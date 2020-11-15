package upsdk

import (
	"crypto"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"sort"
	"strings"
)

// 云闪付 Rsa 验签
func UpRsaSign(params *BodyMap, priKey string, containNilVal bool) (sign string, err error) {
	if priKey == "" {
		return "", errors.New("ERROR:云闪付接口 RSA Sign privateKey 私钥配置不能为空")
	}
	value := rsaSignSortMap(params, containNilVal)
	return RsaSign(value, priKey)
}

func RsaSign(value, priKey string) (sign string, err error) {
	p, _ := pem.Decode([]byte(priKey))
	privateKey, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	if err != nil {
		return
	}

	hash := sha256.New()
	hash.Write([]byte(value))
	shaBytes := hash.Sum(nil)
	b, err := rsa.SignPKCS1v15(cryptorand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, shaBytes)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(b)
	return
}

// @params containNilVal true空字段参与签名 false空字段不参与签名
func rsaSignSortMap(params *BodyMap, containNilVal bool) string {
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range params.m {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if "signature" == k || "symmetricKey" == k {
			continue
		}
		// 不包含value为空的字段
		if !containNilVal && params.Get(k) == "" {
			continue
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(params.Get(k))
		buf.WriteByte('&')
	}
	s := buf.String()
	s = s[0 : len(s)-1]
	return s
}
