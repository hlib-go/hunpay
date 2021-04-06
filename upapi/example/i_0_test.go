package example

import (
	"github.com/hlib-go/hunpay/upapi"
	"io/ioutil"
)

var sbytes, _ = ioutil.ReadFile("../.secret/宁波银联闪券发券-secret")

// 宁波银联-宁波通联
var config = &upapi.Config{
	ServiceUrl:   "https://open.95516.com/open/access/1.0",
	AppId:        "9e211304be4a46fdb7dff03f7a01b2ef",
	Secret:       string(sbytes),
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
