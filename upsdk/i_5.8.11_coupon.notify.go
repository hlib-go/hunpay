package upsdk

import (
	"fmt"
	"github.com/hlib-go/htype"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// 优惠券变动状态通知
func CouponNotifyHandler(cbFn func(req *http.Request, requestId string, notifyResult *CouponNotify) error) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var (
			requestId    = Rand32()
			plog         = logrus.WithField("requestId", requestId)
			err          error
			notifyResult *CouponNotify
		)
		defer func() {
			if err != nil {
				res.Write([]byte(err.Error()))
				return
			}
			err = cbFn(req, requestId, notifyResult)
			if err != nil {
				res.Write([]byte(err.Error()))
				return
			}
			// 响应处理成功
			res.WriteHeader(200)
			res.Write([]byte(`{"code":"00"}`))
		}()

		//接收参数
		paramsBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return
		}
		fmt.Println("CouponNotifyHandler:" + string(paramsBytes))
		plog.Info("CouponNotifyHandler:" + string(paramsBytes))

		// 验证签名

		// 返回结果

	})
}

type CouponNotify struct {
	AppId        string      `json:"appId"`
	TransSeqId   string      `json:"transSeqId"`
	TransTs      string      `json:"transTs"`
	TraceId      string      `json:"traceId"`
	DiscountId   string      `json:"discountId"`   // 优惠券ID
	DiscountName string      `json:"discountName"` // 优惠券名称
	NonceStr     string      `json:"nonceStr"`
	DiscountNum  htype.Int64 `json:"discountNum"`
	EntityTp     string      `json:"entityTp"`
	EntityId     string      `json:"entityId"`
	Timestamp    string      `json:"timestamp"`
	OperaTp      OperaTp     `json:"operaTp"` // 操作类型
	TransTp      TransTp     `json:"transTp"`
	MchntCd      string      `json:"mchntCd"`
	PosTmn       string      `json:"posTmn"`
	OrderNo      string      `json:"orderNo"`
	OrderAt      htype.Int64 `json:"orderAt"`
	DiscountAt   htype.Int64 `json:"discountAt"`
	TransChnl    string      `json:"trans_chnl"`
}

type OperaTp string

const (
	OPERA_TP_01 OperaTp = "01" // 获取优惠券
	OPERA_TP_02 OperaTp = "02" // 核销优惠券
	OPERA_TP_03 OperaTp = "03" // 返还优惠券
	OPERA_TP_04 OperaTp = "04" // 优惠券无操作（如：部分退货等）
	OPERA_TP_05 OperaTp = "05" // 优惠券删除
	OPERA_TP_06 OperaTp = "05" // 优惠券过期
)

type TransTp string

const (
	TRANS_TP_01 TransTp = "01" // 消费
	TRANS_TP_31 TransTp = "31" // 撤销
	TRANS_TP_04 TransTp = "04" // 退货
)
