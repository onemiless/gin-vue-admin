package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PackageRequirementRouter struct {
}

// InitPackageRequirementRouter 初始化 包装要求信息 路由信息
func (s *PackageRequirementRouter) InitPackageRequirementRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	packageRequirementRouter := Router.Group("packageRequirement").Use(middleware.OperationRecord())
	packageRequirementRouterWithoutRecord := Router.Group("packageRequirement")
	packageRequirementRouterWithoutAuth := PublicRouter.Group("packageRequirement")

	var packageRequirementApi = v1.ApiGroupApp.AlphaApiGroup.PackageRequirementApi
	{
		packageRequirementRouter.POST("createPackageRequirement", packageRequirementApi.CreatePackageRequirement)             // 新建包装要求信息
		packageRequirementRouter.DELETE("deletePackageRequirement", packageRequirementApi.DeletePackageRequirement)           // 删除包装要求信息
		packageRequirementRouter.DELETE("deletePackageRequirementByIds", packageRequirementApi.DeletePackageRequirementByIds) // 批量删除包装要求信息
		packageRequirementRouter.PUT("updatePackageRequirement", packageRequirementApi.UpdatePackageRequirement)              // 更新包装要求信息
	}
	{
		packageRequirementRouterWithoutRecord.GET("findPackageRequirement", packageRequirementApi.FindPackageRequirement)       // 根据ID获取包装要求信息
		packageRequirementRouterWithoutRecord.GET("getPackageRequirementList", packageRequirementApi.GetPackageRequirementList) // 获取包装要求信息列表
	}
	{
		packageRequirementRouterWithoutAuth.GET("getPackageRequirementPublic", packageRequirementApi.GetPackageRequirementPublic) // 获取包装要求信息列表
	}
}
