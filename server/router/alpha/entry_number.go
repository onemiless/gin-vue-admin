package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type EntryNumberRouter struct {
}

// InitMdClientRouter 初始化 客户信息 路由信息
func (s *EntryNumberRouter) InitEntryNumberRouter(Router *gin.RouterGroup) {
	//entryNumberRouter := Router.Group("entryNumber").Use(middleware.OperationRecord())
	entryNumberRouterWithoutRecord := Router.Group("entryNumber")
	var entryNumberApi = v1.ApiGroupApp.AlphaApiGroup.EntryNumberApi
	//{
	//	mdClientRouter.POST("createMdClient", mdClientApi.CreateMdClient)             // 新建客户信息
	//	mdClientRouter.DELETE("deleteMdClient", mdClientApi.DeleteMdClient)           // 删除客户信息
	//	mdClientRouter.DELETE("deleteMdClientByIds", mdClientApi.DeleteMdClientByIds) // 批量删除客户信息
	//	mdClientRouter.PUT("updateMdClient", mdClientApi.UpdateMdClient)              // 更新客户信息
	//}
	{
		entryNumberRouterWithoutRecord.POST("getEntryNumber", entryNumberApi.FindEntryNumber) // 根据ID获取客户信息
		//mdClientRouterWithoutRecord.GET("getMdClientList", mdClientApi.GetMdClientList) // 获取客户信息列表
	}
}
