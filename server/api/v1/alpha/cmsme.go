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

type CMSMEApi struct {
}

var cmsmeService = service.ServiceGroupApp.AlphaServiceGroup.CMSMEService

// CreateCMSME 创建ERP部门
// @Tags CMSME
// @Summary 创建ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.CMSME true "创建ERP部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsme/createCMSME [post]
func (cmsmeApi *CMSMEApi) CreateCMSME(c *gin.Context) {
	var cmsme alpha.CMSME
	err := c.ShouldBindJSON(&cmsme)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsmeService.CreateCMSME(&cmsme); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCMSME 删除ERP部门
// @Tags CMSME
// @Summary 删除ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.CMSME true "删除ERP部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsme/deleteCMSME [delete]
func (cmsmeApi *CMSMEApi) DeleteCMSME(c *gin.Context) {
	ME001 := c.Query("ME001")
	if err := cmsmeService.DeleteCMSME(ME001); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCMSMEByIds 批量删除ERP部门
// @Tags CMSME
// @Summary 批量删除ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsme/deleteCMSMEByIds [delete]
func (cmsmeApi *CMSMEApi) DeleteCMSMEByIds(c *gin.Context) {
	ME001s := c.QueryArray("ME001s[]")
	if err := cmsmeService.DeleteCMSMEByIds(ME001s); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCMSME 更新ERP部门
// @Tags CMSME
// @Summary 更新ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.CMSME true "更新ERP部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsme/updateCMSME [put]
func (cmsmeApi *CMSMEApi) UpdateCMSME(c *gin.Context) {
	var cmsme alpha.CMSME
	err := c.ShouldBindJSON(&cmsme)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsmeService.UpdateCMSME(cmsme); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCMSME 用id查询ERP部门
// @Tags CMSME
// @Summary 用id查询ERP部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.CMSME true "用id查询ERP部门"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsme/findCMSME [get]
func (cmsmeApi *CMSMEApi) FindCMSME(c *gin.Context) {
	ME001 := c.Query("ME001")
	if recmsme, err := cmsmeService.GetCMSME(ME001); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsme": recmsme}, c)
	}
}

// GetCMSMEList 分页获取ERP部门列表
// @Tags CMSME
// @Summary 分页获取ERP部门列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.CMSMESearch true "分页获取ERP部门列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsme/getCMSMEList [get]
func (cmsmeApi *CMSMEApi) GetCMSMEList(c *gin.Context) {
	var pageInfo alphaReq.CMSMESearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmsmeService.GetCMSMEInfoList(pageInfo); err != nil {
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
