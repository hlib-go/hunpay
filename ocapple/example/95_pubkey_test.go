package main

import (
	"github.com/hlib-go/hunpay/ocapp"
	"github.com/hlib-go/hunpay/ocapple"
	"testing"
)

// 获取银联平台公钥
func TestPubkey(t *testing.T) {
	_, err := ocapple.Pubkey(cfg)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success................")
}

// 获取银联平台公钥2
func TestPubkey2(t *testing.T) {
	_, err := ocapp.Pubkey(cfg821330248164060)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("success................")
}
