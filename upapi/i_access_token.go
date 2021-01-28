package upapi

//调用云闪付开放平台接口，通过CODE获取accessToken和openId
func AccessToken(conf *Config, code, backendToken string) (r *TokenResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", conf.AppId)
	bm.Set("backendToken", backendToken)
	bm.Set("code", code)
	bm.Set("grantType", "authorization_code")

	err = Call(conf, "/token", bm, &r)
	return
}

type TokenResult struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
	OpenId      string `json:"openId"`
	Scope       string `json:"scope"`
}
