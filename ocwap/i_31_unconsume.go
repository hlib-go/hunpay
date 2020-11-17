package ocwap

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
