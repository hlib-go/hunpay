package ocwap

// 银联全渠道接口通信协议，注意：只对接了5.1.0，其它签名版本未对接
const (
	VERSION             = "5.1.0"                // 银联签名版本：5.1.0
	ENCODING            = "UTF-8"                // 编码
	SIGN_METHOD         = "01"                   // RSA SHA256
	INVALID_REQUEST     = "Invalid request."     // 版本号，交易类型、子类，签名方法，签名值等关键域未上送，返回“Invalid request.”；
	INVALID_REQUEST_URI = "Invalid request URI." // 交易类型和请求地址校验有误，返回“Invalid request URI.”
	BIZ_TYPE            = "000201"               //业务类型，固定值
	CHANNEL_TYPE        = "08"                   // 渠道类型，固定值
	ACCESS_TYPE         = "0"                    // 接入类型 固定值0（0：商户直连接入 1：收单机构接入  2：平台商户接入）
	CURRENCY_CODE       = "156"                  // 交易币种，默认156
)

// 银联商户入网配置参数
type Config struct {
	BaseServiceUrl  string // 服务域名
	MerId           string // 银联商户号
	MerPrivateKey   string // 商户申请的证书私钥，，申请的是pfx格式，需要提取pem格式私钥 ，pfx提取的公钥需要上传到银联商家管理平台，具体步骤参考银联文档
	MerSerialNumber string // 证书序列号，可通过pfx提取序列号

	//EncryptPubKeyCert string // 银联敏感信息加密Key公钥，通过接口查询
	//SignPubKeyCert    string // 银联签名公钥，5.1.0后 通过接口查询
}

func NewConfig(baseServiceUrl, merId, merPrivateKey, merSerialNumber string) *Config {
	return &Config{
		BaseServiceUrl:  baseServiceUrl,
		MerId:           merId,
		MerPrivateKey:   merPrivateKey,
		MerSerialNumber: merSerialNumber,
	}
}
