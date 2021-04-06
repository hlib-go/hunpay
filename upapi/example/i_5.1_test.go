package example

import (
	"fmt"
	"github.com/hlib-go/hunpay/upapi"
	"testing"
)

func TestFrontToken(t *testing.T) {
	r, err := upapi.FrontToken(appid_14525f83bb3c4dd68eca6511c93cb5ef)
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log("FrontToken：", r.FrontToken)
}

// 5.1.1 获取backendToken<backendToken>
func TestBackendToken(t *testing.T) {
	r, err := upapi.BackendToken(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log("BackendToken：", r.BackendToken)
}
