package main

import (
	"encoding/json"
	"fmt"
	"github.com/hlib-go/hunpay/ocapple"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestRefund(t *testing.T) {
	r, err := ocapple.Refund(cfg, &ocapple.RefundParams{
		OrigQryId: "212102011124137231278",
		OrderId:   "135610866920014233616", //"1356108669200142336"
		TxnAmt:    "1",
		BackUrl:   "https://msd.himkt.cn/work/refund/notify",
		TxnTime:   ocapple.TxnTime(),
	})
	t.Log(r)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

// 消费撤销通知测试  https://msd.himkt.cn/work/refund/notify
func TestRefundNotify(t *testing.T) {
	http.Handle("/refund/notify", ocapple.RefundNotifyHandler(func(o *ocapple.RefundNotifyEntity) error {
		rbytes, _ := json.Marshal(o)
		fmt.Println("收到消费退款通知结果JSON：")
		fmt.Println(string(rbytes))
		return nil
	}))
	fmt.Println("Start serve Listen 80 ...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), nil))
}
