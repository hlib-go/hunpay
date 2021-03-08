package example

import (
	"github.com/hlib-go/hunpay/upapi"
)

// 宁波银联-宁波通联
var config = &upapi.Config{
	ServiceUrl:   "https://open.95516.com/open/access/1.0",
	AppId:        "9e211304be4a46fdb7dff03f7a01b2ef",
	Secret:       "",
	SymmetricKey: "bad5200bfe4a91e5cb02f1f2ef1aec08bad5200bfe4a91e5",
	UpPublicKey:  "",
	MchPrivateKey: `-----BEGIN PRIVATE KEY-----
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
-----END PRIVATE KEY-----`,
}

var config2 = &upapi.Config{
	ServiceUrl:   "https://open.95516.com/open/access/1.0",
	AppId:        "14525f83bb3c4dd68eca6511c93cb5ef",
	Secret:       "",
	SymmetricKey: "c761fef26ec7a86e64ba919b913e321cc761fef26ec7a86e",
	UpPublicKey:  "",
	MchPrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAO17/rEMAmFdl/pi
3q3PeT7DGeLOko7s3fupz2ine4oKp1yGI2nNNqSocctcERYQjD100DOVUaBCF6Ei
Fm1vicgqICAjr0bMczgvukkeuDUlLMAunJ7ZVNehpM6TrLRtF+eu0dzF8CVu1mVE
wrE+iYAkHKszueZ7yojayPoes37HAgMBAAECgYAhcV2rdwJ6zaBSTUwIc/giA32I
uAhhgi+8eexQU42NIfTxjZ49Dv4L8ACeX6e0UL4/BU0whm2JQOVs9ozr+lIPjZK3
l8Zp7zDQrcTZlS3dZzhR7HhM8oRsLnkMzVMk+3Pv0JePlPkQtll81rwz0DQewGWE
6SMt/kU6IzqLaCuecQJBAPKJN7aD8iro0lkDGDdcxRMbuhJBzWvO2FaykHmG9dtQ
uZPmQop28RG93L7JaJiwiKN2A9ZXhlXl8OSu/zEcbKkCQQD6qvxwqgVkhcTiBK7d
OFwm4M9jYyDLDEK6cv9uHSMMzRvWpSQ3oKha7Tpgm9EqjRM6E4wUWI1uyLR2yjbE
rMXvAkBcWxw4CC6jYF0ZQDBshIsXJ6vHX/9VWkLPYNfbLyVYCnlgdIJKL9jEpMP2
csO9wRuHA12atWGWPCVrL6hFj0lhAkEApk26naSvXznAnZMt0GcL/F86OF4T66J+
wuR4wr1h+6Q4y/dUR/O2vlyVVnMKGojuMKG3VehKLS3LTORr4aAe0QJAPViFv6ia
ARaw1XJEdrzX1saPFQjUPCRs8Tz+6VBBKn8GiWeatuMmp75DiYTLLPdEnZTYOs/8
XLKWa5Y2rdHbWQ==
-----END RSA PRIVATE KEY-----`,
}

/*
请求URL：https://open.95516.com/open/access/1.0/coupon.download
请求报文：{"acctEntityTp":"03","appId":"14525f83bb3c4dd68eca6511c93cb5ef","cardNo":"","couponId":"3102021022567015","couponNum":1,"mobile":"qPyfpSGrtO1QJEZjOssu1A==","nonceStr":"Umct9GqsJGyfbhs2","openId":"","signature":"LutFFJxOLGANrSIpmdGfWUprioEexZ2U2RvwIxjviCTdeO0elUrfdUBpHlFSRJV1pnYQlpS4+DMegipRg8cvSL76cwheoXMJOksP7fzY/7XYLhU2RIkyVGSJZhLTRmb0DN30v9ZWuBJdpoN3rEIkQxT1KAo6gh/QxXU/5GqYjCU=","timestamp":"1614583547","transSeqId":"9e57ca8531a3f4766fb71c6091dea6f4","transTs":"20210301"}
响应报文：{"resp":"N60003","msg":"加解密失败[N60003]","params":{}}
*/
