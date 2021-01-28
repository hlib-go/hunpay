package upapi

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

type BodyMap struct {
	m  map[string]interface{}
	mu sync.RWMutex
}

func NewBodyMap() *BodyMap {
	return &BodyMap{
		m: make(map[string]interface{}),
	}
}

// 设置参数
func (bm *BodyMap) Set(key string, value interface{}) {
	bm.mu.Lock()
	bm.m[key] = value
	bm.mu.Unlock()
}

// 获取参数
func (bm *BodyMap) Get(key string) string {
	if bm.m == nil {
		return ""
	}
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	value, ok := bm.m[key]
	if !ok {
		return ""
	}
	v, ok := value.(string)
	if !ok {
		return convertToString(value)
	}
	return v
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}

// 删除参数
func (bm *BodyMap) Remove(key string) {
	bm.mu.Lock()
	delete(bm.m, key)
	bm.mu.Unlock()
}

// 拼接云闪付签名参数字符串 secret + body + ts
func (bm *BodyMap) SignParams(secret string) string {
	bm.Set("secret", secret)
	var (
		buf     strings.Builder
		keyList []string
	)
	bm.mu.RLock()
	for k := range bm.m {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	bm.mu.RUnlock()
	for i, k := range keyList {
		if i != 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(bm.Get(k))
	}
	s := buf.String()
	return s
}

// 前端upsdk参数签名
func (bm *BodyMap) FrontSignParams() string {
	var (
		buf     strings.Builder
		keyList []string
	)
	bm.mu.RLock()
	for k := range bm.m {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	bm.mu.RUnlock()
	for i, k := range keyList {
		if i != 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(bm.Get(k))
	}
	s := buf.String()
	return s
}

// 计算SHA256签名
//步骤一：拼接待签名字符串，得到string1
//对所有待签名参数按照字段名的ASCII码进行从小到大排序（字典序），然后使用URL键值对的格式（即key1=value1&key2=value2…）拼接成字符串string1，如：
//appId=a5949221470c4059b9b0b45a90c81527&frontToken=U72eJp21SkuzKRdUK+jyFw==&nonceStr=Wm3WZYTPz0wzccnW&timestamp=1414587457&url=http://mobile.xxx.com?params=value
//>string1 := bm.SignParams(secret)
//步骤二：对待签名字符串进行SHA256签名，得到signature
//将步骤一得到的待签名字符串string1转换成byte数组，传入方法sha256(byte[] data)中，执行后将返回签名结果signature。
//如：a4bb34a2b60aa34ec4f03754547ca3e39a80e628b9760323d10561997935bb42。
//>return fmt.Sprintf("%x", sha256.Sum256([]byte(string1)))
func (bm *BodyMap) Sha256Sign(secret string) string {
	string1 := bm.SignParams(secret)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(string1)))
}

func (bm *BodyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(bm.m)
}
