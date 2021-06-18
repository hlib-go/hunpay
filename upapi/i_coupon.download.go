package upapi

import (
	"encoding/json"
	"strconv"
	"time"
)

/*
5.8.9  赠送优惠券 <coupon.download>
*/
func CouponDownload(c *Config, p *CouponDownloadParams) (r *CouponDownloadResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.AppId)                //是 接入方的唯一标识
	bm.Set("transSeqId", p.TransSeqId)      //是 交易流水号,不重复，最大64位
	bm.Set("transTs", p.TransTs)            //是 请求日期, 格式yyyyMMdd，如：20191227
	bm.Set("couponId", p.CouponId)          //是 优惠券id
	bm.Set("mobile", p.Mobile)              //否 交易手机号，若上送，（使用symmetricKey 对称加密，内容 为base64格式）
	bm.Set("cardNo", p.CardNo)              //否 交易卡号，若上送，（使用symmetricKey 对称加密，内容为 base64格式）
	bm.Set("openId", p.OpenId)              //否 用户唯一标识
	bm.Set("acctEntityTp", "03")            //是 营销活动配置的赠送维度（参见营销平台活动配置），2位， 可选： 02-卡号 03-用户（二选一） 赠送维度为卡号时，则cardNo必填； 赠送维度为用户时，则openId，mobile, cardNo三选一上送
	bm.Set("couponNum", p.CouponNum)        //是  优惠券数量
	bm.Set("nonceStr", GetRandomString(16)) //是	生成签名的随机串
	bm.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	signature, err := UpRsaSign(bm, c.MchPrivateKey, false)
	if err != nil {
		return
	}
	bm.Set("signature", signature)

	if p.Mobile != "" {
		// 注意：签名之后再加密敏感字段
		mobile, err := Encode3DES(c.SymmetricKey, p.Mobile)
		if err != nil {
			return nil, err
		}
		bm.Set("mobile", mobile)
	}
	if p.CardNo != "" {
		// 注意：签名之后再加密敏感字段
		cardNo, err := Encode3DES(c.SymmetricKey, p.CardNo)
		if err != nil {
			return nil, err
		}
		bm.Set("cardNo", cardNo)
	}

	resp, err := Post(c, "/coupon.download", bm)
	if err != nil {
		return
	}
	switch resp.Resp {
	case E3023.Code:
		// 未注册云闪付APP
		err = E3023
	case "GCUP06038":
		// Coupon download failed due to useId is limited.[GCUP06038]
		err = ErrNew(resp.Resp, "超过限制次数")
	case "GCUP06007":
		// Coupon download rules match failed.[GCUP06007]
		err = ErrNew(resp.Resp, "规则验证失败，请确认在有效时间内")
	case "GCUP07056":
		//Coupon download failed due to cityId is invalid.[GCUP07056]
		err = ErrNew(resp.Resp, "Coupon download failed due to cityId is invalid")
	case "GCUP07053":
		// Coupon download failed due to no cardNo is invalid.[GCUP07053]
		err = ErrNew(resp.Resp, "未绑定活动指定银行卡，请先去云闪付绑卡")
	case "GCUP06045":
		// Coupon download failed due to there has no coupon left.[GCUP06045]
		err = ErrNew(resp.Resp, "该优惠券已达到领用上限")
	case "GCUP07052":
		// Coupon download failed due to acct is not samename.[GCUP07052]
		err = ErrNew(resp.Resp, "云闪付没有绑卡或所绑卡不同名")
	case "GCUP06036":
		// Coupon download failed due to cardNo is limited.[GCUP06036]
		err = ErrNew(resp.Resp, "卡号受限无法领券")
	case "GCUP07058":
		//Coupon download failed due to yellowNameList check failed.[GCUP07058]
		err = ErrNew(resp.Resp, "手机号【"+p.Mobile+"】异常-yellowNameList，请咨询银联客服")
	case "GCUP07060":
		// Coupon download failed due to userId check failed.[GCUP07060]
		err = ErrNew(resp.Resp, "您不符合领券要求，请查看活动说明")
	default:
		if resp.Resp != E00.Code {
			err = ErrNew(resp.Resp, resp.Msg)
		}
	}
	if err != nil {
		return
	}

	// 解析响应报文
	pbytes, err := json.Marshal(resp.Params)
	if err != nil {
		return
	}
	err = json.Unmarshal(pbytes, &r)
	if err != nil {
		return
	}
	return
}

type CouponDownloadParams struct {
	TransSeqId string // 交易流水
	TransTs    string // 请求日期
	CouponId   string
	CouponNum  int64
	Mobile     string //Mobile CardNo OpenId 三选一
	CardNo     string
	OpenId     string
}

type CouponDownloadResult struct {
	TransSeqId string
	CouponId   string
}
