package upapi

import (
	"encoding/json"
	"strconv"
	"time"
)

/*
云闪付RSA Sha256 签名问题记录：
1. 非必填字段，需要参与签名，内容为空字符串。
2. 敏感字段需要在签名后，再进行加密处理。
*/

// 抽奖（红包/票券）<qual.reduce>
func QualReduce(c *Config, transNumber, activityNumber, qualNum, qualType, qualValue string) (r *QualReduceResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.AppId)                 // 是  接入方的唯一标识
	bm.Set("activityNumber", activityNumber) // 是	活动 编号
	bm.Set("orderAmount", "")                //否	订单金额
	bm.Set("qualNum", qualNum)               // 是	资格池编号
	bm.Set("qualType", qualType)             // 是	资格类型   固定值“open_id”、“mobile”、“card_no”
	bm.Set("qualValue", qualValue)           // 是	资格值 （使用symmetricKey 对称加密，内容为base64格式）
	bm.Set("certId", "")                     // 否	身份证号 （使用symmetricKey 对称加密，内容为base64格式）
	bm.Set("icTerminal", "")                 // 否	设备终端
	bm.Set("qrCode", "")                     // 否	红包码
	bm.Set("transNumber", transNumber)       // 是	流水号（确保唯一）
	bm.Set("nonceStr", GetRandomString(16))  // 是	生成签名的随机串
	bm.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	signature, err := UpRsaSign(bm, c.MchPrivateKey, true)
	if err != nil {
		return
	}
	bm.Set("signature", signature)

	// 注意：签名之后再加密敏感字段
	qualValueEncode, err := Encode3DES(c.SymmetricKey, qualValue)
	if err != nil {
		return
	}
	bm.Set("qualValue", qualValueEncode)

	resp, err := Post(c, "/qual.reduce", bm)
	if err != nil {
		return
	}
	if resp.Resp != E00.Code {
		err = ErrNew(resp.Resp, resp.Msg)
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

type QualReduceResult struct {
	RespCode    string                     `json:"respCode"`
	RespTime    string                     `json:"respTime"`
	TransNumber string                     `json:"transNumber"`
	AwardInfo   *QualReduceResultAwardInfo `json:"awardInfo"`
}

type QualReduceResultAwardInfo struct {
	ActivityNumber  string `json:"activityNumber"`
	ActivityName    string `json:"activityName"`
	BeginTime       string `json:"beginTime"`
	EndTime         string `json:"endTime"`
	AwardId         string `json:"awardId"`
	AwardType       string `json:"awardType"`
	AwardName       string `json:"awardName"`
	ExtAcctId       string `json:"extAcctId"`
	ExtAcctName     string `json:"extAcctName"`
	DrawDesc        string `json:"drawDesc"`
	CouponStartDate string `json:"couponStartDate"`
	CouponEndDate   string `json:"couponEndDate"`
	CouponGoodsUrl  string `json:"couponGoodsUrl"`
}

/*

序号	参数名	是否必填	备注
1	appId	是	接入方的唯一标识
2	activityNumber	是	活动 编号
3	orderAmount	否	订单金额
4	qualNum	是	资格池编号
5	qualType	是	资格类型
固定值“open_id”、“mobile”、“card_no”
6	qualValue	是	资格值
（使用symmetricKey 对称加密，内容为base64格式）
7	certId	否	身份证号
（使用symmetricKey 对称加密，内容为base64格式）
8	icTerminal	否	设备终端
9	qrCode	否	红包码
10	transNumber	是	流水号（确保唯一）
11	nonceStr	是	生成签名的随机串
12	timestamp	是	生成签名的时间戳
13	signature	是	请使用接入方私钥签名，输出格式为base64.

*/
