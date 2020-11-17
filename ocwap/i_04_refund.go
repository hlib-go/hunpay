package ocwap

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
	pm["reqReserved"] = p.ReqReserved
	// 返回表单POST请求参数
	err = BackTransReqUnmarshal(cfg, pm, &result)
	if err != nil {
		return
	}
	return
}

type RefundParams struct {
	OrigQryId   string `json:"origQryId" description:"原始消费交易的queryId"`
	OrderId     string `json:"orderId"`
	TxnTime     string `json:"txnTime" description:"订单发送时间"`
	TxnAmt      string `json:"txnAmt" description:"交易金额"`
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
