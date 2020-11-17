package main

import (
	"fmt"
	"github.com/hlib-go/hunpay/ocwap"
	"log"
	"net/http"
)

// 消费测试
func main() {
	// https://msd.himkt.cn/work/consume?orderId=T000101&txnAmt=1&accNo=
	http.HandleFunc("/consume", func(writer http.ResponseWriter, request *http.Request) {
		// 跳转银联全渠道手机网页支付界面
		err := ocwap.Consume(cfg, &ocwap.ConsumeParams{
			AccNo:       request.FormValue("accNo"),
			OrderId:     request.FormValue("orderId"),
			TxnAmt:      request.FormValue("txnAmt"),
			FrontUrl:    "https://msd.himkt.cn/work/consume/front",
			BackUrl:     "https://msd.himkt.cn/work/consume/back",
			TxnTime:     ocwap.TxnTime(),
			ReqReserved: "--",
		}, writer)
		if err != nil {
			return
		}
	})
	http.HandleFunc("/consume/front", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("前端通知：", request.RequestURI)
		writer.Write([]byte(request.RequestURI))
	})
	http.Handle("/consume/back", ocwap.ConsumeNotifyHandler(func(o *ocwap.ConsumeNotifyEntity) error {

		return nil
	}))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), nil))
}
