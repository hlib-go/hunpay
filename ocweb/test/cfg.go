package main

import (
	"fmt"
	"github.com/hlib-go/hunpay/ocweb"
	"io/ioutil"
)

// 测试配置
var cfg = ocweb.NewConfig("https://gateway.95516.com", "821330248164056", "", "81628889475")

func init() {
	// 读取商户私钥, 此文件由商户通过银联平台下载的pfx证书导出，对应的公钥通过银联商户平台上传到银联
	bytes, err := ioutil.ReadFile("C:\\www\\certs\\himkt-unionpay.key")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	cfg.MerPrivateKey = string(bytes)
}
