package upapi

import (
	"time"
)

/*
5.8.10  赠送优惠券结果查询 <coupon.query>
*/
func CouponQuery(c *Config, transSeqId, origTransSeqId, origTransTs, backendToken string) (err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.AppId)
	bm.Set("origTransSeqId", origTransSeqId)
	bm.Set("origTransTs", origTransTs)
	bm.Set("transSeqId", transSeqId) // 交易ID，可以用uuid
	bm.Set("transTs", time.Now().Format("20060102"))
	bm.Set("backendToken", backendToken)

	resp, err := Post(c, "/coupon.query", bm)
	if err != nil {
		return
	}
	if resp.Resp != E00.Code {
		err = ErrNew(resp.Resp, resp.Msg)
		return
	}
	return
}
