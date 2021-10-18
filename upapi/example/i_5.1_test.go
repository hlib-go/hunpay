package example

import (
	"fmt"
	"github.com/hlib-go/hunpay/upapi"
	"testing"
)

func TestFrontToken(t *testing.T) {
	r, err := upapi.FrontToken(cfgtoml)
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log("FrontToken：", r.FrontToken)
}

// 5.1.1 获取backendToken<backendToken>
func TestBackendToken(t *testing.T) {
	r, err := upapi.BackendToken(cfgtoml)
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log("BackendToken：", r.BackendToken)
}
