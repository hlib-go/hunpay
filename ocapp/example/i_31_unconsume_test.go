package main

import (
	"encoding/json"
	"fmt"
	"github.com/hlib-go/hunpay/ocapp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

// 消费撤销
func TestUnConsume(t *testing.T) {
	result, err := ocapp.UnConsume(cfg, &ocapp.UnConsumeParams{
		OrigQryId:   "232102011314361711318",
		OrderId:     "1356108669200142336",
		TxnAmt:      "1",
		BackUrl:     "https://msd.himkt.cn/work/consume/back",
		TxnTime:     "20210201131436",
		ReqReserved: "",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	rbytes, _ := json.Marshal(result)
	t.Log(string(rbytes))
}

func TestUnConsumeNotify(t *testing.T) {
	// 消费撤销通知
	http.Handle("/unconsume/notify", ocapp.UnConsumeNotifyHandler(func(o *ocapp.UnConsumeNotifyEntity) error {
		return nil
	}))
	fmt.Println("Start serve Listen 80 ...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), nil))
}
