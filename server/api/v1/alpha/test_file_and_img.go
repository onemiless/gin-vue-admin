package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestFileAndImgApi struct {
}

var testFileAndImgService = service.ServiceGroupApp.AlphaServiceGroup.TestFileAndImgService

// CreateTestFileAndImg 创建TestFileAndImg
// @Tags TestFileAndImg
// @Summary 创建TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TestFileAndImg true "创建TestFileAndImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /testFileAndImg/createTestFileAndImg [post]
func (testFileAndImgApi *TestFileAndImgApi) CreateTestFileAndImg(c *gin.Context) {
	var testFileAndImg alpha.TestFileAndImg
	err := c.ShouldBindJSON(&testFileAndImg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	testFileAndImg.CreatedBy = utils.GetUserID(c)

	if err := testFileAndImgService.CreateTestFileAndImg(&testFileAndImg); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestFileAndImg 删除TestFileAndImg
// @Tags TestFileAndImg
// @Summary 删除TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TestFileAndImg true "删除TestFileAndImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testFileAndImg/deleteTestFileAndImg [delete]
func (testFileAndImgApi *TestFileAndImgApi) DeleteTestFileAndImg(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := testFileAndImgService.DeleteTestFileAndImg(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestFileAndImgByIds 批量删除TestFileAndImg
// @Tags TestFileAndImg
// @Summary 批量删除TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testFileAndImg/deleteTestFileAndImgByIds [delete]
func (testFileAndImgApi *TestFileAndImgApi) DeleteTestFileAndImgByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := testFileAndImgService.DeleteTestFileAndImgByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestFileAndImg 更新TestFileAndImg
// @Tags TestFileAndImg
// @Summary 更新TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TestFileAndImg true "更新TestFileAndImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testFileAndImg/updateTestFileAndImg [put]
func (testFileAndImgApi *TestFileAndImgApi) UpdateTestFileAndImg(c *gin.Context) {
	var testFileAndImg alpha.TestFileAndImg
	err := c.ShouldBindJSON(&testFileAndImg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	testFileAndImg.UpdatedBy = utils.GetUserID(c)

	if err := testFileAndImgService.UpdateTestFileAndImg(testFileAndImg); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestFileAndImg 用id查询TestFileAndImg
// @Tags TestFileAndImg
// @Summary 用id查询TestFileAndImg
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.TestFileAndImg true "用id查询TestFileAndImg"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testFileAndImg/findTestFileAndImg [get]
func (testFileAndImgApi *TestFileAndImgApi) FindTestFileAndImg(c *gin.Context) {
	ID := c.Query("ID")
	if retestFileAndImg, err := testFileAndImgService.GetTestFileAndImg(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestFileAndImg": retestFileAndImg}, c)
	}
}

// GetTestFileAndImgList 分页获取TestFileAndImg列表
// @Tags TestFileAndImg
// @Summary 分页获取TestFileAndImg列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.TestFileAndImgSearch true "分页获取TestFileAndImg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testFileAndImg/getTestFileAndImgList [get]
func (testFileAndImgApi *TestFileAndImgApi) GetTestFileAndImgList(c *gin.Context) {
	var pageInfo alphaReq.TestFileAndImgSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testFileAndImgService.GetTestFileAndImgInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetTestFileAndImgPublic 不需要鉴权的TestFileAndImg接口
// @Tags TestFileAndImg
// @Summary 不需要鉴权的TestFileAndImg接口
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.TestFileAndImgSearch true "分页获取TestFileAndImg列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testFileAndImg/getTestFileAndImgPublic [get]
func (testFileAndImgApi *TestFileAndImgApi) GetTestFileAndImgPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的TestFileAndImg接口信息",
	}, "获取成功", c)
}
