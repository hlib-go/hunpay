package upsdk

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

// 银联云闪付接口请求客户端
type Unionpay struct {
	Config          *Config
	GetBackendToken func(refresh bool) string
	GetFrontToken   func(refresh bool) string

	MchId string
	log   *log.Entry
}

func New(conf *Config, mchId string) *Unionpay {
	up := &Unionpay{Config: conf, MchId: mchId}
	up.log = log.WithField("mchId", up.MchId)
	return up
}

type RespBody struct {
	Resp   string      `json:"resp"`
	Msg    string      `json:"msg"`
	Params interface{} `json:"params"`
}

// 云闪付接口 POST请求
func (c *Unionpay) Post(path string, bodyMap *BodyMap) (respBody *RespBody, err error) {
	var (
		reqBody string
		resBody string
	)
	defer func() {
		fmt.Println("请求URL：" + c.Config.BaseServiceUrl + path)
		fmt.Println("请求报文：" + reqBody)
		fmt.Println("响应报文：" + resBody)
	}()

	plog := c.log.WithField("path", path)
	plog.WithField("serviceUrl", c.Config.BaseServiceUrl).WithField("appid", c.Config.AppId).Info("云闪付请求URL")
	bodyBytes, _ := json.Marshal(bodyMap)
	reqBody = string(bodyBytes)
	plog.WithField("requestBody", reqBody).Info("云闪付请求报文")
	resp, err := http.Post(c.Config.BaseServiceUrl+path, "application/json", strings.NewReader(reqBody))
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

func (c *Unionpay) Call(path string, bm *BodyMap, result interface{}) (err error) {
	plog := c.log.WithField("path", path)
	begmillisecond := time.Now().UnixNano() / 1e6
	//计算签名
	signature := bm.Sha256Sign(c.Config.Secret)
	bm.Set("signature", signature)

	body, err := json.Marshal(bm.m)
	if err != nil {
		return
	}
	plog.WithField("requestBody", string(body)).Info("云闪付请求报文")
	resp, err := http.Post(c.Config.BaseServiceUrl+path, "application/json", strings.NewReader(string(body)))
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
	/*
		a10	不合法的backend_token，或已过期（参见6.1.1获取backendToken章节，重新获取backend_token）
		a20	不合法的frontend_token，或已过期（参见6.1.2获取frontToken章节，重新获取front_token）
		a31	不合法的授权code，或已过期（参见5.3系统对接步骤章节，参见常见问题解答）
	*/
	if respBody.Resp == "a10" {
		c.GetBackendToken(true)
	}
	if respBody.Resp == "a20" {
		c.GetFrontToken(true)
	}

	if respBody.Resp != "00" {
		return errors.New(respBody.Resp + ":" + respBody.Msg)
	}
	b, err := json.Marshal(respBody.Params)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, result)
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
func (c *Unionpay) Decode3DES(v string) (val string, err error) {
	if v == "" {
		val = ""
		return
	}
	bytes, err := hex.DecodeString(c.Config.SymmetricKey)
	if err != nil {
		return
	}
	key := string(bytes)
	val, err = TRIPLE_DES_ECB_PKCS5_Decode(v, key)
	return
}

// 云闪付敏感数据加密
func (c *Unionpay) Encode3DES(v string) (val string, err error) {
	if v == "" {
		val = ""
		return
	}
	bytes, err := hex.DecodeString(c.Config.SymmetricKey)
	if err != nil {
		return
	}
	key := string(bytes)
	val, err = TRIPLE_DES_ECB_PKCS5_Encode(v, key)
	return
}
