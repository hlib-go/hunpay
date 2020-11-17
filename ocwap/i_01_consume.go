package ocwap

import (
	"net/http"
	"strings"
)

// Consume 手机（wap）消费接口
func Consume(cfg *Config, p *ConsumeParams, writer http.ResponseWriter) (err error) {
	var pm = make(map[string]string)
	// 以下参数根据接口文档与示例填写
	pm["txnType"] = "01"
	pm["txnSubType"] = "01"
	pm["bizType"] = "000201"
	pm["merId"] = cfg.MerId
	pm["orderId"] = p.OrderId
	pm["txnAmt"] = p.TxnAmt
	pm["txnTime"] = p.TxnTime // 十四位字符串
	pm["currencyCode"] = "156"
	pm["accessType"] = "0"
	pm["channelType"] = "08"
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
	OrderId     string `json:"orderId" description:"必填，业务系统订单号，不能重复"`
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
