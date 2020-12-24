package ocwap

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 消费撤销接口
func UnConsume(cfg *Config, p *UnConsumeParams) (result *UnConsumeResult, err error) {
	var pm = make(map[string]string)
	// 以下参数根据接口文档与示例填写
	pm["txnType"] = "31"
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
	pm["reqReserved"] = p.ReqReserved
	// 返回表单POST请求参数
	err = BackTransReqUnmarshal(cfg, pm, &result)
	if err != nil {
		return
	}
	return
}

type UnConsumeParams struct {
	OrigQryId   string `json:"origQryId" description:"原始消费交易的queryId"`
	OrderId     string `json:"orderId" description:"必填，业务系统订单号，不能重复"`
	TxnAmt      string `json:"txnAmt" description:"必填，交易金额，单位分"`
	BackUrl     string `json:"backUrl" description:"必填，后台接受撤销结果的通知URL"`
	TxnTime     string `json:"txnTime" description:"必填，订单发送时间  商户代码merId、商户订单号orderId、订单发送时间txnTime三要素唯一确定一笔交易。"`
	ReqReserved string `json:"reqReserved" description:"非必填，请求方保留域,通知原样返回"`
}

// 消费撤销接口响应
type UnConsumeResult struct {
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
	Reserved    string `json:"reserved"`
}

//UnConsumeNotifyHandler 消费撤销异步通知结果
func UnConsumeNotifyHandler(cbFunc func(o *UnConsumeNotifyEntity) error) http.Handler {
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
			log.Info(requestId, orderId, "全渠道消费撤销通知 request.RequestURI：", request.RequestURI)
			log.Info(requestId, orderId, "全渠道消费撤销通知 reqParams：", reqParams)
			log.Info(requestId, orderId, "全渠道消费撤销通知 reqBody：", reqBody)
			log.Info(requestId, orderId, "全渠道消费撤销通知 resBody：", resBody)
			if err != nil {
				log.Error(requestId, "全渠道消费通知处理异常：", err.Error())
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

		var entity *UnConsumeNotifyEntity
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

type UnConsumeNotifyEntity struct {
	QueryId            string `json:"queryId"`
	CurrencyCode       string `json:"currencyCode"`
	TraceTime          string `json:"traceTime"`
	Signature          string `json:"signature"`
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
	OrderId            string `json:"orderId"`
	OrigQryId          string `json:"origQryId"`
	Reserved           string `json:"reserved"`
}
