package upsdk

//调用云闪付开放平台接口，通过CODE获取accessToken和openId
func (c *Unionpay) AccessToken(code string) (r *TokenResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.Config.AppId)
	bm.Set("backendToken", c.GetBackendToken(false))
	bm.Set("code", code)
	bm.Set("grantType", "authorization_code")

	err = c.Call("/token", bm, &r)
	return
}

type TokenResult struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
	OpenId      string `json:"openId"`
	Scope       string `json:"scope"`
}
