package upapi

import "time"

// 基础服务令牌，通过appId、secret换取，有效期为7200秒，upsdk初始化凭证。
func FrontToken(conf *Config) (r *FrontTokenResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", conf.AppId)
	bm.Set("nonceStr", GetRandomString(16))
	bm.Set("timestamp", time.Now().Unix())

	err = Call(conf, "/frontToken", bm, &r)
	return
}

type FrontTokenResult struct {
	FrontToken string `json:"frontToken"`
	ExpiresIn  int64  `json:"expiresIn"`
}
