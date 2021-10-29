package example

import (
	"github.com/BurntSushi/toml"
	"github.com/hlib-go/hunpay/upapi"
)

var cfgtoml *upapi.Config

func init() {
	_, err := toml.DecodeFile("D:\\Projects\\hlib-go\\hunpay\\upapi\\.secret\\宁波银联-U惠天天转.toml", &cfgtoml)
	if err != nil {
		panic(err)
	}
}
