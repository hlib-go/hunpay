package example

import (
	"github.com/hlib-go/hunpay/upapi"
	"testing"
)

func TestOauth2(t *testing.T) {
	url := upapi.Oauth2("b25226d0324d42499e66a1cd1d83b802", "https://ms.himkt.cn/mswork/result", "upapi_mobile", "")
	t.Log(url)
}
