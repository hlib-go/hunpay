package upapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
返回码	描述
94	没有通过人像验证，请次日后再试
95	您没有通过人像验证
96	当前用户未开通人像照片认证
a10	不合法的backend_token，或已过期（参见6.1.1获取backendToken章节，重新获取backend_token）
a20	不合法的frontend_token，或已过期（参见6.1.2获取frontToken章节，重新获取front_token）
a31	不合法的授权code，或已过期（参见5.3系统对接步骤章节，参见常见问题解答）
BC2126	缓存中不包含此商户appId
BC2127	缓存中不包含此backendToken
BC2133	获取商户信息异常（检查appId参数信息，若是无感支付，并检查planId）
BC0024	非法请求（接口返回检查时间戳，upsdk报错检查时间戳和安全域名）
BC0025	验签失败（检查签名因子和签名方法）
N60005	cdhdUsrId为空
N62100	商户appId不能为空
N62101	backendToken不能为空
N62102	frontToken不能为空
N62103	scope不能为空
N62104	签名字段值不能为空
N62105	responseType应为固定值code
N62106	grantType应为固定值authoriz
N62107	随机字符串不能为空
N62108	时间戳不能为空
N62109	accessToken不能为空
N62110	openId不能为空
N62111	url不能为空
N62112	code不能为空
N62113	contractCode不能为空
N62114	planId不能为空
N62115	contractId不能为空
N62116	content不能为空
N62117	secret不能为空
N62118	actionType不能为空
N62119	channelNo不能为空
N62120	region不能为空
N62121	result不能为空
N62122	bizOrderId不能为空
N62123	bizType不能为空
N62124	merchantId不能为空
N62125	notifyUrl不能为空
N62126	缓存中不包含此商户appId
N62127	缓存中不包含此backendToken
N62128	接口名不能为空
N62129	relateId不能为空
N62130	trId不能为空
N62131	无权限访问此接口（请检查授权联登上送scope是否正确，咨询业务人员是否具有该接口调用权限）
N62132	不支持此访问域名（请检查申请对接云闪付开放平台的申请表中配置的域名与授权联登请求链接中的redirect_url是否一致）
N62134	退税号不能为空
N62135	用户授权时，没有提供手机号
N62136	用户没有手机号
N62137	获取签约商户信息异常（检查上送的appId和planId）
N62138	验证支付密码异常
N62142	待更新的缓存key不能为空
N62143	待更新的缓存expire不能为空
S52131	无权限访问此接口（请检查授权联登上送scope是否正确，咨询业务人员是否具有该接口调用权限）
S52136	用户没有手机号
S52172	接入方编码非法（请检查活动id/机构账户号是否和申请单填写一致）

*/

var (
	E00   = ErrNew("00", "ok")
	EA10  = ErrNew("a10", "不合法的backend_token，或已过期")
	EA20  = ErrNew("a20", "不合法的frontend_token，或已过期")
	E3023 = ErrNew("3023", "用户未注册云闪付APP")
)

type Err struct {
	Code string `json:"errno"`
	Msg  string `json:"error"`
}

func ErrNew(code, msg string) *Err {
	return &Err{Code: code, Msg: msg}
}

func NewF(err error) *Err {
	e := err.Error()
	i := strings.Index(e, ":")
	if i == -1 {
		return &Err{Code: "UPERR", Msg: e}
	}
	return &Err{Code: e[0:i], Msg: e[i+1:]}
}

func (e *Err) NewMsg(msg string) *Err {
	return ErrNew(e.Code, msg)
}

func (e *Err) NewMsgF(args ...interface{}) *Err {
	return ErrNew(e.Code, fmt.Sprintf(e.Msg, args...))
}

func (e *Err) Error() string {
	return e.Msg + ".[" + e.Code + "]"
}

func (e *Err) JsonMarshal() []byte {
	b, _ := json.Marshal(e)
	return b
}
