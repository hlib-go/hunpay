package main

import (
	"fmt"
	"github.com/hlib-go/hunpay/ocapple"
	"io/ioutil"
)

// 测试配置
var cfg821330248164060 = ocapple.NewConfig("https://gateway.95516.com", "821330248164060", "", "82031637015")

func init() {
	// 读取商户私钥, 此文件由商户通过银联平台下载的pfx证书导出，对应的公钥通过银联商户平台上传到银联
	bytes, err := ioutil.ReadFile("C:\\www\\certs\\unionpay-821330248164060.key")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	cfg821330248164060.MerPrivateKey = string(bytes)
}
