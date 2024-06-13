package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type TecSNRuleRouter struct {
}

// InitTecSNRuleRouter 初始化 snrul路由
func (s *TecSNRuleRouter) InitTecSNRuleRouter(Router *gin.RouterGroup) {
	//mdThirdLevelRouter := Router.Group("tecSNRule").Use(middleware.OperationRecord())
	TecSNRuleRouterWithoutRecord := Router.Group("tecSNRule")
	var tecSNRuleApi = v1.ApiGroupApp.AlphaApiGroup.TecSNRuleApi
	//{
	//	mdThirdLevelRouter.POST("createMdThirdLevel", mdThirdLevelApi.CreateMdThirdLevel)             // 新建层级三
	//	mdThirdLevelRouter.DELETE("deleteMdThirdLevel", mdThirdLevelApi.DeleteMdThirdLevel)           // 删除层级三
	//	mdThirdLevelRouter.DELETE("deleteMdThirdLevelByIds", mdThirdLevelApi.DeleteMdThirdLevelByIds) // 批量删除层级三
	//	mdThirdLevelRouter.PUT("updateMdThirdLevel", mdThirdLevelApi.UpdateMdThirdLevel)              // 更新层级三
	//}
	{
		TecSNRuleRouterWithoutRecord.POST("findSNRuler", tecSNRuleApi.FindSNRuler) // 根据ID获取层级三
		//mdThirdLevelRouterWithoutRecord.GET("getMdThirdLevelList", mdThirdLevelApi.GetMdThirdLevelList) // 获取层级三列表
	}
}
