package upsdk

import (
	"time"
)

/*
5.8.10  赠送优惠券结果查询 <coupon.query>
*/
func (up *Unionpay) CouponQuery(transSeqId, origTransSeqId, origTransTs string) (err error) {
	bm := NewBodyMap()
	bm.Set("appId", up.Config.AppId)
	bm.Set("origTransSeqId", origTransSeqId)
	bm.Set("origTransTs", origTransTs)
	bm.Set("transSeqId", transSeqId) // 交易ID，可以用uuid
	bm.Set("transTs", time.Now().Format("20060102"))
	bm.Set("backendToken", up.GetBackendToken(false))

	resp, err := up.Post("/coupon.query", bm)
	if err != nil {
		return
	}
	if resp.Resp != E00.Code {
		err = ErrNew(resp.Resp, resp.Msg)
		return
	}
	return
}
