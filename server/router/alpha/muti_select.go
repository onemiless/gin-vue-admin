package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MutiSelectRouter struct {
}

// InitMutiSelectRouter 初始化 多级选择信息 路由信息
func (s *MutiSelectRouter) InitMutiSelectRouter(Router *gin.RouterGroup) {
	mutiSelectRouter := Router.Group("mutiSelect").Use(middleware.OperationRecord())
	mutiSelectRouterWithoutRecord := Router.Group("mutiSelect")
	var mutiSelectApi = v1.ApiGroupApp.AlphaApiGroup.MutiSelectApi
	{
		mutiSelectRouter.POST("createMutiSelect", mutiSelectApi.CreateMutiSelect)             // 新建多级选择信息
		mutiSelectRouter.DELETE("deleteMutiSelect", mutiSelectApi.DeleteMutiSelect)           // 删除多级选择信息
		mutiSelectRouter.DELETE("deleteMutiSelectByIds", mutiSelectApi.DeleteMutiSelectByIds) // 批量删除多级选择信息
		mutiSelectRouter.PUT("updateMutiSelect", mutiSelectApi.UpdateMutiSelect)              // 更新多级选择信息
	}
	{
		mutiSelectRouterWithoutRecord.GET("findMutiSelect", mutiSelectApi.FindMutiSelect)       // 根据ID获取多级选择信息
		mutiSelectRouterWithoutRecord.GET("getMutiSelectList", mutiSelectApi.GetMutiSelectList) // 获取多级选择信息列表
	}
}
