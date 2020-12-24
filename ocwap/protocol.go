package ocwap

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// Method ：POST
// ContentType ： application/x-www-form-urlencoded;charset=utf-8
func Post(serviceUrl, body string) (resBytes []byte, err error) {
	contentType := "application/x-www-form-urlencoded;charset=utf-8"
	var (
		begTime   = time.Now().UnixNano()
		endTime   int64
		requestId = Rand32()
	)
	defer func() {
		log.Info(requestId, "ocwap请求URL：POST "+serviceUrl+"    "+contentType)
		log.Info(requestId, "ocwap请求报文", body)
		log.Info(requestId, "ocwap响应报文", url.QueryEscape(string(resBytes)))
		log.Info(requestId, "ocwap请求耗时", (endTime-begTime)/1e6, "ms")
	}()
	resp, err := http.Post(serviceUrl, contentType, strings.NewReader(body))
	endTime = time.Now().UnixNano()
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("ERROR:" + resp.Status)
		return
	}
	resBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if INVALID_REQUEST == string(resBytes) {
		err = errors.New(INVALID_REQUEST)
		return
	}
	if INVALID_REQUEST_URI == string(resBytes) {
		err = errors.New(INVALID_REQUEST_URI)
		return
	}
	return
}

// 前台接口交易
func FrontTransReq(cfg *Config, bm map[string]string) (url string, kv map[string]string, err error) {
	var (
		requestId = Rand32()
		reqBody   string
	)
	defer func() {
		log.Info(requestId, "ocwap前端请求地址", url)
		log.Info(requestId, "ocwap前端请求报文", reqBody)
	}()
	url = cfg.BaseServiceUrl + "/gateway/api/frontTransReq.do"
	bm["version"] = VERSION
	bm["encoding"] = ENCODING
	bm["signMethod"] = SIGN_METHOD
	bm["certId"] = cfg.MerSerialNumber //签名方式01需要上送
	bm["merId"] = cfg.MerId

	// 签名字符串
	signString := RsaSignSortMap(bm)

	// 计算签名
	sign, err := RsaWithSha256Sign(signString, cfg.MerPrivateKey)
	if err != nil {
		return
	}
	bm["signature"] = sign
	kv = bm

	requestBodyBytes, _ := json.Marshal(bm)
	reqBody = string(requestBodyBytes)
	return
}

// 后台接口交易
func BackTransReq(conf *Config, bm map[string]string) (resMap map[string]string, err error) {
	var (
		url = conf.BaseServiceUrl + "/gateway/api/backTransReq.do"
	)
	bm["version"] = VERSION
	bm["encoding"] = ENCODING
	bm["signMethod"] = SIGN_METHOD      //01（表示采用RSA签名） HASH表示散列算法
	bm["certId"] = conf.MerSerialNumber //签名方式01需要上送
	bm["merId"] = conf.MerId

	// 签名字符串
	signString := RsaSignSortMap(bm)

	// 计算签名
	sign, err := RsaWithSha256Sign(signString, conf.MerPrivateKey)
	if err != nil {
		return
	}
	bm["signature"] = sign

	// 请求字符串
	reqBody := MapConvertParams(bm)

	// HTTP POST
	resBytes, err := Post(url, reqBody)
	if err != nil {
		return
	}
	resStr := string(resBytes)

	// 响应报文参数转Map
	resMap = ParamsConvertMap(resStr)

	// 签名字符串
	signResString := RsaSignSortMap(resMap)

	// 验证签名
	err = RsaWithSha256Verify(signResString, resMap["signature"], resMap["signPubKeyCert"])
	if err != nil {
		return
	}

	// 验证响应状态码
	if resMap["respCode"] != RESP_OK {
		err = errors.New("UP" + resMap["respCode"] + ":" + resMap["respMsg"])
	}
	return
}

// 后台请求，响应结果转为结构体
func BackTransReqUnmarshal(cfg *Config, bm map[string]string, result interface{}) (err error) {
	resMap, err := BackTransReq(cfg, bm)
	if err != nil {
		return
	}
	resBytes, err := json.Marshal(resMap)
	if err != nil {
		return
	}
	err = json.Unmarshal(resBytes, result)
	if err != nil {
		return
	}
	return
}

// 参数字符串转Map
func ParamsConvertMap(params string) (bmap map[string]string) {
	bmap = make(map[string]string)
	s1 := strings.Split(params, "&")
	for _, item := range s1 {
		index := strings.Index(item, "=")
		if index == -1 {
			continue
		}
		k := item[0:index]
		v := item[index+1:]
		bmap[k] = v
	}
	return
}

// map转参数字符串
func MapConvertParams(bmap map[string]string) (params string) {
	for k, v := range bmap {
		params += strings.TrimSpace(k) + "=" + url.QueryEscape(strings.TrimSpace(v)) + "&"
	}
	params = params[0 : len(params)-1]
	return
}

// 拼接待签名字符串
func RsaSignSortMap(params map[string]string) string {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range params {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		k = strings.TrimSpace(k)
		if "signature" == k || k == "" {
			continue
		}
		v := strings.TrimSpace(params[k])
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(v)
		buf.WriteByte('&')
	}
	s := buf.String()
	s = s[0 : len(s)-1]
	return s
}

// 合并map
func mergeMap(bm map[string]string, pm map[string]string) map[string]string {
	if bm == nil {
		bm = make(map[string]string)
	}
	if pm == nil {
		pm = make(map[string]string)
	}
	for k, v := range pm {
		bm[k] = v
	}
	return bm
}

// 生成交易时间
func TxnTime() string {
	return time.Now().Format("20060102150405")
}
