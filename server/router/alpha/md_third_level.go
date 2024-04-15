package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MdThirdLevelRouter struct {
}

// InitMdThirdLevelRouter 初始化 层级三 路由信息
func (s *MdThirdLevelRouter) InitMdThirdLevelRouter(Router *gin.RouterGroup) {
	mdThirdLevelRouter := Router.Group("mdThirdLevel").Use(middleware.OperationRecord())
	mdThirdLevelRouterWithoutRecord := Router.Group("mdThirdLevel")
	var mdThirdLevelApi = v1.ApiGroupApp.AlphaApiGroup.MdThirdLevelApi
	{
		mdThirdLevelRouter.POST("createMdThirdLevel", mdThirdLevelApi.CreateMdThirdLevel)             // 新建层级三
		mdThirdLevelRouter.DELETE("deleteMdThirdLevel", mdThirdLevelApi.DeleteMdThirdLevel)           // 删除层级三
		mdThirdLevelRouter.DELETE("deleteMdThirdLevelByIds", mdThirdLevelApi.DeleteMdThirdLevelByIds) // 批量删除层级三
		mdThirdLevelRouter.PUT("updateMdThirdLevel", mdThirdLevelApi.UpdateMdThirdLevel)              // 更新层级三
	}
	{
		mdThirdLevelRouterWithoutRecord.GET("findMdThirdLevel", mdThirdLevelApi.FindMdThirdLevel)       // 根据ID获取层级三
		mdThirdLevelRouterWithoutRecord.GET("getMdThirdLevelList", mdThirdLevelApi.GetMdThirdLevelList) // 获取层级三列表
	}
}
