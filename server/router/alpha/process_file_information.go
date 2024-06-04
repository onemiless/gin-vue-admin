package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProcessFileInformationRouter struct {
}

// InitProcessFileInformationRouter 初始化 工艺文件信息 路由信息
func (s *ProcessFileInformationRouter) InitProcessFileInformationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	processFileInformationRouter := Router.Group("processFileInformation").Use(middleware.OperationRecord())
	processFileInformationRouterWithoutRecord := Router.Group("processFileInformation")
	processFileInformationRouterWithoutAuth := PublicRouter.Group("processFileInformation")

	var processFileInformationApi = v1.ApiGroupApp.AlphaApiGroup.ProcessFileInformationApi
	{
		processFileInformationRouter.POST("createProcessFileInformation", processFileInformationApi.CreateProcessFileInformation)             // 新建工艺文件信息
		processFileInformationRouter.DELETE("deleteProcessFileInformation", processFileInformationApi.DeleteProcessFileInformation)           // 删除工艺文件信息
		processFileInformationRouter.DELETE("deleteProcessFileInformationByIds", processFileInformationApi.DeleteProcessFileInformationByIds) // 批量删除工艺文件信息
		processFileInformationRouter.PUT("updateProcessFileInformation", processFileInformationApi.UpdateProcessFileInformation)              // 更新工艺文件信息
	}
	{
		processFileInformationRouterWithoutRecord.GET("findProcessFileInformation", processFileInformationApi.FindProcessFileInformation)       // 根据ID获取工艺文件信息
		processFileInformationRouterWithoutRecord.GET("getProcessFileInformationList", processFileInformationApi.GetProcessFileInformationList) // 获取工艺文件信息列表
	}
	{
		processFileInformationRouterWithoutAuth.GET("getProcessFileInformationPublic", processFileInformationApi.GetProcessFileInformationPublic) // 获取工艺文件信息列表
	}
}
