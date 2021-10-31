package example

import (
	"github.com/BurntSushi/toml"
	"github.com/hlib-go/hunpay/upapi"
)

var cfgtoml *upapi.Config

func init() {
	_, err := toml.DecodeFile("D:\\Projects\\hlib-go\\hunpay\\upapi\\.secret\\宁波银联-闪券发券.toml", &cfgtoml)
	if err != nil {
		panic(err)
	}
}
