package main

import (
	"github.com/hlib-go/hunpay/ocweb"
	"testing"
)

// 获取银联平台公钥
func TestPubkey(t *testing.T) {
	_, err := ocweb.Pubkey(cfg)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success................")
}
