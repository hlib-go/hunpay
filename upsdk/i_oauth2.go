package upsdk

import (
	"fmt"
	"net/url"
)

// OAUTH2 用户联登
/*
upapi_base	基础授权，云闪付不会弹出授权页面，仅可以获取用户openId
upapi_mobile	用户手机号授权，云闪付会弹出授权页面，仅可获取用户手机号
upapi_auth	用户实名授权（手机号[可选]+姓名+证件号），云闪付会弹出授权页面，获取用户实名信息，亦可以获取用户手机号，但仍需调用获取手机号接口，具体返回请以用户是否授权获取为准
upapi_contract	无感支付 【授权链接参照如下】
https://open.95516.com/s/open/noPwd/html/open.html?
appId=XXX
&redirectUri=https://XXX/XXX
&responseType=code
&scope=upapi_contract
&planId=9ca088374a8141ff95dca35a12d73240
upapi_mshpush	消息推送授权，云闪付不会弹出授权页面，授权完可以向用户推送消息

注：云闪付只会向特定行业开放实名认证要素信息
*/

// 仅用于云闪付APP内
func (c *Unionpay) Oauth2(redirectUri, scope, state string) string {
	return fmt.Sprintf("https://open.95516.com/s/open/html/oauth.html?appId=%s&redirectUri=%s&responseType=code&scope=%s&state=%s",
		c.Config.AppId,
		url.QueryEscape(redirectUri),
		scope,
		state)
}
