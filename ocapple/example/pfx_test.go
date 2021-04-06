package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/hlib-go/hunpay/ocwap"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"log"
	"testing"
)

// 解析 pfx 证书私钥与序列号
func TestPfxToBase64(t *testing.T) {
	fbytes, _ := ioutil.ReadFile("D:\\certs\\unionpay-821330248164060.pfx")
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
	sign, err := ocwap.RsaWithSha256Sign(value, string(pemBytes))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("*************** 解析银联pfx证书私钥与序列号 ***************")
	fmt.Println("SerialNumber=" + cert.SerialNumber.String())
	fmt.Println("私钥证书 \n" + string(pemBytes))
	fmt.Println("value=" + value + " sign=" + sign)

}

// 解析 pfx 证书私钥与序列号  ,输出pem格式
func TestPfxToPem(t *testing.T) {
	fbytes, _ := ioutil.ReadFile("D:\\certs\\unionpay-821330248164060.pfx")
	pass := "142857"
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
	sign, err := ocwap.RsaWithSha256Sign(value, string(pemBytes))
	if err != nil {
		t.Error(err)
		return
	}

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
