package upsdk

func (c *Unionpay) UserMobile(accessToken, openId string) (r *UserMobileResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.Config.AppId)
	bm.Set("backendToken", c.GetBackendToken(false))
	bm.Set("accessToken", accessToken)
	bm.Set("openId", openId)

	err = c.Call("/user.mobile", bm, &r)
	// 解密密文手机号
	r.Mobile, err = c.Decode3DES(r.Mobile)
	if err != nil {
		return
	}
	return
}

type UserMobileResult struct {
	Mobile string `json:"mobile"`
}
