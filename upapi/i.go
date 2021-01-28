package upapi

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RespBody struct {
	Resp   string      `json:"resp"`
	Msg    string      `json:"msg"`
	Params interface{} `json:"params"`
}

// 云闪付接口 POST请求
func Post(c *Config, path string, bodyMap *BodyMap) (respBody *RespBody, err error) {
	var (
		reqBody string
		resBody string
	)
	defer func() {
		fmt.Println("请求URL：" + c.ServiceUrl + path)
		fmt.Println("请求报文：" + reqBody)
		fmt.Println("响应报文：" + resBody)
	}()

	plog := log.WithField("path", path)
	plog.WithField("serviceUrl", c.ServiceUrl).WithField("appid", c.AppId).Info("云闪付请求URL")
	bodyBytes, _ := json.Marshal(bodyMap)
	reqBody = string(bodyBytes)
	plog.WithField("requestBody", reqBody).Info("云闪付请求报文")
	resp, err := http.Post(c.ServiceUrl+path, "application/json", strings.NewReader(reqBody))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		plog.WithField("httpStatus", resp.Status).Warn("云闪付响应状态异常")
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(bytes)
	plog.WithField("responseBody", resBody).Warn("云闪付响应报文")
	err = json.Unmarshal(bytes, &respBody)
	if err != nil {
		return
	}
	return
}

func Call(c *Config, path string, bm *BodyMap, result interface{}) (err error) {
	plog := log.WithField("path", path)
	begmillisecond := time.Now().UnixNano() / 1e6
	//计算签名
	signature := bm.Sha256Sign(c.Secret)
	bm.Set("signature", signature)

	body, err := json.Marshal(bm.m)
	if err != nil {
		return
	}
	plog.WithField("requestBody", string(body)).Info("云闪付请求报文")
	resp, err := http.Post(c.ServiceUrl+path, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	endmillisecond := time.Now().UnixNano() / 1e6
	plog.WithField("millisecond", strconv.FormatInt(endmillisecond-begmillisecond, 10)).WithField("responseBody", string(bytes)).Info("云闪付响应报文")
	var respBody *RespBody
	err = json.Unmarshal(bytes, &respBody)
	if err != nil {
		return
	}
	/*
		a10	不合法的backend_token，或已过期（参见6.1.1获取backendToken章节，重新获取backend_token）
		a20	不合法的frontend_token，或已过期（参见6.1.2获取frontToken章节，重新获取front_token）
		a31	不合法的授权code，或已过期（参见5.3系统对接步骤章节，参见常见问题解答）
	*/
	if respBody.Resp == EA10.Code {
		if c.RefreshBackendToken != nil {
			c.RefreshBackendToken()
		}
		return EA10
	}
	if respBody.Resp == EA20.Code {
		if c.RefreshFrontToken != nil {
			c.RefreshFrontToken()
		}
		return EA20
	}

	if respBody.Resp != E00.Code {
		return errors.New(respBody.Resp + ":" + respBody.Msg)
	}
	b, err := json.Marshal(respBody.Params)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, result)
	if err != nil {
		return
	}
	return
}

// 获取随机字符串
//    length：字符串长度
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(str)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

// 云闪付敏感数据解密
func Decode3DES(symmetricKey string, v string) (val string, err error) {
	if v == "" {
		val = ""
		return
	}
	bytes, err := hex.DecodeString(symmetricKey)
	if err != nil {
		return
	}
	key := string(bytes)
	val, err = TRIPLE_DES_ECB_PKCS5_Decode(v, key)
	return
}

// 云闪付敏感数据加密
func Encode3DES(symmetricKey string, v string) (val string, err error) {
	if v == "" {
		val = ""
		return
	}
	bytes, err := hex.DecodeString(symmetricKey)
	if err != nil {
		return
	}
	key := string(bytes)
	val, err = TRIPLE_DES_ECB_PKCS5_Encode(v, key)
	return
}
