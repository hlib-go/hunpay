package ocapp

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 文档： https://open.unionpay.com/tjweb/acproduct/APIList?acpAPIId=336&apiservId=453&version=V2.2&bussType=0

// 退货接口
func Refund(cfg *Config, p *RefundParams) (result *RefundResult, err error) {
	var pm = make(map[string]string)
	// 以下参数根据接口文档与示例填写
	pm["txnType"] = "04"
	pm["txnSubType"] = "00"
	pm["bizType"] = "000201"
	pm["merId"] = cfg.MerId
	pm["origQryId"] = p.OrigQryId
	pm["orderId"] = p.OrderId
	pm["txnAmt"] = p.TxnAmt
	pm["txnTime"] = p.TxnTime // 十四位字符串
	pm["accessType"] = "0"
	pm["channelType"] = "08"
	pm["currencyCode"] = "156" // 交易币种
	pm["backUrl"] = p.BackUrl
	if p.ReqReserved != "" {
		pm["reqReserved"] = p.ReqReserved
	}
	// 返回表单POST请求参数
	err = BackTransReqUnmarshal("refund", cfg, pm, &result)
	if err != nil {
		return
	}
	// respCode=45&respMsg=已被成功退货或已被成功撤销[2004003]
	return
}

type RefundParams struct {
	OrigQryId   string `json:"origQryId" description:"原始消费交易的queryId"`
	OrderId     string `json:"orderId" description:"退款交易订单号，非原交易订单号"`
	TxnTime     string `json:"txnTime" description:"退款交易时间，非原交易订单号"`
	TxnAmt      string `json:"txnAmt" description:"退款金额"`
	BackUrl     string `json:"backUrl" description:"退货结果通知URL"`
	ReqReserved string `json:"reqReserved" description:"商户自定义保留域，交易应答时会原样返回"`
}

type RefundResult struct {
	QueryId     string `json:"queryId"`
	RespCode    string `json:"respCode"`
	RespMsg     string `json:"respMsg"`
	BizType     string `json:"bizType"`
	TxnTime     string `json:"txnTime"`
	TxnAmt      string `json:"txnAmt"`
	TxnType     string `json:"txnType"`
	TxnSubType  string `json:"txnSubType"`
	AccessType  string `json:"accessType"`
	ReqReserved string `json:"reqReserved"`
	MerId       string `json:"merId"`
	OrderId     string `json:"orderId"`
	OrigQryId   string `json:"origQryId"`
}

//RefundNotifyHandler 消费退款异步通知结果
func RefundNotifyHandler(cbFunc func(o *RefundNotifyEntity) error) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		resBody, err := RefundNotify("", request, cbFunc)
		if err != nil {
			writer.WriteHeader(500)
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write([]byte(resBody))
	})
}

func RefundNotify(requestId string, request *http.Request, cbFunc func(o *RefundNotifyEntity) error) (resBody string, err error) {
	if requestId == "" {
		requestId = Rand32()
	}
	var (
		nlog      = log.WithField("requestId", requestId)
		orderId   string
		reqParams string
		reqBody   string
		//resBody   string
	)
	defer func() {
		nlog.Info(orderId, " 全渠道消费退款通知 request.RequestURI：", request.RequestURI)
		nlog.Info(orderId, " 全渠道消费退款通知 reqParams：", reqParams)
		nlog.Info(orderId, " 全渠道消费退款通知 reqBody：", reqBody)
		nlog.Info(orderId, " 全渠道消费退款通知 resBody：", resBody)
		if err != nil {
			nlog.Warn(orderId, " 全渠道消费退款通知处理异常：", err.Error())
			return
		}
	}()

	// 接收通知参数
	rbytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}
	if rbytes == nil || string(rbytes) == "" {
		err = errors.New("无效请求，没有接收到消费退款通知参数")
		return
	}

	params, err := url.QueryUnescape(string(rbytes))
	if err != nil {
		return
	}
	reqParams = params

	bmap := ParamsConvertMap(params)
	orderId = bmap["orderId"]

	// 验签
	signValue := RsaSignSortMap(bmap)
	err = RsaWithSha256Verify(signValue, bmap["signature"], bmap["signPubKeyCert"])
	if err != nil {
		return
	}

	// 验证返回码
	if bmap["respCode"] != RESP_OK {
		err = errors.New("UP" + bmap["respCode"] + ":" + bmap["respMsg"])
		return
	}

	pbytes, err := json.Marshal(bmap)
	if err != nil {
		return
	}
	reqBody = string(pbytes)

	var entity *RefundNotifyEntity
	err = json.Unmarshal(pbytes, &entity)
	if err != nil {
		return
	}

	// 回调业务函数
	err = cbFunc(entity)
	if err != nil {
		return
	}
	resBody = `{"respCode":"00","requestId":` + requestId + `,"orderId":"` + orderId + `"}`
	return
}

type RefundNotifyEntity struct {
	QueryId            string `json:"queryId"`
	CurrencyCode       string `json:"currencyCode"`
	TraceTime          string `json:"traceTime"`
	SignMethod         string `json:"signMethod"`
	SettleCurrencyCode string `json:"settleCurrencyCode"`
	SettleAmt          string `json:"settleAmt"`
	SettleDate         string `json:"settleDate"`
	TraceNo            string `json:"traceNo"`
	RespCode           string `json:"respCode"`
	RespMsg            string `json:"respMsg"`
	ExchangeDate       string `json:"exchangeDate"`
	SignPubKeyCert     string `json:"signPubKeyCert"`
	ExchangeRate       string `json:"exchangeRate"`
	AcqInsCode         string `json:"acqInsCode"`
	AccNo              string `json:"accNo"`
	Version            string `json:"version"`
	Encoding           string `json:"encoding"`
	BizType            string `json:"bizType"`
	TxnTime            string `json:"txnTime"`
	TxnAmt             string `json:"txnAmt"`
	TxnType            string `json:"txnType"`
	TxnSubType         string `json:"txnSubType"`
	AccessType         string `json:"accessType"`
	ReqReserved        string `json:"reqReserved"`
	MerId              string `json:"merId"`
	OrderId            string `json:"orderId"` // 订单号是退款交易的订单号，非消费交易订单号
	OrigQryId          string `json:"origQryId"`
	Reserved           string `json:"reserved"`
}
