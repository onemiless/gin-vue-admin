package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MdFirstLevelRouter struct {
}

// InitMdFirstLevelRouter 初始化 层级一 路由信息
func (s *MdFirstLevelRouter) InitMdFirstLevelRouter(Router *gin.RouterGroup) {
	mdFirstLevelRouter := Router.Group("mdFirstLevel").Use(middleware.OperationRecord())
	mdFirstLevelRouterWithoutRecord := Router.Group("mdFirstLevel")
	var mdFirstLevelApi = v1.ApiGroupApp.AlphaApiGroup.MdFirstLevelApi
	{
		mdFirstLevelRouter.POST("createMdFirstLevel", mdFirstLevelApi.CreateMdFirstLevel)             // 新建层级一
		mdFirstLevelRouter.DELETE("deleteMdFirstLevel", mdFirstLevelApi.DeleteMdFirstLevel)           // 删除层级一
		mdFirstLevelRouter.DELETE("deleteMdFirstLevelByIds", mdFirstLevelApi.DeleteMdFirstLevelByIds) // 批量删除层级一
		mdFirstLevelRouter.PUT("updateMdFirstLevel", mdFirstLevelApi.UpdateMdFirstLevel)              // 更新层级一
	}
	{
		mdFirstLevelRouterWithoutRecord.GET("findMdFirstLevel", mdFirstLevelApi.FindMdFirstLevel)       // 根据ID获取层级一
		mdFirstLevelRouterWithoutRecord.GET("getMdFirstLevelList", mdFirstLevelApi.GetMdFirstLevelList) // 获取层级一列表
	}
}
