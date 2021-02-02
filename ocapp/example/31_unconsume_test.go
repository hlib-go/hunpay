package main

import (
	"encoding/json"
	"fmt"
	"github.com/hlib-go/hunpay/ocapp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

// 消费撤销未测试成功

// 消费撤销
func TestUnConsume(t *testing.T) {
	result, err := ocapp.UnConsume(cfg, &ocapp.UnConsumeParams{
		OrigQryId:   "232102011314361711318",
		OrderId:     "135610866920014233615",
		TxnAmt:      "1",
		BackUrl:     "https://msd.himkt.cn/work/unconsume/notify",
		TxnTime:     ocapp.TxnTime(),
		ReqReserved: "",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	rbytes, _ := json.Marshal(result)
	t.Log(string(rbytes))
}

// 消费撤销通知测试  https://msd.himkt.cn/work/unconsume/notify
func TestUnConsumeNotify(t *testing.T) {
	http.Handle("/unconsume/notify", ocapp.UnConsumeNotifyHandler(func(o *ocapp.UnConsumeNotifyEntity) error {
		rbytes, _ := json.Marshal(o)
		fmt.Println("收到消费撤销通知结果JSON：")
		fmt.Println(rbytes)
		return nil
	}))
	fmt.Println("Start serve Listen 80 ...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), nil))
}
