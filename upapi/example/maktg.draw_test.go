package example

import (
	"github.com/hlib-go/hunpay/upapi"
	"testing"
	"time"
)

// 未测试成功
func TestMaktgDraw(t *testing.T) {
	err := upapi.MaktgDraw(cfgtoml, &upapi.MaktgDrawParams{
		TransSeqId: upapi.Rand32(),
		TransTs:    time.Now().Format("20060102"),
		Mobile:     "13611703040",
		ActivityNo: "1320210425336616",
	})
	if err != nil {
		t.Error(err)
		return
	}
}
