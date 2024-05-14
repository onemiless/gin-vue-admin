package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CostCollectionRouter struct {
}

// InitCostCollectionRouter 初始化 成本信息收集信息 路由信息
func (s *CostCollectionRouter) InitCostCollectionRouter(Router *gin.RouterGroup) {
	costCollectionRouter := Router.Group("costCollection").Use(middleware.OperationRecord())
	costCollectionRouterWithoutRecord := Router.Group("costCollection")
	var costCollectionApi = v1.ApiGroupApp.AlphaApiGroup.CostCollectionApi
	{
		costCollectionRouter.POST("createCostCollection", costCollectionApi.CreateCostCollection)             // 新建成本信息收集信息
		costCollectionRouter.DELETE("deleteCostCollection", costCollectionApi.DeleteCostCollection)           // 删除成本信息收集信息
		costCollectionRouter.DELETE("deleteCostCollectionByIds", costCollectionApi.DeleteCostCollectionByIds) // 批量删除成本信息收集信息
		costCollectionRouter.PUT("updateCostCollection", costCollectionApi.UpdateCostCollection)              // 更新成本信息收集信息
	}
	{
		costCollectionRouterWithoutRecord.GET("findCostCollection", costCollectionApi.FindCostCollection)         // 根据ID获取成本信息收集信息
		costCollectionRouterWithoutRecord.GET("getCostCollectionList", costCollectionApi.GetCostCollectionList)   // 获取成本信息收集信息列表
		costCollectionRouterWithoutRecord.GET("getCostCollectionNumer", costCollectionApi.GetCostCollectionNumer) // 获取成本信息收集信息流水号
	}
}
