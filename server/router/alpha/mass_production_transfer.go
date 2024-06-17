package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MassProductionTransferRouter struct {
}

// InitMassProductionTransferRouter 初始化 量产转移 路由信息
func (s *MassProductionTransferRouter) InitMassProductionTransferRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	massProductionTransferRouter := Router.Group("massProductionTransfer").Use(middleware.OperationRecord())
	massProductionTransferRouterWithoutRecord := Router.Group("massProductionTransfer")
	massProductionTransferRouterWithoutAuth := PublicRouter.Group("massProductionTransfer")

	var massProductionTransferApi = v1.ApiGroupApp.AlphaApiGroup.MassProductionTransferApi
	{
		massProductionTransferRouter.POST("createMassProductionTransfer", massProductionTransferApi.CreateMassProductionTransfer)             // 新建量产转移
		massProductionTransferRouter.DELETE("deleteMassProductionTransfer", massProductionTransferApi.DeleteMassProductionTransfer)           // 删除量产转移
		massProductionTransferRouter.DELETE("deleteMassProductionTransferByIds", massProductionTransferApi.DeleteMassProductionTransferByIds) // 批量删除量产转移
		massProductionTransferRouter.PUT("updateMassProductionTransfer", massProductionTransferApi.UpdateMassProductionTransfer)              // 更新量产转移
	}
	{
		massProductionTransferRouterWithoutRecord.GET("findMassProductionTransfer", massProductionTransferApi.FindMassProductionTransfer)       // 根据ID获取量产转移
		massProductionTransferRouterWithoutRecord.GET("getMassProductionTransferList", massProductionTransferApi.GetMassProductionTransferList) // 获取量产转移列表
	}
	{
		massProductionTransferRouterWithoutAuth.GET("getMassProductionTransferPublic", massProductionTransferApi.GetMassProductionTransferPublic) // 获取量产转移列表
	}
}
