package test

import (
	"fmt"
	"testing"
)

func TestFrontToken(t *testing.T) {
	r, err := c.FrontToken()
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log("FrontToken：", r.FrontToken)
}

// 5.1.1 获取backendToken<backendToken>
func TestBackendToken(t *testing.T) {
	r, err := c.BackendToken()
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log("BackendToken：", r.BackendToken)
}
