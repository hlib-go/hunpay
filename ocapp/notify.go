package ocapp

import (
	"errors"
	"net/url"
)

// 验证通知签名与返回码
func NotifyVerify(reqParams string) (bmap map[string]string, err error) {
	params, err := url.QueryUnescape(reqParams)
	if err != nil {
		return
	}

	// 参数转为Map
	bmap = ParamsConvertMap(params)

	// 验签
	signValue := RsaSignSortMap(bmap)
	err = RsaWithSha256Verify(signValue, bmap["signature"], bmap["signPubKeyCert"])
	if err != nil {
		return
	}

	// 验证返回码
	if bmap["respCode"] != RESP_OK {
		err = errors.New("UP" + bmap["respCode"] + ":" + bmap["respMsg"])
		return
	}
	return
}
