package upapi

import (
	"strconv"
	"time"
)

// 直接抽奖
func MaktgDraw(c *Config, p *MaktgDrawParams) (err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.AppId)           //是 接入方的唯一标识
	bm.Set("transSeqId", p.TransSeqId) //是 交易流水号,不重复，最大64位
	bm.Set("transTs", p.TransTs)       //是 请求日期, 格式yyyyMMdd，如：20191227
	bm.Set("activityNo", p.ActivityNo)
	bm.Set("mobile", p.Mobile)
	bm.Set("cardNo", p.CardNo)              //否 交易卡号，若上送，（使用symmetricKey 对称加密，内容为 base64格式）
	bm.Set("openId", p.OpenId)              //否 用户唯一标识
	bm.Set("acctEntityTp", "03")            // 01-手机号 02-卡号 03-用户(手机、卡号、openid三选一)
	bm.Set("nonceStr", GetRandomString(16)) // 是	生成签名的随机串
	bm.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	signature, err := UpRsaSign(bm, c.MchPrivateKey, false)
	if err != nil {
		return
	}
	bm.Set("signature", signature)
	if p.Mobile != "" {
		// 注意：签名之后再加密敏感字段
		var mobile string
		mobile, err = Encode3DES(c.SymmetricKey, p.Mobile)
		if err != nil {
			return
		}
		bm.Set("mobile", mobile)
	}

	resp, err := Post(c, "/maktg.draw", bm)
	if err != nil {
		return
	}
	if resp.Resp != E00.Code {
		err = ErrNew(resp.Resp, resp.Msg)
		return
	}
	return
}

type MaktgDrawParams struct {
	TransSeqId string // 交易流水
	TransTs    string // 请求日期
	ActivityNo string // 活动编号
	Mobile     string //
	CardNo     string
	OpenId     string
}
