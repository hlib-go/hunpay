package ocapp

import (
	"crypto"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

/*
银联接口签名步骤：
1. sha256,16进制输出【注意这里需要转为16进制】
2. rsa sha256签名 base64输出
*/

// 签名
// @params value 待签名字符串
// @params priKey 格式为pem格式私钥，通过银联pfx证书提取私钥，  对应的提取cer格式公钥上传到银联商户平台
func RsaWithSha256Sign(value string, priKey string) (sign string, err error) {
	if priKey == "" {
		return "", errors.New("ERROR:RSA Sign 商户私钥不能为空")
	}

	p, _ := pem.Decode([]byte(priKey))
	privateKey, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	if err != nil {
		return
	}

	sha256hex := fmt.Sprintf("%x", sha256.Sum256([]byte(value)))
	hash := sha256.New()
	hash.Write([]byte(sha256hex))
	shaBytes := hash.Sum(nil)
	b, err := rsa.SignPKCS1v15(cryptorand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, shaBytes)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(b)
	return
}

// 验签
// @params value 待签名字符串
// @params 验签签名字符串，待签名字符串签名结果与此字符串不一致则返回异常
// @params cerKey 5.1.0 通过接口“银联加密公钥更新查询交易”获取，其它版本可在银联商户平台下载
func RsaWithSha256Verify(value string, sig string, cerKey string) (err error) {
	if cerKey == "" {
		return errors.New("ERROR:RSA Sign 银联签名公钥不能为空")
	}
	oriSign, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return
	}

	block, _ := pem.Decode([]byte(cerKey))
	if block == nil {
		err = errors.New("cer key error")
		return
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		err = errors.New("RsaWithSha256Verify ParseCertificate Error")
		return
	}
	sha256hex := fmt.Sprintf("%x", sha256.Sum256([]byte(value)))
	hashed := sha256.Sum256([]byte(sha256hex))
	err = rsa.VerifyPKCS1v15(cert.PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], oriSign)
	if err != nil {
		log.Error(err.Error())
		err = errors.New("验签失败")
		return
	}
	return
}
