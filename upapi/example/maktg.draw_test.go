package example

import (
	"github.com/hlib-go/hunpay/upapi"
	"testing"
	"time"
)

// wei ce shi cheng gong
func TestMaktgDraw(t *testing.T) {
	err := upapi.MaktgDraw(config, &upapi.MaktgDrawParams{
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
