package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TecBaseInfoRouter struct {
}

// InitTecBaseInfoRouter 初始化 技术部基础信息 路由信息
func (s *TecBaseInfoRouter) InitTecBaseInfoRouter(Router *gin.RouterGroup) {
	tecBaseInfoRouter := Router.Group("tecBaseInfo").Use(middleware.OperationRecord())
	tecBaseInfoRouterWithoutRecord := Router.Group("tecBaseInfo")
	var tecBaseInfoApi = v1.ApiGroupApp.AlphaApiGroup.TecBaseInfoApi
	{
		tecBaseInfoRouter.POST("createTecBaseInfo", tecBaseInfoApi.CreateTecBaseInfo)             // 新建技术部基础信息
		tecBaseInfoRouter.DELETE("deleteTecBaseInfo", tecBaseInfoApi.DeleteTecBaseInfo)           // 删除技术部基础信息
		tecBaseInfoRouter.DELETE("deleteTecBaseInfoByIds", tecBaseInfoApi.DeleteTecBaseInfoByIds) // 批量删除技术部基础信息
		tecBaseInfoRouter.PUT("updateTecBaseInfo", tecBaseInfoApi.UpdateTecBaseInfo)              // 更新技术部基础信息
	}
	{
		tecBaseInfoRouterWithoutRecord.GET("findTecBaseInfo", tecBaseInfoApi.FindTecBaseInfo)       // 根据ID获取技术部基础信息
		tecBaseInfoRouterWithoutRecord.GET("getTecBaseInfoList", tecBaseInfoApi.GetTecBaseInfoList) // 获取技术部基础信息列表
	}
}
