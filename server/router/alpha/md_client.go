package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MdClientRouter struct {
}

// InitMdClientRouter 初始化 客户信息 路由信息
func (s *MdClientRouter) InitMdClientRouter(Router *gin.RouterGroup) {
	mdClientRouter := Router.Group("mdClient").Use(middleware.OperationRecord())
	mdClientRouterWithoutRecord := Router.Group("mdClient")
	var mdClientApi = v1.ApiGroupApp.AlphaApiGroup.MdClientApi
	{
		mdClientRouter.POST("createMdClient", mdClientApi.CreateMdClient)             // 新建客户信息
		mdClientRouter.DELETE("deleteMdClient", mdClientApi.DeleteMdClient)           // 删除客户信息
		mdClientRouter.DELETE("deleteMdClientByIds", mdClientApi.DeleteMdClientByIds) // 批量删除客户信息
		mdClientRouter.PUT("updateMdClient", mdClientApi.UpdateMdClient)              // 更新客户信息
	}
	{
		mdClientRouterWithoutRecord.GET("findMdClient", mdClientApi.FindMdClient)       // 根据ID获取客户信息
		mdClientRouterWithoutRecord.GET("getMdClientList", mdClientApi.GetMdClientList) // 获取客户信息列表
	}
}
