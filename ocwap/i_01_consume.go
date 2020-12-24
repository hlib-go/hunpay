package ocwap

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

// Consume 手机（wap）消费接口
func Consume(cfg *Config, p *ConsumeParams, writer http.ResponseWriter) (err error) {
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
	// 如果存在卡号，打开链接时无填写银行卡号的步骤
	if p.AccNo != "" {
		pm["accType"] = "01"  //  账号类型，01：银行卡
		pm["accNo"] = p.AccNo // 银行卡号
	}
	pm["frontUrl"] = p.FrontUrl
	pm["backUrl"] = p.BackUrl
	pm["reqReserved"] = p.ReqReserved
	// 返回表单POST请求参数
	action, items, err := FrontTransReq(cfg, pm)
	if err != nil {
		return
	}
	payFormHtml := postForm(action, items)

	// 跳转银联支付，POST表单请求
	writer.Header().Set("Content-Type", "text/html;charset=UTF-8")
	writer.Write([]byte(payFormHtml))
	return
}

type ConsumeParams struct {
	AccNo       string `json:"accNo" description:"非必填,银行卡号，用于帮用户自动填写卡号"`
	OrderId     string `json:"orderId" description:"必填，业务系统订单号，不能重复，长度8到40，不能存在符号"`
	TxnAmt      string `json:"txnAmt" description:"必填，交易金额，单位分"`
	FrontUrl    string `json:"frontUrl" description:"必填，支付完成后跳转的页面"`
	BackUrl     string `json:"backUrl" description:"必填，后台接受支付结果通知URL"`
	TxnTime     string `json:"txnTime" description:"必填，交易时间 ,14位 yyyyMMddHHmmss  商户代码merId、商户订单号orderId、订单发送时间txnTime三要素唯一确定一笔交易。"`
	ReqReserved string `json:"reqReserved" description:"非必填，请求方保留域,通知原样返回"`
}

// 构建表单POST请求字符串
func postForm(action string, items map[string]string) string {
	var html strings.Builder
	html.WriteString("<meta charset='UTF-8'>")
	html.WriteString("<div style='display:none;'>")
	html.WriteString("<form method='POST' action='" + action + "' id='form'>")
	for k, v := range items {
		html.WriteString(k + " <input type='hidden' name='" + k + "' value='" + v + "'><br>")
	}
	html.WriteString("</form>")
	html.WriteString("</div>")
	html.WriteString("<script>document.getElementById('form').submit();</script>")
	return html.String()
}

// 消费同步通知结果（页面跳转）
func ConsumeNotifyFrontHandler(cbFunc func(writer http.ResponseWriter, request *http.Request, entity *ConsumeNotifyEntity, err error)) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var (
			entity    *ConsumeNotifyEntity
			err       error
			requestId = Rand32()
			reqParams string
			reqBody   string
		)
		defer func() {
			log.Info(requestId, "全渠道消费前端通知 request.RequestURI：", request.RequestURI)
			log.Info(requestId, "全渠道消费前端通知 reqParams：", reqParams)
			log.Info(requestId, "全渠道消费前端通知 reqBody JSON：", reqBody)
			if err != nil {
				log.Info(requestId, "全渠道消费前端通知处理异常：", err.Error())
			}
			cbFunc(writer, request, entity, err)
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

		err = json.Unmarshal(pbytes, &entity)
		if err != nil {
			return
		}
	})
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
				log.Info(requestId, "全渠道消费通知处理异常：", err.Error())
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
