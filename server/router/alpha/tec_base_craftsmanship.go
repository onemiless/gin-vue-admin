package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TecBaseCraftsmanshipRouter struct{}

// InitTecBaseCraftsmanshipRouter 初始化 入篮量和程序号 路由信息
func (s *TecBaseCraftsmanshipRouter) InitTecBaseCraftsmanshipRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	tecBaseCraftsmanshipRouter := Router.Group("tecBaseCraftsmanship").Use(middleware.OperationRecord())
	tecBaseCraftsmanshipRouterWithoutRecord := Router.Group("tecBaseCraftsmanship")
	tecBaseCraftsmanshipRouterWithoutAuth := PublicRouter.Group("tecBaseCraftsmanship")

	var tecBaseCraftsmanshipApi = v1.ApiGroupApp.AlphaApiGroup.TecBaseCraftsmanshipApi
	{
		tecBaseCraftsmanshipRouter.POST("createTecBaseCraftsmanship", tecBaseCraftsmanshipApi.CreateTecBaseCraftsmanship)             // 新建入篮量和程序号
		tecBaseCraftsmanshipRouter.DELETE("deleteTecBaseCraftsmanship", tecBaseCraftsmanshipApi.DeleteTecBaseCraftsmanship)           // 删除入篮量和程序号
		tecBaseCraftsmanshipRouter.DELETE("deleteTecBaseCraftsmanshipByIds", tecBaseCraftsmanshipApi.DeleteTecBaseCraftsmanshipByIds) // 批量删除入篮量和程序号
		tecBaseCraftsmanshipRouter.PUT("updateTecBaseCraftsmanship", tecBaseCraftsmanshipApi.UpdateTecBaseCraftsmanship)              // 更新入篮量和程序号
	}
	{
		tecBaseCraftsmanshipRouterWithoutRecord.GET("findTecBaseCraftsmanship", tecBaseCraftsmanshipApi.FindTecBaseCraftsmanship)       // 根据ID获取入篮量和程序号
		tecBaseCraftsmanshipRouterWithoutRecord.GET("getTecBaseCraftsmanshipList", tecBaseCraftsmanshipApi.GetTecBaseCraftsmanshipList) // 获取入篮量和程序号列表
	}
	{
		tecBaseCraftsmanshipRouterWithoutAuth.GET("getTecBaseCraftsmanshipDataSource", tecBaseCraftsmanshipApi.GetTecBaseCraftsmanshipDataSource) // 获取入篮量和程序号数据源
		tecBaseCraftsmanshipRouterWithoutAuth.GET("getTecBaseCraftsmanshipPublic", tecBaseCraftsmanshipApi.GetTecBaseCraftsmanshipPublic)         // 获取入篮量和程序号列表
	}
}
