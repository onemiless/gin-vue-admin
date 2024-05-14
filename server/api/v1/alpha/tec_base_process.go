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

type TecBaseProcessApi struct {
}

var tecBaseProcessService = service.ServiceGroupApp.AlphaServiceGroup.TecBaseProcessService

// CreateTecBaseProcess 创建工艺、设备及治具基本信息
// @Tags TecBaseProcess
// @Summary 创建工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseProcess true "创建工艺、设备及治具基本信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tecBaseProcess/createTecBaseProcess [post]
func (tecBaseProcessApi *TecBaseProcessApi) CreateTecBaseProcess(c *gin.Context) {
	var tecBaseProcess alpha.TecBaseProcess
	err := c.ShouldBindJSON(&tecBaseProcess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tecBaseProcess.CreatedBy = utils.GetUserID(c)

	if err := tecBaseProcessService.CreateTecBaseProcess(&tecBaseProcess); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTecBaseProcess 删除工艺、设备及治具基本信息
// @Tags TecBaseProcess
// @Summary 删除工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseProcess true "删除工艺、设备及治具基本信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseProcess/deleteTecBaseProcess [delete]
func (tecBaseProcessApi *TecBaseProcessApi) DeleteTecBaseProcess(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := tecBaseProcessService.DeleteTecBaseProcess(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTecBaseProcessByIds 批量删除工艺、设备及治具基本信息
// @Tags TecBaseProcess
// @Summary 批量删除工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tecBaseProcess/deleteTecBaseProcessByIds [delete]
func (tecBaseProcessApi *TecBaseProcessApi) DeleteTecBaseProcessByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := tecBaseProcessService.DeleteTecBaseProcessByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTecBaseProcess 更新工艺、设备及治具基本信息
// @Tags TecBaseProcess
// @Summary 更新工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseProcess true "更新工艺、设备及治具基本信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tecBaseProcess/updateTecBaseProcess [put]
func (tecBaseProcessApi *TecBaseProcessApi) UpdateTecBaseProcess(c *gin.Context) {
	var tecBaseProcess alpha.TecBaseProcess
	err := c.ShouldBindJSON(&tecBaseProcess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tecBaseProcess.UpdatedBy = utils.GetUserID(c)

	if err := tecBaseProcessService.UpdateTecBaseProcess(tecBaseProcess); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTecBaseProcess 用id查询工艺、设备及治具基本信息
// @Tags TecBaseProcess
// @Summary 用id查询工艺、设备及治具基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.TecBaseProcess true "用id查询工艺、设备及治具基本信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tecBaseProcess/findTecBaseProcess [get]
func (tecBaseProcessApi *TecBaseProcessApi) FindTecBaseProcess(c *gin.Context) {
	ID := c.Query("ID")
	ParenId := c.Query("ParentId")
	if ParenId != "" {
		if tecBaseProcess, err := tecBaseProcessService.GetTecBaseProcessParentID(ParenId); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"tecBaseProcess": tecBaseProcess}, c)
		}
	}
	if ID != "" {

		if retecBaseProcess, err := tecBaseProcessService.GetTecBaseProcess(ID); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"retecBaseProcess": retecBaseProcess}, c)
		}
	}
}

// GetTecBaseProcessList 分页获取工艺、设备及治具基本信息列表
// @Tags TecBaseProcess
// @Summary 分页获取工艺、设备及治具基本信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.TecBaseProcessSearch true "分页获取工艺、设备及治具基本信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tecBaseProcess/getTecBaseProcessList [get]
func (tecBaseProcessApi *TecBaseProcessApi) GetTecBaseProcessList(c *gin.Context) {
	var pageInfo alphaReq.TecBaseProcessSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tecBaseProcessService.GetTecBaseProcessInfoList(pageInfo); err != nil {
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
