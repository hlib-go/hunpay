package ocwap

// 文档：https://open.unionpay.com/tjweb/acproduct/APIList?acpAPIId=337&apiservId=453&version=V2.2&bussType=0

// Query 交易状态查询接口
func Query(cfg *Config, orderId string) (result *QueryResult, err error) {
	var pm = make(map[string]string)
	// 以下参数根据接口文档与示例填写
	pm["txnType"] = "00"
	pm["txnSubType"] = "00"
	pm["bizType"] = BIZ_TYPE
	pm["accessType"] = ACCESS_TYPE
	pm["txnTime"] = TxnTime()
	pm["orderId"] = orderId

	err = BackTransReqUnmarshal(cfg, pm, &result)
	if err != nil {
		return
	}
	return
}

// 注意：响应字段所有都是string类型，有其它类型需求在业务系统中转换
type QueryResult struct {
	QueryId            string `json:"queryId"`
	TraceTime          string `json:"traceTime"`
	TxnType            string `json:"txnType"`
	TxnSubType         string `json:"txnSubType"`
	SettleCurrencyCode string `json:"settleCurrencyCode"`
	SettleAmt          string `json:"settleAmt"`
	SettleDate         string `json:"settleDate"`
	TraceNo            string `json:"traceNo"`
	RespCode           string `json:"respCode"`
	RespMsg            string `json:"respMsg"`
	BindId             string `json:"bindId"`
	ExchangeDate       string `json:"exchangeDate"`
	IssuerIdentifyMode string `json:"issuerIdentifyMode"`
	CurrencyCode       string `json:"currencyCode"`
	TxnAmt             string `json:"txnAmt"`
	ExchangeRate       string `json:"exchangeRate"`
	AcqInsCode         string `json:"acqInsCode"`
	CardTransData      string `json:"cardTransData"`
	OrigRespCode       string `json:"origRespCode"`
	OrigRespMsg        string `json:"origRespMsg"`
	AccNo              string `json:"accNo"`
	PayType            string `json:"payType"`
	PayCardNo          string `json:"payCardNo"`
	PayCardType        string `json:"payCardType"`
	PayCardIssueName   string `json:"payCardIssueName"`
	Version            string `json:"version"`
	Encoding           string `json:"encoding"`
	TxnTime            string `json:"txnTime"`
	AccessType         string `json:"accessType"`
	MerId              string `json:"merId"`
	OrderId            string `json:"orderId"`
	Reserved           string `json:"reserved"`
	ReqReserved        string `json:"reqReserved"`
}
