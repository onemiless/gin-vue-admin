package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TecBaseInfoExtRouter struct {
}

// InitTecBaseInfoExtRouter 初始化 零件基础信息扩展 路由信息
func (s *TecBaseInfoExtRouter) InitTecBaseInfoExtRouter(Router *gin.RouterGroup) {
	tecBaseInfoExtRouter := Router.Group("tecBaseInfoExt").Use(middleware.OperationRecord())
	tecBaseInfoExtRouterWithoutRecord := Router.Group("tecBaseInfoExt")
	var tecBaseInfoExtApi = v1.ApiGroupApp.AlphaApiGroup.TecBaseInfoExtApi
	{
		tecBaseInfoExtRouter.POST("createTecBaseInfoExt", tecBaseInfoExtApi.CreateTecBaseInfoExt)             // 新建零件基础信息扩展
		tecBaseInfoExtRouter.DELETE("deleteTecBaseInfoExt", tecBaseInfoExtApi.DeleteTecBaseInfoExt)           // 删除零件基础信息扩展
		tecBaseInfoExtRouter.DELETE("deleteTecBaseInfoExtByIds", tecBaseInfoExtApi.DeleteTecBaseInfoExtByIds) // 批量删除零件基础信息扩展
		tecBaseInfoExtRouter.PUT("updateTecBaseInfoExt", tecBaseInfoExtApi.UpdateTecBaseInfoExt)              // 更新零件基础信息扩展
	}
	{
		tecBaseInfoExtRouterWithoutRecord.GET("findTecBaseInfoExt", tecBaseInfoExtApi.FindTecBaseInfoExt)       // 根据ID获取零件基础信息扩展
		tecBaseInfoExtRouterWithoutRecord.GET("getTecBaseInfoExtList", tecBaseInfoExtApi.GetTecBaseInfoExtList) // 获取零件基础信息扩展列表
		tecBaseInfoExtRouterWithoutRecord.GET("checkIsDuplicate", tecBaseInfoExtApi.CheckIsDuplicate)           // 检查是否重复
	}
}
