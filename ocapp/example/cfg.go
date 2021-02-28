package main

import (
	"fmt"
	"github.com/hlib-go/hunpay/ocapp"
	"github.com/hlib-go/hunpay/upsdk"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

// 测试配置
var cfg = ocapp.NewConfig("https://gateway.95516.com", "821330248164056", "", "81628889475")

func init() {
	// 读取商户私钥, 此文件由商户通过银联平台下载的pfx证书导出，对应的公钥通过银联商户平台上传到银联
	bytes, err := ioutil.ReadFile("C:\\www\\certs\\himkt-unionpay.key")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	cfg.MerPrivateKey = string(bytes)
}

// config
/*
	"upServiceUrl" : "https://open.95516.com/open/access/1.0",
	"upAppId" : "50f9fdbb53b64209ae3690d22dff778e",
	"upSecret" : "127a5ec59b2241bea8df8f0d784fb403",
	"upSymmetricKey" : "20a140eabc4343f2c8cd62549e7a0bf420a140eabc4343f2",
*/
var cfgApp = upsdk.New(&upsdk.Config{
	/*BaseServiceUrl: "https://open.95516.com/open/access/1.0",
	AppId:          "", // appId关联了域名 ms.himkt.cn, 使用其它域名无法使用
	Secret:         "127a5ec59b2241bea8df8f0d784fb403",
	SymmetricKey:   "20a140eabc4343f2c8cd62549e7a0bf420a140eabc4343f2",
	UpPublicKey:    "",
	MchPrivateKey:  ``,*/
})

func init() {
	cfgApp.GetFrontToken = func(refresh bool) string {
		r, err := cfgApp.FrontToken()
		if err != nil {
			log.Error(err)
			return ""
		}
		return r.FrontToken
	}
}
