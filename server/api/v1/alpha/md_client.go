package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MdClientApi struct {
}

var mdClientService = service.ServiceGroupApp.AlphaServiceGroup.MdClientService

// CreateMdClient 创建客户信息
// @Tags MdClient
// @Summary 创建客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdClient true "创建客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdClient/createMdClient [post]
func (mdClientApi *MdClientApi) CreateMdClient(c *gin.Context) {
	var mdClient alpha.MdClient
	err := c.ShouldBindJSON(&mdClient)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mdClientService.CreateMdClient(&mdClient); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMdClient 删除客户信息
// @Tags MdClient
// @Summary 删除客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdClient true "删除客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdClient/deleteMdClient [delete]
func (mdClientApi *MdClientApi) DeleteMdClient(c *gin.Context) {
	ID := c.Query("ID")
	if err := mdClientService.DeleteMdClient(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMdClientByIds 批量删除客户信息
// @Tags MdClient
// @Summary 批量删除客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mdClient/deleteMdClientByIds [delete]
func (mdClientApi *MdClientApi) DeleteMdClientByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := mdClientService.DeleteMdClientByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMdClient 更新客户信息
// @Tags MdClient
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdClient true "更新客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdClient/updateMdClient [put]
func (mdClientApi *MdClientApi) UpdateMdClient(c *gin.Context) {
	var mdClient alpha.MdClient
	err := c.ShouldBindJSON(&mdClient)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mdClientService.UpdateMdClient(mdClient); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMdClient 用id查询客户信息
// @Tags MdClient
// @Summary 用id查询客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.MdClient true "用id查询客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdClient/findMdClient [get]
func (mdClientApi *MdClientApi) FindMdClient(c *gin.Context) {
	ID := c.Query("ID")
	if remdClient, err := mdClientService.GetMdClient(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remdClient": remdClient}, c)
	}
}

// GetMdClientList 分页获取客户信息列表
// @Tags MdClient
// @Summary 分页获取客户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.MdClientSearch true "分页获取客户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdClient/getMdClientList [get]
func (mdClientApi *MdClientApi) GetMdClientList(c *gin.Context) {
	var pageInfo alphaReq.MdClientSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mdClientService.GetMdClientInfoList(pageInfo); err != nil {
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
