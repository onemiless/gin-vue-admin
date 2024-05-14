package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProofingInformationRouter struct {
}

// InitProofingInformationRouter 初始化 打样信息 路由信息
func (s *ProofingInformationRouter) InitProofingInformationRouter(Router *gin.RouterGroup) {
	proofingInformationRouter := Router.Group("proofingInformation").Use(middleware.OperationRecord())
	proofingInformationRouterWithoutRecord := Router.Group("proofingInformation")
	var proofingInformationApi = v1.ApiGroupApp.AlphaApiGroup.ProofingInformationApi
	{
		proofingInformationRouter.POST("createProofingInformation", proofingInformationApi.CreateProofingInformation)             // 新建打样信息
		proofingInformationRouter.DELETE("deleteProofingInformation", proofingInformationApi.DeleteProofingInformation)           // 删除打样信息
		proofingInformationRouter.DELETE("deleteProofingInformationByIds", proofingInformationApi.DeleteProofingInformationByIds) // 批量删除打样信息
		proofingInformationRouter.PUT("updateProofingInformation", proofingInformationApi.UpdateProofingInformation)              // 更新打样信息
	}
	{
		proofingInformationRouterWithoutRecord.GET("findProofingInformation", proofingInformationApi.FindProofingInformation)       // 根据ID获取打样信息
		proofingInformationRouterWithoutRecord.GET("getProofingInformationList", proofingInformationApi.GetProofingInformationList) // 获取打样信息列表
	}
}
