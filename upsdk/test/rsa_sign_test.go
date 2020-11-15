package test

import (
	"encoding/base64"
	"encoding/pem"
	"fmt"
	log "github.com/sirupsen/logrus"
	"hunpay/upsdk"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var channelpriKey = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7kdKuAnMgu7AV
D4hfaT9i4TpDBVxN0xA2A6vGppaHB5F8N9vHRCBJLdAm04+HkdDmxG41Hbq7lABn
rM92vrCbSA1Lo6asOawaif8dnu/i92SOnoK+r7v0vNR+hfIAXqO2xCJB9a1IvprN
Uw9V8m8ALr8eLHBJ+0sFsnrgvJG4gYl4q0/+pzERJK0SKqFceUKpcqunfsZPv4ko
XK9Q7kSZxi+i3iwwmQ/7IBnsEVB78v4KGrkXDI/nrydauyGGXYXkTCbFiOw8CuxA
Skx4k0kwpA8nuvFXzBG0V7EMox9oQtiOkfGJgDUJvmYYqg8rpgDb54iAvKZMmz+4
x/4UupJjAgMBAAECggEAFVVqnvwMWCbAykRwAFoaKYbwd3r+mqNs7pfQS9HawRTt
STGZP7rR6UDase/SHVtKZVTmLAhrmrYkraYMGrdpot+5E2dTp7cPih0z9QyEwE3f
FBGXUVTvjdCEYredZMle2YTJWLM2uFVligDud5oRYfXvKuFnDCMWz1kTfMg10sRH
Aw5lLHJk1cJSsB+s8swHzch+IsFg5oyA6VcpFKPiKvMwy8m1A923nH+mVVVcj4wM
B0/qCtxlmIAUUg5MKp/RGgKPPxReTg6bqF2t6wNrHZevVOsFhDStizwf6dwLYCLV
wK3I6Szwp+7uR284hVKZLu2uwSuREdxi1Xc6cGMuAQKBgQDrWpVo0LKaLNjnbcVL
UuFCyuN0M8aWMfMrG1NtoYYPC5hFK5ZTAV3euf7yTGHazfswJ9A6U+h9+0XsIhfY
8lN9HvO2ah05Uo2EPSGhaLa0ziFw4Nbi5usd3X6vnQq270Q/BoHi8fSoDou9oelQ
neG3zIZab3cHYzEzKnq1rbUEgQKBgQDMBiJCTKx94lU8UdIIRDg6VEfIrya3+9F4
BTQI4xiSdikdz5iZC2tt5gGHnaKeYVwsAioHJmFBWyWu0YQXb0qt1X+vbM/d9v1z
T/mJ3KSpx928RdkwVKQsGQFsYjPgDVpweQkzybFEJFstOaNHXIw3+RhQG37UXuOr
6sD05B2U4wKBgQDongyEn5mXpvHvs+BH9a/te2japn4GX3JPzd9kwTwmTLiAzXbz
rashA8cHpxUk1WgLDZ7St7JYKm3O2VemtsRsK5aIWlNuH7j91goSZdQH2qDU13Ws
qL4EM7MOUfKQIubaQE1KiQjevhnCIXDgnFvHdV/prLgB1jl/r+G/BeSfgQKBgFqH
fjwc+Y0CGQAi7ids3eZD73ZFAdExk8jFxkkLO6QBek0YCIYgYxLotFUQxU+xs8xz
SWLSzOTLJPVlUk9zupdX3MhiZ/n91oiMPBXIKeiMHv+jnrOrWw2WKuOEz6/jPPYb
PtIT9OxflXWD1ceccTuE9BzXlnd1g2CNUgFYFygxAoGAbM2zG5dJSR4icH6LqBoH
EzgxZuLlRIyWS2wnyxoIneRfKEoXAnaeZTho0jYBqFsKsButRIVP1DLlzS99NCi0
N2MTKkYWZZbs2kFin7sB92Xy3QYoHeru4fZK3MdBRj85e17n9MBfVCnGTq6cbxbm
MG7BFjV7aoR/h2bkkVv6mxw=
-----END PRIVATE KEY-----
`

var channelpubkey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu5HSrgJzILuwFQ+IX2k/
YuE6QwVcTdMQNgOrxqaWhweRfDfbx0QgSS3QJtOPh5HQ5sRuNR26u5QAZ6zPdr6w
m0gNS6OmrDmsGon/HZ7v4vdkjp6Cvq+79LzUfoXyAF6jtsQiQfWtSL6azVMPVfJv
AC6/HixwSftLBbJ64LyRuIGJeKtP/qcxESStEiqhXHlCqXKrp37GT7+JKFyvUO5E
mcYvot4sMJkP+yAZ7BFQe/L+Chq5FwyP568nWrshhl2F5EwmxYjsPArsQEpMeJNJ
MKQPJ7rxV8wRtFexDKMfaELYjpHxiYA1Cb5mGKoPK6YA2+eIgLymTJs/uMf+FLqS
YwIDAQAB
-----END PUBLIC KEY-----
`

func TestRsaSign(t *testing.T) {
	value := "activityNumber=1320200413190749&appId=9e211304be4a46fdb7dff03f7a01b2ef&certId=430211198412100435&icTerminal=430211198412100435&nonceStr=CrjzALhpVkqBZ4ca&orderAmount=&qrCode=&qualNum=3a5d4792-48c3-4416-809d-ada5da535f84&qualType=mobile&qualValue=13611703040&timestamp=1594198929&transNumber=1594198929"
	sign, err := upsdk.RsaSign(value, channelpriKey)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(sign)
	// Pkb3CaLsfex+bthOVSGG+3SVon8avYknBaTeSh812HSuySOWd1RMOz+K5lRIMNAR6wzMWm15qx/U2IAN3PFLGxcH66aCgsvGBnJ9vZBUtZz+rlAmkRx01gXso8v2u9Vq9I9sSVbTGn+p1QbcfTv91avSOZghsRubHGKt7TCHNfHahoDByfFfA9vYOBlAJgJh8pH9meI61MqRhqjL+SmQZfGWE9zogrinYpl+spAJWgjKjq1LIdVutZFSOoMgrAOYZUHTCi4DR5msAGTYStNdC49f56WGUJE2zKnl7wZ0HZh/AfoLz+/Yame9wukYERccjhLWjb4c5UWvYEqqK8zsUw==
	// Pkb3CaLsfex+bthOVSGG+3SVon8avYknBaTeSh812HSuySOWd1RMOz+K5lRIMNAR6wzMWm15qx/U2IAN3PFLGxcH66aCgsvGBnJ9vZBUtZz+rlAmkRx01gXso8v2u9Vq9I9sSVbTGn+p1QbcfTv91avSOZghsRubHGKt7TCHNfHahoDByfFfA9vYOBlAJgJh8pH9meI61MqRhqjL+SmQZfGWE9zogrinYpl+spAJWgjKjq1LIdVutZFSOoMgrAOYZUHTCi4DR5msAGTYStNdC49f56WGUJE2zKnl7wZ0HZh/AfoLz+/Yame9wukYERccjhLWjb4c5UWvYEqqK8zsUw==
}

func TestUnionpay_Post(t *testing.T) {
	body := "{\"signature\":\"Z1f2gMj2LPrHPQAErhAEK015AEgvuaDWugRXlJFD7Nis14OoChh2NIvp2qcaKvtjIKqPigRyqEF09hT+Cx4YlyPtX8YqzfwV9ebrttHfWOsnU05Wq2CLbvbJ0D0EqGqayiOMyP9Le6o8zGAw+/ksQdf7hWEFuzj2ofdNXbLGxbJAf/X83vaCXCWUNCzwdf9anCjlUMpLilZvjKB36BcL8/nsqPx3Lb0muO54+haFvlRmGFd+eOTGyhAfs3bznPf48jW1uCVTxV/NWRcKFJpie5smDxRSyDTScL6QTWLoLcDPKe3B0GVNfNKuET6eZGRuRnpaNnDrHzyNqkoYdY1Vqw==\",\"qualNum\":\"3a5d4792-48c3-4416-809d-ada5da535f84\",\"certId\":\"G7Q5s8VhE4BJtvGeFpwogA+jKTOTtOZZ\",\"qualValue\":\"ORvPvDbB3XumZwSosRb/DA==\",\"nonceStr\":\"CrjzALhpVkqBZ4ca\",\"orderAmount\":\"\",\"qualType\":\"mobile\",\"qrCode\":\"\",\"appId\":\"9e211304be4a46fdb7dff03f7a01b2ef\",\"activityNumber\":\"1320200413190749\",\"icTerminal\":\"430211198412100435\",\"transNumber\":\"1594198928\",\"timestamp\":\"1594198929\"}"
	resp, err := http.Post("https://open.95516.com/open/access/1.0/qual.reduce", "application/json", strings.NewReader(body))
	if err != nil {
		t.Error(err)
		return
	}
	bytesBody, _ := ioutil.ReadAll(resp.Body)
	t.Log("resp.Body = " + string(bytesBody))
}

func TestRsaPub(t *testing.T) {
	p, _ := pem.Decode([]byte(channelpubkey))

	fmt.Println("base64:" + base64.StdEncoding.EncodeToString(p.Bytes))

	s := base64.StdEncoding.EncodeToString(p.Bytes)
	t.Log(s)
}
