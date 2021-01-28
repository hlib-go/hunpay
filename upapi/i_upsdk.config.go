package upapi

import (
	"crypto/sha256"
	"fmt"
	"time"
)

/*
6.2 UPSDK使用步骤
步骤一：引入JS文件
在需要调用JS接口的页面引入如下JS文件（仅支持https）：
https://open.95516.com/s/open/js/upsdk.js

注意：
 upsdk.js文件依赖Zepto或Jquery。
 Zepto的版本要求1.0及以上，Jquery的版本要求1.4及以上（建议用最新的Zepto及Jquery）。

步骤二：通过config接口注入权限验证配置
所有需要使用UPSDK的页面必须先注入配置信息，否则将无法调用（同一个url仅需调用一次，如果跳转的页面无需使用插件，则无需config，否则需要重新执行config）。
//初始化示例代码
upsdk.config({
	appId:' ',   //必填，接入方的唯一标识，由云闪付分配
	timestamp:' ' ,   //必填，生成签名的时间戳，从1970年1月1日00:00:00至今的秒数，取东八区的北京时间，例1556096162
	nonceStr: ' ',    //必填，生成签名的随机串，参见附录三生成签名随机字符串
	signature: ' ',   //必填，生成签名的摘要，参见附录一、附录二
	debug: true   //开发阶段可打开此标记，云闪付APP会将调试信息toast出来
});
建议接入方在开发联调时，打开debug: true 开关，UPSDK只有在开关打开时才会输出状态信息，帮助开发者定位错误。 请务必在最终生产版本关闭此开关。

步骤三：配置信息验证
通过ready接口处理成功验证
upsdk.ready(function(){
//config信息验证后会执行ready方法
});
通过error接口处理失败验证
upsdk.error(function(err){
//config信息验证失败会执行error方法
});

*/

func UpsdkConfig(conf *Config, url, frontToken string, debug bool) (bm *BodyMap) {
	bm = NewBodyMap()
	bm.Set("appId", conf.AppId)
	bm.Set("nonceStr", GetRandomString(16))
	bm.Set("timestamp", time.Now().Unix())
	bm.Set("url", url)
	bm.Set("frontToken", frontToken)

	//计算签名,注意，upsdk的签名不用拼接secret，服务端的接口签名需要拼接secret参数
	//signature := bm.SignValue(c.Config.Secret)
	string1 := bm.FrontSignParams()
	signature := fmt.Sprintf("%x", sha256.Sum256([]byte(string1)))

	bm.Set("signature", signature)
	bm.Set("debug", debug)
	bm.Remove("url")
	bm.Remove("frontToken")
	bm.Remove("secret")
	return bm
}
