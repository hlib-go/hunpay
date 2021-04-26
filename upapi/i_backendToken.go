package upapi

import "time"

// 基础服务令牌，通过appId、secret换取，有效期为7200秒，后台接口调用凭证。
func BackendToken(conf *Config) (r *BackendTokenResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", conf.AppId)
	bm.Set("nonceStr", GetRandomString(16))
	bm.Set("timestamp", time.Now().Unix())

	err = Call(conf, "/backendToken", bm, &r)
	return
}

type BackendTokenResult struct {
	BackendToken string `json:"backendToken"`
	ExpiresIn    int64  `json:"expiresIn"`
}
