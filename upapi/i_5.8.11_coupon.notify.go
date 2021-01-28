package upapi

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 优惠券变动状态通知
func CouponNotifyHandler(getConfig func(appid string) (cfg *Config, err error), cbFn func(req *http.Request, requestId string, notifyResult *CouponNotify) error) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var (
			requestId    = Rand32()
			plog         = logrus.WithField("requestId", requestId)
			err          error
			notifyResult *CouponNotify
		)
		plog.Info("闪券通知 RequestURI ", req.RequestURI)
		defer func() {
			if err != nil {
				plog.Error(err.Error())
				res.WriteHeader(500)
				res.Write([]byte(`{"error":"` + err.Error() + `","requestId":"` + requestId + `"}`))
				return
			}
			err = cbFn(req, requestId, notifyResult)
			if err != nil {
				plog.Error(err.Error())
				res.Write([]byte(`{"error":"` + err.Error() + `","requestId":"` + requestId + `"}`))
				res.WriteHeader(500)
				return
			}
			plog.Info("通知处理成功", ` {"code":"00","requestId":"`+requestId+`"}`)
			// 响应处理成功
			res.WriteHeader(200)
			res.Write([]byte(`{"code":"00","error":"ok","requestId":"` + requestId + `"}`))
		}()

		err = req.ParseForm()
		if err != nil {
			return
		}
		form := req.Form

		bmap := NewBodyMap()
		for k, v := range form {
			if len(v) > 0 {
				bmap.Set(k, v[0])
			}
		}
		signature := bmap.Get("signature")
		appid := bmap.Get("appId")
		plog = plog.WithField("discountId", bmap.Get("discountId")).WithField("appid", appid)
		if appid == "" {
			err = errors.New("非法请求，无效appid")
			return
		}
		if signature == "" {
			err = errors.New("非法请求，无效signature")
			return
		}
		pbytes, err := bmap.MarshalJSON()
		if err != nil {
			return
		}
		plog.Info("闪券通知JSON格式报文：", string(pbytes))

		// 根据appid获取银联公钥
		cfg, err := getConfig(appid)
		if err != nil {
			return
		}
		if cfg == nil {
			err = errors.New("回调通知根据appid未读取到云闪付配置")
			return
		}

		// 验证签名
		err = UpRsaVerify(signature, bmap, cfg.UpPublicKey, true)
		if err != nil {
			return
		}

		err = json.Unmarshal(pbytes, &notifyResult)
		if err != nil {
			return
		}
	})
}

type CouponNotify struct {
	AppId        string  `json:"appId"`
	NonceStr     string  `json:"nonceStr"`
	Timestamp    string  `json:"timestamp"`
	TransSeqId   string  `json:"transSeqId"`   // 交易流水号，唯一
	TransTs      string  `json:"transTs"`      // 交易时间
	TraceId      string  `json:"traceId"`      // 跟踪号(原交易流水号) ，根据此交易号更新券状态
	DiscountId   string  `json:"discountId"`   // 优惠券ID
	DiscountName string  `json:"discountName"` // 优惠券名称
	DiscountNum  string  `json:"discountNum"`  //优惠券数量
	EntityTp     string  `json:"entityTp"`     //主体类型(2位，：用户、卡号、 手机号（三选一）01-手机号 02- 卡号 03-用户)
	EntityId     string  `json:"entityId"`     // 用户标识(03:代表openId,02:代 表卡号,01:代表手机号,其中,卡 号,手机号,为加密返回,不可逆转)
	OperaTp      OperaTp `json:"operaTp"`      // 操作类型,根据操作更新券状态
	TransTp      TransTp `json:"transTp"`      // 交易号
	MchntCd      string  `json:"mchntCd"`      // 商户编号
	PosTmn       string  `json:"posTmn"`
	OrderNo      string  `json:"orderNo"`
	OrderAt      string  `json:"orderAt"`
	DiscountAt   string  `json:"discountAt"`
	TransChnl    string  `json:"transChnl"`
}

func (c *CouponNotify) JsonString() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type OperaTp string

const (
	OPERA_TP_01 OperaTp = "01" // 获取优惠券
	OPERA_TP_02 OperaTp = "02" // 核销优惠券
	OPERA_TP_03 OperaTp = "03" // 返还优惠券
	OPERA_TP_04 OperaTp = "04" // 优惠券无操作（如：部分退货等）
	OPERA_TP_05 OperaTp = "05" // 优惠券删除
	OPERA_TP_06 OperaTp = "06" // 优惠券过期
)

type TransTp string

const (
	TRANS_TP_01 TransTp = "01" // 消费
	TRANS_TP_31 TransTp = "31" // 撤销
	TRANS_TP_04 TransTp = "04" // 退货
)
