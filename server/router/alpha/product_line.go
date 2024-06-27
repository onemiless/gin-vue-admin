package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductLineRouter struct{}

// InitProductLineRouter 初始化 产线信息 路由信息
func (s *ProductLineRouter) InitProductLineRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	productLineRouter := Router.Group("productLine").Use(middleware.OperationRecord())
	productLineRouterWithoutRecord := Router.Group("productLine")
	productLineRouterWithoutAuth := PublicRouter.Group("productLine")

	var productLineApi = v1.ApiGroupApp.AlphaApiGroup.ProductLineApi
	{
		productLineRouter.POST("createProductLine", productLineApi.CreateProductLine)             // 新建产线信息
		productLineRouter.DELETE("deleteProductLine", productLineApi.DeleteProductLine)           // 删除产线信息
		productLineRouter.DELETE("deleteProductLineByIds", productLineApi.DeleteProductLineByIds) // 批量删除产线信息
		productLineRouter.PUT("updateProductLine", productLineApi.UpdateProductLine)              // 更新产线信息
	}
	{
		productLineRouterWithoutRecord.GET("findProductLine", productLineApi.FindProductLine)       // 根据ID获取产线信息
		productLineRouterWithoutRecord.GET("getProductLineList", productLineApi.GetProductLineList) // 获取产线信息列表
	}
	{
		productLineRouterWithoutAuth.GET("getProductLinePublic", productLineApi.GetProductLinePublic) // 获取产线信息列表
	}
}
