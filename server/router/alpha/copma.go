package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type COPMARouter struct {
}

// InitCOPMARouter 初始化 ERP客户信息 路由信息
func (s *COPMARouter) InitCOPMARouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	copmaRouter := Router.Group("copma").Use(middleware.OperationRecord())
	copmaRouterWithoutRecord := Router.Group("copma")
	copmaRouterWithoutAuth := PublicRouter.Group("copma")

	var copmaApi = v1.ApiGroupApp.AlphaApiGroup.COPMAApi
	{
		copmaRouter.POST("createCOPMA", copmaApi.CreateCOPMA)             // 新建ERP客户信息
		copmaRouter.DELETE("deleteCOPMA", copmaApi.DeleteCOPMA)           // 删除ERP客户信息
		copmaRouter.DELETE("deleteCOPMAByIds", copmaApi.DeleteCOPMAByIds) // 批量删除ERP客户信息
		copmaRouter.PUT("updateCOPMA", copmaApi.UpdateCOPMA)              // 更新ERP客户信息
	}
	{
		copmaRouterWithoutRecord.GET("findCOPMA", copmaApi.FindCOPMA)       // 根据ID获取ERP客户信息
		copmaRouterWithoutRecord.GET("getCOPMAList", copmaApi.GetCOPMAList) // 获取ERP客户信息列表
	}
	{
		copmaRouterWithoutAuth.GET("getCOPMAPublic", copmaApi.GetCOPMAPublic) // 获取ERP客户信息列表
	}
}
