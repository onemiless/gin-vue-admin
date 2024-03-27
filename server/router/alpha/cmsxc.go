package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CMSXCRouter struct {
}

// InitCMSXCRouter 初始化 车厂规范信息 路由信息
func (s *CMSXCRouter) InitCMSXCRouter(Router *gin.RouterGroup) {
	cmsxcRouter := Router.Group("cmsxc").Use(middleware.OperationRecord())
	cmsxcRouterWithoutRecord := Router.Group("cmsxc")
	var cmsxcApi = v1.ApiGroupApp.AlphaApiGroup.CMSXCApi
	{
		cmsxcRouter.POST("createCMSXC", cmsxcApi.CreateCMSXC)             // 新建车厂规范信息
		cmsxcRouter.DELETE("deleteCMSXC", cmsxcApi.DeleteCMSXC)           // 删除车厂规范信息
		cmsxcRouter.DELETE("deleteCMSXCByIds", cmsxcApi.DeleteCMSXCByIds) // 批量删除车厂规范信息
		cmsxcRouter.PUT("updateCMSXC", cmsxcApi.UpdateCMSXC)              // 更新车厂规范信息
	}
	{
		cmsxcRouterWithoutRecord.GET("findCMSXC", cmsxcApi.FindCMSXC)       // 根据ID获取车厂规范信息
		cmsxcRouterWithoutRecord.GET("getCMSXCList", cmsxcApi.GetCMSXCList) // 获取车厂规范信息列表
	}
}
