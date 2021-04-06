package main

import (
	"encoding/json"
	"github.com/hlib-go/hunpay/ocapple"
	"testing"
)

// 交易状态查询（）
func TestQuery(t *testing.T) {
	result, err := ocapple.Query(cfg, "1356108669200142336", "20210201131436")
	if err != nil {
		t.Error(err.Error())
		return
	}
	rbytes, _ := json.Marshal(result)
	t.Log(string(rbytes))
}
