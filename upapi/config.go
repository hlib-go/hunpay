package upapi

// 银联云闪付接口配置
type Config struct {
	BaseServiceUrl string `json:"baseServiceUrl"  description:"接口链接前缀，结尾不带'/'"` // https://open.95516.com/open/access/1.0
	AppId          string `json:"appId,omitempty" description:"接入方唯一标识"`
	Secret         string `json:"secret,omitempty" description:"接入方秘钥，用于基础令牌接口的签名"`
	SymmetricKey   string `json:"symmetricKey,omitempty" description:"对称密钥（3DES，16进制格式） 。用于后台敏感数据解密"`
	UpPublicKey    string `json:"upPublicKey,omitempty" description:"（银联方）使用openssl生成，base64形式输出"`
	MchPrivateKey  string `json:"mchPrivateKey" description:"接入商户证书私钥"`

	RefreshFrontToken   func() // frontToken失效时调用此函数，业务实现刷新逻辑
	RefreshBackendToken func() // backendToken失效时调用此函数，业务实现刷新逻辑
}
