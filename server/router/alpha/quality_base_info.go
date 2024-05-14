package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QualityBaseInfoRouter struct {
}

// InitQualityBaseInfoRouter 初始化 质量要求信息 路由信息
func (s *QualityBaseInfoRouter) InitQualityBaseInfoRouter(Router *gin.RouterGroup) {
	qualityBaseInfoRouter := Router.Group("qualityBaseInfo").Use(middleware.OperationRecord())
	qualityBaseInfoRouterWithoutRecord := Router.Group("qualityBaseInfo")
	var qualityBaseInfoApi = v1.ApiGroupApp.AlphaApiGroup.QualityBaseInfoApi
	{
		qualityBaseInfoRouter.POST("createQualityBaseInfo", qualityBaseInfoApi.CreateQualityBaseInfo)             // 新建质量要求信息
		qualityBaseInfoRouter.DELETE("deleteQualityBaseInfo", qualityBaseInfoApi.DeleteQualityBaseInfo)           // 删除质量要求信息
		qualityBaseInfoRouter.DELETE("deleteQualityBaseInfoByIds", qualityBaseInfoApi.DeleteQualityBaseInfoByIds) // 批量删除质量要求信息
		qualityBaseInfoRouter.PUT("updateQualityBaseInfo", qualityBaseInfoApi.UpdateQualityBaseInfo)              // 更新质量要求信息
	}
	{
		qualityBaseInfoRouterWithoutRecord.GET("findQualityBaseInfo", qualityBaseInfoApi.FindQualityBaseInfo)       // 根据ID获取质量要求信息
		qualityBaseInfoRouterWithoutRecord.GET("getQualityBaseInfoList", qualityBaseInfoApi.GetQualityBaseInfoList) // 获取质量要求信息列表
	}
}
