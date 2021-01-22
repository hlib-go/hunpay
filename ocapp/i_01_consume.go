package ocapp

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// Consume 消费接口
func Consume(cfg *Config, p *ConsumeParams) (result *ConsumeResult, err error) {
	var pm = make(map[string]string)
	// 以下参数根据接口文档与示例填写
	pm["txnType"] = "01"
	pm["txnSubType"] = "01"
	pm["bizType"] = BIZ_TYPE
	pm["merId"] = cfg.MerId
	pm["orderId"] = p.OrderId
	pm["txnAmt"] = p.TxnAmt
	pm["txnTime"] = p.TxnTime // 十四位字符串
	pm["currencyCode"] = CURRENCY_CODE
	pm["accessType"] = ACCESS_TYPE
	pm["channelType"] = CHANNEL_TYPE
	if p.AccNo != "" {
		pm["accNo"] = p.AccNo // 银行卡号
	}
	//pm["frontUrl"] = p.FrontUrl // 前台返回商户结果时使用，前台类交易需上送
	pm["backUrl"] = p.BackUrl //后台返回商户结果时使用，如上送，则发送商户后台交易结果通知，不支持换行符等不可见字符
	if p.ReqReserved != "" {
		pm["reqReserved"] = p.ReqReserved
	}
	err = AppTransReqUnmarshal(cfg, pm, &result)
	if err != nil {
		return
	}

	return
}

type ConsumeParams struct {
	AccNo       string `json:"accNo" description:"非必填,银行卡号，用于帮用户自动填写卡号"`
	OrderId     string `json:"orderId" description:"必填，业务系统订单号，不能重复，长度8到40，不能存在符号"`
	TxnAmt      string `json:"txnAmt" description:"必填，交易金额，单位分"`
	BackUrl     string `json:"backUrl" description:"必填，后台接受支付结果通知URL"`
	TxnTime     string `json:"txnTime" description:"必填，交易时间 ,14位 yyyyMMddHHmmss  商户代码merId、商户订单号orderId、订单发送时间txnTime三要素唯一确定一笔交易。"`
	ReqReserved string `json:"reqReserved" description:"非必填，请求方保留域,通知原样返回"`
}

type ConsumeResult struct {
	RespCode    string `json:"respCode"`
	RespMsg     string `json:"respMsg"`
	Tn          string `json:"tn"`
	BizType     string `json:"bizType"`
	TxnTime     string `json:"txnTime"`
	TxnType     string `json:"txnType"`
	TxnSubType  string `json:"txnSubType"`
	AccessType  string `json:"accessType"`
	ReqReserved string `json:"reqReserved"`
	MerId       string `json:"merId"`
	OrderId     string `json:"orderId"`
	Reserved    string `json:"reserved"`
}

//ConsumeNotifyHandler 消费异步通知结果
func ConsumeNotifyHandler(cbFunc func(o *ConsumeNotifyEntity) error) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var (
			err       error
			requestId = Rand32()
			orderId   string
			reqParams string
			reqBody   string
			resBody   string
		)
		defer func() {
			log.Info(requestId, orderId, "全渠道消费通知 request.RequestURI：", request.RequestURI)
			log.Info(requestId, orderId, "全渠道消费通知 reqParams：", reqParams)
			log.Info(requestId, orderId, "全渠道消费通知 reqBody JSON：", reqBody)
			log.Info(requestId, orderId, "全渠道消费通知 resBody JSON：", resBody)
			if err != nil {
				log.Warn(requestId, "全渠道消费通知处理异常：", err.Error())
				writer.WriteHeader(500)
				writer.Write([]byte(err.Error()))
				return
			}
			writer.Write([]byte(resBody))
		}()

		// 接收通知参数
		rbytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return
		}
		reqParams = string(rbytes)
		// 验证签名与返回码
		bmap, err := NotifyVerify(reqParams)
		if err != nil {
			return
		}

		pbytes, err := json.Marshal(bmap)
		if err != nil {
			return
		}
		reqBody = string(pbytes)

		var entity *ConsumeNotifyEntity
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
	})
}

type ConsumeNotifyEntity struct {
	Version          string `json:"version"`
	Encoding         string `json:"encoding"`
	TxnType          string `json:"txnType"`
	TxnSubType       string `json:"txnSubType"`
	BizType          string `json:"bizType"`
	AccessType       string `json:"accessType"`
	AcqInsCode       string `json:"acqInsCode"`
	MerId            string `json:"merId"`
	OrderId          string `json:"orderId"`
	TxnTime          string `json:"txnTime"`
	TxnAmt           string `json:"txnAmt"`
	CurrencyCode     string `json:"currencyCode"`
	ReqReserved      string `json:"reqReserved"`
	Reserved         string `json:"reserved"`
	QueryId          string `json:"queryId"`
	RespCode         string `json:"respCode"`
	RespMsg          string `json:"respMsg"`
	SettleAmt        string `json:"settleAmt"`
	TraceNo          string `json:"traceNo"`
	TraceTime        string `json:"traceTime"`
	ExchangeDate     string `json:"exchangeDate"`
	ExchangeRate     string `json:"exchangeRate"`
	AccNo            string `json:"accNo"`
	PayCardType      string `json:"payCardType"`
	PayType          string `json:"payType"`
	PayCardNo        string `json:"payCardNo"`
	PayCardIssueName string `json:"payCardIssueName"`
	BindId           string `json:"bindId"`
}
