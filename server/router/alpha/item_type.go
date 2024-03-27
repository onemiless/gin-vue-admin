package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ItemTypeRouter struct {
}

// InitItemTypeRouter 初始化 零件类型 路由信息
func (s *ItemTypeRouter) InitItemTypeRouter(Router *gin.RouterGroup) {
	itemtypeRouter := Router.Group("itemtype").Use(middleware.OperationRecord())
	itemtypeRouterWithoutRecord := Router.Group("itemtype")
	var itemtypeApi = v1.ApiGroupApp.AlphaApiGroup.ItemTypeApi
	{
		itemtypeRouter.POST("createItemType", itemtypeApi.CreateItemType)             // 新建零件类型
		itemtypeRouter.DELETE("deleteItemType", itemtypeApi.DeleteItemType)           // 删除零件类型
		itemtypeRouter.DELETE("deleteItemTypeByIds", itemtypeApi.DeleteItemTypeByIds) // 批量删除零件类型
		itemtypeRouter.PUT("updateItemType", itemtypeApi.UpdateItemType)              // 更新零件类型
	}
	{
		itemtypeRouterWithoutRecord.GET("findItemType", itemtypeApi.FindItemType)                   // 根据ID获取零件类型
		itemtypeRouterWithoutRecord.GET("getItemTypeList", itemtypeApi.GetItemTypeList)             // 获取零件类型列表
		itemtypeRouterWithoutRecord.GET("getItemTypeListParent", itemtypeApi.GetItemTypeListParent) // 获取零件类型列表

	}
}
