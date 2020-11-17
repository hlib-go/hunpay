package main

import (
	"github.com/hlib-go/hunpay/ocwap"
	"testing"
)

// 获取银联平台公钥
func TestPubkey(t *testing.T) {
	_, err := ocwap.Pubkey(cfg)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success................")
}
