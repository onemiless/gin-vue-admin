package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/passport"
)

type ApiGroup struct {
	WxLoginApi
}

var ApiGroupApp = new(ApiGroup)

var wxLoginPassport = new(passport.WxLoginPassport)
