package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/config"
	"sync"
)

var WxGlobal = new(config.WxGlobanConfig)

type CodeInfo struct {
	CanLogin   bool
	OpenId     string
	CreateTime int64
}

var LoginCodeMap = sync.Map{}

var SUBSCRIBE = "subscribe"     //(订阅 关注)
var UNSUBSCRIBE = "unsubscribe" //(退订 取消关注)
var SCAN = "SCAN"               //(扫码)
