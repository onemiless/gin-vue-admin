package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TecBaseProcessRouter struct {
}

// InitTecBaseProcessRouter 初始化 工艺、设备及治具基本信息 路由信息
func (s *TecBaseProcessRouter) InitTecBaseProcessRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	tecBaseProcessRouter := Router.Group("tecBaseProcess").Use(middleware.OperationRecord())
	tecBaseProcessRouterWithoutRecord := Router.Group("tecBaseProcess")
	tecBaseProcessRouterWithoutAuth := PublicRouter.Group("tecBaseProcess")

	var tecBaseProcessApi = v1.ApiGroupApp.AlphaApiGroup.TecBaseProcessApi
	{
		tecBaseProcessRouter.POST("createTecBaseProcess", tecBaseProcessApi.CreateTecBaseProcess)             // 新建工艺、设备及治具基本信息
		tecBaseProcessRouter.DELETE("deleteTecBaseProcess", tecBaseProcessApi.DeleteTecBaseProcess)           // 删除工艺、设备及治具基本信息
		tecBaseProcessRouter.DELETE("deleteTecBaseProcessByIds", tecBaseProcessApi.DeleteTecBaseProcessByIds) // 批量删除工艺、设备及治具基本信息
		tecBaseProcessRouter.PUT("updateTecBaseProcess", tecBaseProcessApi.UpdateTecBaseProcess)              // 更新工艺、设备及治具基本信息
	}
	{
		tecBaseProcessRouterWithoutRecord.GET("findTecBaseProcess", tecBaseProcessApi.FindTecBaseProcess)       // 根据ID获取工艺、设备及治具基本信息
		tecBaseProcessRouterWithoutRecord.GET("getTecBaseProcessList", tecBaseProcessApi.GetTecBaseProcessList) // 获取工艺、设备及治具基本信息列表
	}
	{
		tecBaseProcessRouterWithoutAuth.GET("getTecBaseProcessPublic", tecBaseProcessApi.GetTecBaseProcessPublic) // 获取工艺、设备及治具基本信息列表
	}
}
