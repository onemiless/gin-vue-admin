package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestFileAndImgRouter struct {
}

// InitTestFileAndImgRouter 初始化 TestFileAndImg 路由信息
func (s *TestFileAndImgRouter) InitTestFileAndImgRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	testFileAndImgRouter := Router.Group("testFileAndImg").Use(middleware.OperationRecord())
	testFileAndImgRouterWithoutRecord := Router.Group("testFileAndImg")
	testFileAndImgRouterWithoutAuth := PublicRouter.Group("testFileAndImg")

	var testFileAndImgApi = v1.ApiGroupApp.AlphaApiGroup.TestFileAndImgApi
	{
		testFileAndImgRouter.POST("createTestFileAndImg", testFileAndImgApi.CreateTestFileAndImg)             // 新建TestFileAndImg
		testFileAndImgRouter.DELETE("deleteTestFileAndImg", testFileAndImgApi.DeleteTestFileAndImg)           // 删除TestFileAndImg
		testFileAndImgRouter.DELETE("deleteTestFileAndImgByIds", testFileAndImgApi.DeleteTestFileAndImgByIds) // 批量删除TestFileAndImg
		testFileAndImgRouter.PUT("updateTestFileAndImg", testFileAndImgApi.UpdateTestFileAndImg)              // 更新TestFileAndImg
	}
	{
		testFileAndImgRouterWithoutRecord.GET("findTestFileAndImg", testFileAndImgApi.FindTestFileAndImg)       // 根据ID获取TestFileAndImg
		testFileAndImgRouterWithoutRecord.GET("getTestFileAndImgList", testFileAndImgApi.GetTestFileAndImgList) // 获取TestFileAndImg列表
	}
	{
		testFileAndImgRouterWithoutAuth.GET("getTestFileAndImgPublic", testFileAndImgApi.GetTestFileAndImgPublic) // 获取TestFileAndImg列表
	}
}
