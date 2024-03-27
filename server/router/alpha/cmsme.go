package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CMSMERouter struct {
}

// InitCMSMERouter 初始化 ERP部门 路由信息
func (s *CMSMERouter) InitCMSMERouter(Router *gin.RouterGroup) {
	cmsmeRouter := Router.Group("cmsme").Use(middleware.OperationRecord())
	cmsmeRouterWithoutRecord := Router.Group("cmsme")
	var cmsmeApi = v1.ApiGroupApp.AlphaApiGroup.CMSMEApi
	{
		cmsmeRouter.POST("createCMSME", cmsmeApi.CreateCMSME)             // 新建ERP部门
		cmsmeRouter.DELETE("deleteCMSME", cmsmeApi.DeleteCMSME)           // 删除ERP部门
		cmsmeRouter.DELETE("deleteCMSMEByIds", cmsmeApi.DeleteCMSMEByIds) // 批量删除ERP部门
		cmsmeRouter.PUT("updateCMSME", cmsmeApi.UpdateCMSME)              // 更新ERP部门
	}
	{
		cmsmeRouterWithoutRecord.GET("findCMSME", cmsmeApi.FindCMSME)       // 根据ID获取ERP部门
		cmsmeRouterWithoutRecord.GET("getCMSMEList", cmsmeApi.GetCMSMEList) // 获取ERP部门列表
	}
}
