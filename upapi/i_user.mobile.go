package upapi

// 取用户手机号码
func UserMobile(c *Config, accessToken, backendToken, openId string) (r *UserMobileResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.AppId)
	bm.Set("backendToken", backendToken)
	bm.Set("accessToken", accessToken)
	bm.Set("openId", openId)

	err = Call(c, "/user.mobile", bm, &r)
	// 解密密文手机号
	r.Mobile, err = Decode3DES(c.SymmetricKey, r.Mobile)
	if err != nil {
		return
	}
	return
}

type UserMobileResult struct {
	Mobile string `json:"mobile"`
}
