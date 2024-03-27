package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/api"
	"github.com/gin-gonic/gin"
)

type WxLoginRouter struct {
}

func (s *WxLoginRouter) InitWxLoginRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.WxLoginApi
	{
		plugRouter.GET("pingpang", plugApi.Signin)
		plugRouter.POST("pingpang", plugApi.PingPang)
		plugRouter.GET("getLoginPic", plugApi.GetLoginPic)
		plugRouter.GET("loginOrCreate", plugApi.LoginOrCreate)
		plugRouter.POST("clearWx", plugApi.ClearWx)
		plugRouter.POST("register", plugApi.Register)
	}
}
