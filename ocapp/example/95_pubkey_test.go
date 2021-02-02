package main

import (
	"github.com/hlib-go/hunpay/ocapp"
	"testing"
)

// 获取银联平台公钥
func TestPubkey(t *testing.T) {
	_, err := ocapp.Pubkey(cfg)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success................")
}
