package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EngineeringChangeRouter struct {
}

// InitEngineeringChangeRouter 初始化 工程变更信息 路由信息
func (s *EngineeringChangeRouter) InitEngineeringChangeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	engineeringChangeRouter := Router.Group("engineeringChange").Use(middleware.OperationRecord())
	engineeringChangeRouterWithoutRecord := Router.Group("engineeringChange")
	engineeringChangeRouterWithoutAuth := PublicRouter.Group("engineeringChange")

	var engineeringChangeApi = v1.ApiGroupApp.AlphaApiGroup.EngineeringChangeApi
	{
		engineeringChangeRouter.POST("createEngineeringChange", engineeringChangeApi.CreateEngineeringChange)             // 新建工程变更信息
		engineeringChangeRouter.DELETE("deleteEngineeringChange", engineeringChangeApi.DeleteEngineeringChange)           // 删除工程变更信息
		engineeringChangeRouter.DELETE("deleteEngineeringChangeByIds", engineeringChangeApi.DeleteEngineeringChangeByIds) // 批量删除工程变更信息
		engineeringChangeRouter.PUT("updateEngineeringChange", engineeringChangeApi.UpdateEngineeringChange)              // 更新工程变更信息
	}
	{
		engineeringChangeRouterWithoutRecord.GET("findEngineeringChange", engineeringChangeApi.FindEngineeringChange)       // 根据ID获取工程变更信息
		engineeringChangeRouterWithoutRecord.GET("getEngineeringChangeList", engineeringChangeApi.GetEngineeringChangeList) // 获取工程变更信息列表
	}
	{
		engineeringChangeRouterWithoutAuth.GET("getEngineeringChangePublic", engineeringChangeApi.GetEngineeringChangePublic) // 获取工程变更信息列表
	}
}
