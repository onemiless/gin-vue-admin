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

type MassProductionTransferApi struct {
}

var massProductionTransferService = service.ServiceGroupApp.AlphaServiceGroup.MassProductionTransferService

// CreateMassProductionTransfer 创建量产转移
// @Tags MassProductionTransfer
// @Summary 创建量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MassProductionTransfer true "创建量产转移"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /massProductionTransfer/createMassProductionTransfer [post]
func (massProductionTransferApi *MassProductionTransferApi) CreateMassProductionTransfer(c *gin.Context) {
	var massProductionTransfer alpha.MassProductionTransfer
	err := c.ShouldBindJSON(&massProductionTransfer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	massProductionTransfer.CreatedBy = utils.GetUserID(c)

	if err := massProductionTransferService.CreateMassProductionTransfer(&massProductionTransfer); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMassProductionTransfer 删除量产转移
// @Tags MassProductionTransfer
// @Summary 删除量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MassProductionTransfer true "删除量产转移"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /massProductionTransfer/deleteMassProductionTransfer [delete]
func (massProductionTransferApi *MassProductionTransferApi) DeleteMassProductionTransfer(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := massProductionTransferService.DeleteMassProductionTransfer(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMassProductionTransferByIds 批量删除量产转移
// @Tags MassProductionTransfer
// @Summary 批量删除量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /massProductionTransfer/deleteMassProductionTransferByIds [delete]
func (massProductionTransferApi *MassProductionTransferApi) DeleteMassProductionTransferByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := massProductionTransferService.DeleteMassProductionTransferByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMassProductionTransfer 更新量产转移
// @Tags MassProductionTransfer
// @Summary 更新量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MassProductionTransfer true "更新量产转移"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /massProductionTransfer/updateMassProductionTransfer [put]
func (massProductionTransferApi *MassProductionTransferApi) UpdateMassProductionTransfer(c *gin.Context) {
	var massProductionTransfer alpha.MassProductionTransfer
	err := c.ShouldBindJSON(&massProductionTransfer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	massProductionTransfer.UpdatedBy = utils.GetUserID(c)

	if err := massProductionTransferService.UpdateMassProductionTransfer(massProductionTransfer); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMassProductionTransfer 用id查询量产转移
// @Tags MassProductionTransfer
// @Summary 用id查询量产转移
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.MassProductionTransfer true "用id查询量产转移"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /massProductionTransfer/findMassProductionTransfer [get]
func (massProductionTransferApi *MassProductionTransferApi) FindMassProductionTransfer(c *gin.Context) {
	ID := c.Query("ID")
	if remassProductionTransfer, err := massProductionTransferService.GetMassProductionTransfer(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remassProductionTransfer": remassProductionTransfer}, c)
	}
}

// GetMassProductionTransferList 分页获取量产转移列表
// @Tags MassProductionTransfer
// @Summary 分页获取量产转移列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.MassProductionTransferSearch true "分页获取量产转移列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /massProductionTransfer/getMassProductionTransferList [get]
func (massProductionTransferApi *MassProductionTransferApi) GetMassProductionTransferList(c *gin.Context) {
	var pageInfo alphaReq.MassProductionTransferSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := massProductionTransferService.GetMassProductionTransferInfoList(pageInfo); err != nil {
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

// GetMassProductionTransferPublic 不需要鉴权的量产转移接口
// @Tags MassProductionTransfer
// @Summary 不需要鉴权的量产转移接口
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.MassProductionTransferSearch true "分页获取量产转移列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /massProductionTransfer/getMassProductionTransferPublic [get]
func (massProductionTransferApi *MassProductionTransferApi) GetMassProductionTransferPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的量产转移接口信息",
	}, "获取成功", c)
}
