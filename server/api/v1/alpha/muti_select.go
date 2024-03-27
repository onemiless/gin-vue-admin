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

type MutiSelectApi struct {
}

var mutiSelectService = service.ServiceGroupApp.AlphaServiceGroup.MutiSelectService

// CreateMutiSelect 创建多级选择信息
// @Tags MutiSelect
// @Summary 创建多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MutiSelect true "创建多级选择信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mutiSelect/createMutiSelect [post]
func (mutiSelectApi *MutiSelectApi) CreateMutiSelect(c *gin.Context) {
	var mutiSelect alpha.MutiSelect
	err := c.ShouldBindJSON(&mutiSelect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mutiSelectService.CreateMutiSelect(&mutiSelect); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMutiSelect 删除多级选择信息
// @Tags MutiSelect
// @Summary 删除多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MutiSelect true "删除多级选择信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mutiSelect/deleteMutiSelect [delete]
func (mutiSelectApi *MutiSelectApi) DeleteMutiSelect(c *gin.Context) {
	ID := c.Query("ID")
	if err := mutiSelectService.DeleteMutiSelect(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMutiSelectByIds 批量删除多级选择信息
// @Tags MutiSelect
// @Summary 批量删除多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mutiSelect/deleteMutiSelectByIds [delete]
func (mutiSelectApi *MutiSelectApi) DeleteMutiSelectByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := mutiSelectService.DeleteMutiSelectByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMutiSelect 更新多级选择信息
// @Tags MutiSelect
// @Summary 更新多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MutiSelect true "更新多级选择信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mutiSelect/updateMutiSelect [put]
func (mutiSelectApi *MutiSelectApi) UpdateMutiSelect(c *gin.Context) {
	var mutiSelect alpha.MutiSelect
	err := c.ShouldBindJSON(&mutiSelect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mutiSelectService.UpdateMutiSelect(mutiSelect); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMutiSelect 用id查询多级选择信息
// @Tags MutiSelect
// @Summary 用id查询多级选择信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.MutiSelect true "用id查询多级选择信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mutiSelect/findMutiSelect [get]
func (mutiSelectApi *MutiSelectApi) FindMutiSelect(c *gin.Context) {
	ID := c.Query("ID")
	if remutiSelect, err := mutiSelectService.GetMutiSelect(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remutiSelect": remutiSelect}, c)
	}
}

// GetMutiSelectList 分页获取多级选择信息列表
// @Tags MutiSelect
// @Summary 分页获取多级选择信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.MutiSelectSearch true "分页获取多级选择信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mutiSelect/getMutiSelectList [get]
func (mutiSelectApi *MutiSelectApi) GetMutiSelectList(c *gin.Context) {
	var pageInfo alphaReq.MutiSelectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mutiSelectService.GetMutiSelectInfoList(pageInfo); err != nil {
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
