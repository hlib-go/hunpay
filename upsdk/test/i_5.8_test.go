package test

import (
	"encoding/json"
	"github.com/hlib-go/hunpay/upsdk"
	"testing"
	"time"
)

// 5.8.6  抽奖（红包/票券）<qual.reduce>
func TestQualReduce(t *testing.T) {
	/*
	  "mchId": "100002",
	  "transNumber":"423456789039",
	  "qualNum":"3a5d4792-48c3-4416-809d-ada5da535f84",
	  "qualType": "mobile",
	  "qualValue":"13912300661",
	  "activityNumber":"1320200615282465"*/
	_, err := c.QualReduce("423456789039", "1320200615282465", "3a5d4792-48c3-4416-809d-ada5da535f84", "mobile", "13611703040")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("TestQualReduce success..................")
}

/*
测试活动ID：
云闪付删券活动ID: 1320200615282456   couponId:3102020072729846


活动ID1：1320200615282427
活动ID2：1320200615282448

        外部活动ID           之前给的立减活动ID（编码非法）     现在给的立减活动ID
满减券1  1320200615282456   3102020072429626            3102020072729846 （测试可以用）
满减券2  1320200615282465   3102020072429625            3102020072429635
*/
// 5.8.9  赠送优惠券 <coupon.download>
func Test_CouponDownload(t *testing.T) {
	r, err := c.CouponDownload(&upsdk.CouponDownloadParams{
		TransSeqId: upsdk.Rand32(),
		TransTs:    time.Now().Format("20060102"),
		CouponId:   "3102020103047596",
		CouponNum:  1,
		Mobile:     "15657477628",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	b, _ := json.Marshal(r)
	t.Log(string(b))
	/*
		尝试了之前31和13开头的活动ID，结果如下：
		13开头的ID，赠送优惠券接口返回：The coupon activity is not exist.[GCUP06003]
		31开头的ID，赠送优惠券接口返回：接入方编码非法[S52172]
	*/
}

// 5.8.10  赠送优惠券结果查询 <coupon.query>
func Test_CouponQuery(t *testing.T) {
	//"transSeqId":"4ff007f90a384eee869d99d7166ed342","transTs":"20200906"
	c.GetBackendToken = func(refresh bool) string {
		t, _ := c.BackendToken()
		return t.BackendToken
	}
	err := c.CouponQuery(upsdk.Rand32(), "2922b52272a04cc9a6cec290c2a6d324", "20201015")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

// 5.8.12  优惠券活动剩余名额查询
func Test_ActivityQuota(t *testing.T) {
	_, err := c.ActivityQuota(upsdk.Rand32(), "3102020072729846", "5")
	if err != nil {
		t.Error(err.Error())
		return
	}
	// 测试结果：404 Not Found
}
