package ocwap

import "errors"

// 文档：https://open.unionpay.com/tjweb/acproduct/APIList?acpAPIId=338&apiservId=453&version=V2.2&bussType=0

// 全渠道公钥更新,首次调用，需要从银联商户平台下载公钥
func Pubkey(cfg *Config) (signPubKeyCert string, err error) {
	// 以下字段根据文档添加，值为约定固定值，都是必填字段
	var bm = make(map[string]string)
	bm["txnType"] = "95"
	bm["txnSubType"] = "00"
	bm["channelType"] = "08"
	bm["bizType"] = "000201"
	bm["accessType"] = "0"
	bm["orderId"] = Rand32()
	bm["txnTime"] = txnTime()

	respMap, err := BackTransReq(cfg, bm)
	if err != nil {
		return
	}
	signPubKeyCert = respMap["signPubKeyCert"]
	if signPubKeyCert == "" {
		err = errors.New("接口返回参数 signPubKeyCert 为空")
		return
	}
	return
}
