package main

import (
	"encoding/json"
	"fmt"
	"github.com/hlib-go/hunpay/ocapp"
	"github.com/hlib-go/hunpay/ocwap"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// 消费测试
func main() {
	// https://msd.himkt.cn/work/consume?orderId=T0000001&txnAmt=1&accNo=6251211100976741
	// https://ms.himkt.cn/mswork/consume?orderId=T00000021112&txnAmt=1&accNo=6214830213065526   ，云闪付扫码方式访问此链接，直接调起控件支付
	// 消费，跳转云闪付控件支付
	http.HandleFunc("/consume", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html;charset=utf-8")
		var err error
		defer func() {
			if e := recover(); e != nil {
				err = e.(error)
			}

			if err != nil {
				writer.Write([]byte("Error: " + err.Error()))
				return
			}
		}()

		txnAmt, _ := strconv.ParseInt(request.FormValue("txnAmt"), 10, 64)

		// 跳转银联全渠道手机网页支付界面
		result, err := ocapp.Consume(cfg821330248164060, &ocapp.ConsumeParams{
			AccNo:       request.FormValue("accNo"),
			OrderId:     request.FormValue("orderId"),
			TxnAmt:      txnAmt,
			BackUrl:     "https://ms.himkt.cn/mswork/consume/back",
			TxnTime:     ocwap.TxnTime(),
			ReqReserved: "-",
		})
		if err != nil {
			return
		}
		// 获取TN，前端调用控件支付
		tn := result.Tn

		fmt.Println("RU-> ", "https://ms.himkt.cn/mswork"+request.RequestURI)
		bm := cfgApp.UpsdkConfig("https://ms.himkt.cn/mswork"+request.RequestURI, true)

		p := make(map[string]string)
		p["tn"] = tn
		p["appId"] = bm.Get("appId")
		p["nonceStr"] = bm.Get("nonceStr")
		p["signature"] = bm.Get("signature")
		p["timestamp"] = bm.Get("timestamp")

		html := `
<html>
<head>
	<title>云闪付控件支付</title>
    <script src="https://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
    <!-- 云闪付UPSDK,需要依赖jquery -->
    <script src="https://open.95516.com/s/open/js/upsdk.js"></script>
	<script>
		$(function(){
			console.log("ocapp pay ......");
			window.upsdk.config({
				appId: '{{.appId}}',
				nonceStr: '{{.nonceStr}}',
				signature: '{{.signature}}',
				timestamp: {{.timestamp}},
				debug: true,
			});
			window.upsdk.ready(function () {
				//config信息验证后会执行ready方法
				alert("upsdk config success");
				console.log("upsdk config success...");
			});
			window.upsdk.error(function (err) {
				//config信息验证失败会执行error方法
				alert("upsdk config error -> "+err.message)
				console.log("upsdk config error:", err);
			});
			$("#pay").click(function(){
   				if (!window.upsdk){
					alert("upsdk未加载成功")
					return;
				}
				window.upsdk.pay({
				tn: '{{.tn}}',
				success: function () {
				  alert("支付成功")
				},
				fail: function () {
				  alert("支付失败")
				},
			  });
			});
		});
	</script>
</head>
<body>
云闪付控件支付测试
<hr>
<button id=pay>去支付</button>
</body>
</html>
		`

		tpl := template.New("consume.tpl")
		if err != nil {
			return
		}
		tpl, err = tpl.Parse(html)
		if err != nil {
			return
		}
		tpl.Execute(writer, p)
	})

	// 接受通知，判断状态00，调用查询接口，oriRespCode等于00，执行业务发货逻辑
	http.Handle("/consume/back", ocapp.ConsumeNotifyHandler(func(o *ocapp.ConsumeNotifyEntity) error {
		bytes, _ := json.Marshal(o)
		fmt.Println("JSON ConsumeNotifyEntity ", string(bytes))
		return nil
	}))

	// https://msd.himkt.cn/work/query?orderId=T0000002
	http.HandleFunc("/query", func(writer http.ResponseWriter, request *http.Request) {
		orderId := request.FormValue("orderId")
		result, err := ocapp.Query(cfg, orderId, "")
		if err != nil {
			fmt.Println("Query Error", err.Error())
			return
		}
		bytes, _ := json.Marshal(result)
		fmt.Println("resultJSON", string(bytes))
		writer.Write(bytes)
	})

	fmt.Println("Start serve Listen 80 ...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 80), nil))
}
