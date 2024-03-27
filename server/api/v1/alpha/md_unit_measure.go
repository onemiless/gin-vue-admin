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

type MdUnitMeasureApi struct {
}

var mdUnitMeasureService = service.ServiceGroupApp.AlphaServiceGroup.MdUnitMeasureService

// CreateMdUnitMeasure 创建品号单位设置
// @Tags MdUnitMeasure
// @Summary 创建品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdUnitMeasure true "创建品号单位设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdUnitMeasure/createMdUnitMeasure [post]
func (mdUnitMeasureApi *MdUnitMeasureApi) CreateMdUnitMeasure(c *gin.Context) {
	var mdUnitMeasure alpha.MdUnitMeasure
	err := c.ShouldBindJSON(&mdUnitMeasure)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mdUnitMeasureService.CreateMdUnitMeasure(&mdUnitMeasure); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMdUnitMeasure 删除品号单位设置
// @Tags MdUnitMeasure
// @Summary 删除品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdUnitMeasure true "删除品号单位设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdUnitMeasure/deleteMdUnitMeasure [delete]
func (mdUnitMeasureApi *MdUnitMeasureApi) DeleteMdUnitMeasure(c *gin.Context) {
	ID := c.Query("ID")
	if err := mdUnitMeasureService.DeleteMdUnitMeasure(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMdUnitMeasureByIds 批量删除品号单位设置
// @Tags MdUnitMeasure
// @Summary 批量删除品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mdUnitMeasure/deleteMdUnitMeasureByIds [delete]
func (mdUnitMeasureApi *MdUnitMeasureApi) DeleteMdUnitMeasureByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := mdUnitMeasureService.DeleteMdUnitMeasureByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMdUnitMeasure 更新品号单位设置
// @Tags MdUnitMeasure
// @Summary 更新品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdUnitMeasure true "更新品号单位设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdUnitMeasure/updateMdUnitMeasure [put]
func (mdUnitMeasureApi *MdUnitMeasureApi) UpdateMdUnitMeasure(c *gin.Context) {
	var mdUnitMeasure alpha.MdUnitMeasure
	err := c.ShouldBindJSON(&mdUnitMeasure)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mdUnitMeasureService.UpdateMdUnitMeasure(mdUnitMeasure); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMdUnitMeasure 用id查询品号单位设置
// @Tags MdUnitMeasure
// @Summary 用id查询品号单位设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.MdUnitMeasure true "用id查询品号单位设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdUnitMeasure/findMdUnitMeasure [get]
func (mdUnitMeasureApi *MdUnitMeasureApi) FindMdUnitMeasure(c *gin.Context) {
	ID := c.Query("ID")
	if remdUnitMeasure, err := mdUnitMeasureService.GetMdUnitMeasure(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remdUnitMeasure": remdUnitMeasure}, c)
	}
}

// GetMdUnitMeasureList 分页获取品号单位设置列表
// @Tags MdUnitMeasure
// @Summary 分页获取品号单位设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.MdUnitMeasureSearch true "分页获取品号单位设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdUnitMeasure/getMdUnitMeasureList [get]
func (mdUnitMeasureApi *MdUnitMeasureApi) GetMdUnitMeasureList(c *gin.Context) {
	var pageInfo alphaReq.MdUnitMeasureSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mdUnitMeasureService.GetMdUnitMeasureInfoList(pageInfo); err != nil {
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
