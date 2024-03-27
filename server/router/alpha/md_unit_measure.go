package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MdUnitMeasureRouter struct {
}

// InitMdUnitMeasureRouter 初始化 品号单位设置 路由信息
func (s *MdUnitMeasureRouter) InitMdUnitMeasureRouter(Router *gin.RouterGroup) {
	mdUnitMeasureRouter := Router.Group("mdUnitMeasure").Use(middleware.OperationRecord())
	mdUnitMeasureRouterWithoutRecord := Router.Group("mdUnitMeasure")
	var mdUnitMeasureApi = v1.ApiGroupApp.AlphaApiGroup.MdUnitMeasureApi
	{
		mdUnitMeasureRouter.POST("createMdUnitMeasure", mdUnitMeasureApi.CreateMdUnitMeasure)             // 新建品号单位设置
		mdUnitMeasureRouter.DELETE("deleteMdUnitMeasure", mdUnitMeasureApi.DeleteMdUnitMeasure)           // 删除品号单位设置
		mdUnitMeasureRouter.DELETE("deleteMdUnitMeasureByIds", mdUnitMeasureApi.DeleteMdUnitMeasureByIds) // 批量删除品号单位设置
		mdUnitMeasureRouter.PUT("updateMdUnitMeasure", mdUnitMeasureApi.UpdateMdUnitMeasure)              // 更新品号单位设置
	}
	{
		mdUnitMeasureRouterWithoutRecord.GET("findMdUnitMeasure", mdUnitMeasureApi.FindMdUnitMeasure)       // 根据ID获取品号单位设置
		mdUnitMeasureRouterWithoutRecord.GET("getMdUnitMeasureList", mdUnitMeasureApi.GetMdUnitMeasureList) // 获取品号单位设置列表
	}
}
