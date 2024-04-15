package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MdSecondLevelRouter struct {
}

// InitMdSecondLevelRouter 初始化 层级二 路由信息
func (s *MdSecondLevelRouter) InitMdSecondLevelRouter(Router *gin.RouterGroup) {
	mdSecondLevelRouter := Router.Group("mdSecondLevel").Use(middleware.OperationRecord())
	mdSecondLevelRouterWithoutRecord := Router.Group("mdSecondLevel")
	var mdSecondLevelApi = v1.ApiGroupApp.AlphaApiGroup.MdSecondLevelApi
	{
		mdSecondLevelRouter.POST("createMdSecondLevel", mdSecondLevelApi.CreateMdSecondLevel)             // 新建层级二
		mdSecondLevelRouter.DELETE("deleteMdSecondLevel", mdSecondLevelApi.DeleteMdSecondLevel)           // 删除层级二
		mdSecondLevelRouter.DELETE("deleteMdSecondLevelByIds", mdSecondLevelApi.DeleteMdSecondLevelByIds) // 批量删除层级二
		mdSecondLevelRouter.PUT("updateMdSecondLevel", mdSecondLevelApi.UpdateMdSecondLevel)              // 更新层级二
	}
	{
		mdSecondLevelRouterWithoutRecord.GET("findMdSecondLevel", mdSecondLevelApi.FindMdSecondLevel)       // 根据ID获取层级二
		mdSecondLevelRouterWithoutRecord.GET("getMdSecondLevelList", mdSecondLevelApi.GetMdSecondLevelList) // 获取层级二列表
	}
}
