package ocwap

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"log"
	"testing"
)

// 解析 pfx 证书私钥与序列号
func TestPfxToBase64(t *testing.T) {
	//fbytes,_:= ioutil.ReadFile("D:\\Information\\Company\\中国银联\\手机网页支付\\手机网页（WAP）支付产品技术开发包1.1.8\\Java Version SDK (通用版)\\ACPSample_B2C\\src\\assets\\测试环境证书\\acp_test_sign.pfx")
	//fbytes, _ := ioutil.ReadFile("D:\\Projects\\himkt-go\\hm-unionpay\\sdk\\cert\\acptest\\acp_test_sign.pfx")
	fbytes, _ := ioutil.ReadFile("D:\\Information\\Company\\中国银联\\海脉云\\导出证书\\himkt-unionpay.pfx")
	pass := "142857"
	value := "123"

	priKey, cert, err := pkcs12.Decode(fbytes, pass)
	if err != nil {
		t.Error(err)
		return
	}

	// 格式化私钥
	derStream, _ := x509.MarshalPKCS8PrivateKey(priKey)
	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	})

	// 计算签名
	sign, err := RsaWithSha256Sign(value, string(pemBytes))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("*************** 解析银联pfx证书私钥与序列号 ***************")
	fmt.Println("SerialNumber=" + cert.SerialNumber.String())
	fmt.Println("私钥证书 \n" + string(pemBytes))
	fmt.Println("value=" + value + " sign=" + sign)

}

// 解析 pfx 证书私钥与序列号
func TestPfxToPem(t *testing.T) {
	//fbytes, _ := ioutil.ReadFile("D:\\Information\\Company\\中国银联\\海脉云\\导出证书\\himkt-unionpay.pfx")
	fbytes, _ := ioutil.ReadFile("D:\\Projects\\himkt-go\\hm-chunpay\\sdk\\cert\\MER777290058185160\\acp_test_sign.pfx")
	pass := "000000"
	value := "123"

	priKey, cert, err := pkcs12DecodeAll(fbytes, pass)
	if err != nil {
		t.Error(err)
		return
	}

	// 格式化私钥
	derStream, _ := x509.MarshalPKCS8PrivateKey(priKey[0])
	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	})

	// 计算签名
	sign, err := RsaWithSha256Sign(value, string(pemBytes))
	if err != nil {
		t.Error(err)
		return
	}

	// 81628924165
	// 405099140
	// 68719476761
	fmt.Println("*************** 解析银联pfx证书私钥与序列号 ***************")
	fmt.Println("SerialNumber=" + cert[0].SerialNumber.String())
	fmt.Println("私钥证书 \n" + string(pemBytes))
	fmt.Println("value=" + value + " sign=" + sign)
}

// pkcs12DecodeAll extracts all certificate and private keys from pfxData.
func pkcs12DecodeAll(pfxData []byte, password string) ([]interface{}, []*x509.Certificate, error) {
	var privateKeys []interface{}
	var certificates []*x509.Certificate

	blocks, err := pkcs12.ToPEM(pfxData, password)
	if err != nil {
		log.Printf("error while converting to PEM: %s", err)
		return nil, nil, err
	}

	for _, b := range blocks {
		if b.Type == "CERTIFICATE" {
			certs, err := x509.ParseCertificates(b.Bytes)
			if err != nil {
				return nil, nil, err
			}
			certificates = append(certificates, certs...)

		} else if b.Type == "PRIVATE KEY" {
			privateKey, err := x509.ParsePKCS1PrivateKey(b.Bytes)
			if err != nil {
				return nil, nil, err
			}
			privateKeys = append(privateKeys, privateKey)
		}
	}
	return privateKeys, certificates, err
}
