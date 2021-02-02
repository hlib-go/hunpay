package ocapp

import "errors"

// 文档：https://open.unionpay.com/tjweb/acproduct/APIList?acpAPIId=338&apiservId=453&version=V2.2&bussType=0

// Pubkey 全渠道平台敏感信息加密公钥查询
// 由业务系统缓存，缓存有效期12个小时，失效重新查询
func Pubkey(cfg *Config) (encryptPubKeyCert string, err error) {
	// 以下字段根据文档添加，值为约定固定值，都是必填字段
	var bm = make(map[string]string)
	bm["txnType"] = "95"
	bm["txnSubType"] = "00"
	bm["channelType"] = CHANNEL_TYPE
	bm["bizType"] = BIZ_TYPE
	bm["accessType"] = ACCESS_TYPE
	bm["certType"] = "01" // 证书类型固定值 01：敏感信息加密公钥
	bm["orderId"] = Rand32()
	bm["txnTime"] = TxnTime()

	respMap, err := BackTransReq("pubkey", cfg, bm)
	if err != nil {
		return
	}
	encryptPubKeyCert = respMap["encryptPubKeyCert"]
	if encryptPubKeyCert == "" {
		err = errors.New("接口返回参数 encryptPubKeyCert 为空.[UNPAY_OCAPP_PUBKEY]")
		return
	}
	return
}
