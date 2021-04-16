package example

import (
	"encoding/json"
	"github.com/hlib-go/hunpay/upapi"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

var appid_14525f83bb3c4dd68eca6511c93cb5ef *upapi.Config

func init() {
	log.SetReportCaller(true)
	fbs, err := ioutil.ReadFile("/certs/cfg.appid.14525f83bb3c4dd68eca6511c93cb5ef.json")
	if err != nil {
		log.Error(err)
	}
	err = json.Unmarshal(fbs, &appid_14525f83bb3c4dd68eca6511c93cb5ef)
	if err != nil {
		log.Error(err)
	}
	appid_14525f83bb3c4dd68eca6511c93cb5ef.MchPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

}

/*
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDte/6xDAJhXZf6Yt6tz3k+wxnizpKO7N37qc9op3uKCqdchiNpzTakqHHLXBEWEIw9dNAzlVGgQhehIhZtb4nIKiAgI69GzHM4L7pJHrg1JSzALpye2VTXoaTOk6y0bRfnrtHcxfAlbtZlRMKxPomAJByrM7nme8qI2sj6HrN+xwIDAQAB

*/

/*
请求URL：https://open.95516.com/open/access/1.0/coupon.download
请求报文：{"acctEntityTp":"03","appId":"14525f83bb3c4dd68eca6511c93cb5ef","cardNo":"","couponId":"3102021022567015","couponNum":1,"mobile":"qPyfpSGrtO1QJEZjOssu1A==","nonceStr":"Umct9GqsJGyfbhs2","openId":"","signature":"LutFFJxOLGANrSIpmdGfWUprioEexZ2U2RvwIxjviCTdeO0elUrfdUBpHlFSRJV1pnYQlpS4+DMegipRg8cvSL76cwheoXMJOksP7fzY/7XYLhU2RIkyVGSJZhLTRmb0DN30v9ZWuBJdpoN3rEIkQxT1KAo6gh/QxXU/5GqYjCU=","timestamp":"1614583547","transSeqId":"9e57ca8531a3f4766fb71c6091dea6f4","transTs":"20210301"}
响应报文：{"resp":"N60003","msg":"加解密失败[N60003]","params":{}}
*/
