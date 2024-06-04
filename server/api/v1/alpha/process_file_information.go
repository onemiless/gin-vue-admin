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

type ProcessFileInformationApi struct {
}

var processFileInformationService = service.ServiceGroupApp.AlphaServiceGroup.ProcessFileInformationService

// CreateProcessFileInformation 创建工艺文件信息
// @Tags ProcessFileInformation
// @Summary 创建工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ProcessFileInformation true "创建工艺文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /processFileInformation/createProcessFileInformation [post]
func (processFileInformationApi *ProcessFileInformationApi) CreateProcessFileInformation(c *gin.Context) {
	var processFileInformation alpha.ProcessFileInformation
	err := c.ShouldBindJSON(&processFileInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	processFileInformation.CreatedBy = utils.GetUserID(c)

	if err := processFileInformationService.CreateProcessFileInformation(&processFileInformation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProcessFileInformation 删除工艺文件信息
// @Tags ProcessFileInformation
// @Summary 删除工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ProcessFileInformation true "删除工艺文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /processFileInformation/deleteProcessFileInformation [delete]
func (processFileInformationApi *ProcessFileInformationApi) DeleteProcessFileInformation(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := processFileInformationService.DeleteProcessFileInformation(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProcessFileInformationByIds 批量删除工艺文件信息
// @Tags ProcessFileInformation
// @Summary 批量删除工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /processFileInformation/deleteProcessFileInformationByIds [delete]
func (processFileInformationApi *ProcessFileInformationApi) DeleteProcessFileInformationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := processFileInformationService.DeleteProcessFileInformationByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProcessFileInformation 更新工艺文件信息
// @Tags ProcessFileInformation
// @Summary 更新工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ProcessFileInformation true "更新工艺文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /processFileInformation/updateProcessFileInformation [put]
func (processFileInformationApi *ProcessFileInformationApi) UpdateProcessFileInformation(c *gin.Context) {
	var processFileInformation alpha.ProcessFileInformation
	err := c.ShouldBindJSON(&processFileInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	processFileInformation.UpdatedBy = utils.GetUserID(c)

	if err := processFileInformationService.UpdateProcessFileInformation(processFileInformation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProcessFileInformation 用id查询工艺文件信息
// @Tags ProcessFileInformation
// @Summary 用id查询工艺文件信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.ProcessFileInformation true "用id查询工艺文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /processFileInformation/findProcessFileInformation [get]
func (processFileInformationApi *ProcessFileInformationApi) FindProcessFileInformation(c *gin.Context) {
	ID := c.Query("ID")
	if reprocessFileInformation, err := processFileInformationService.GetProcessFileInformation(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reprocessFileInformation": reprocessFileInformation}, c)
	}
}

// GetProcessFileInformationList 分页获取工艺文件信息列表
// @Tags ProcessFileInformation
// @Summary 分页获取工艺文件信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.ProcessFileInformationSearch true "分页获取工艺文件信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /processFileInformation/getProcessFileInformationList [get]
func (processFileInformationApi *ProcessFileInformationApi) GetProcessFileInformationList(c *gin.Context) {
	var pageInfo alphaReq.ProcessFileInformationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := processFileInformationService.GetProcessFileInformationInfoList(pageInfo); err != nil {
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

// GetProcessFileInformationPublic 不需要鉴权的工艺文件信息接口
// @Tags ProcessFileInformation
// @Summary 不需要鉴权的工艺文件信息接口
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.ProcessFileInformationSearch true "分页获取工艺文件信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /processFileInformation/getProcessFileInformationPublic [get]
func (processFileInformationApi *ProcessFileInformationApi) GetProcessFileInformationPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的工艺文件信息接口信息",
	}, "获取成功", c)
}
