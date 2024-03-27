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

type CMSXCApi struct {
}

var cmsxcService = service.ServiceGroupApp.AlphaServiceGroup.CMSXCService

// CreateCMSXC 创建车厂规范信息
// @Tags CMSXC
// @Summary 创建车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.CMSXC true "创建车厂规范信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmsxc/createCMSXC [post]
func (cmsxcApi *CMSXCApi) CreateCMSXC(c *gin.Context) {
	var cmsxc alpha.CMSXC
	err := c.ShouldBindJSON(&cmsxc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsxcService.CreateCMSXC(&cmsxc); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCMSXC 删除车厂规范信息
// @Tags CMSXC
// @Summary 删除车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.CMSXC true "删除车厂规范信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsxc/deleteCMSXC [delete]
func (cmsxcApi *CMSXCApi) DeleteCMSXC(c *gin.Context) {
	XC001 := c.Query("XC001")
	if err := cmsxcService.DeleteCMSXC(XC001); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCMSXCByIds 批量删除车厂规范信息
// @Tags CMSXC
// @Summary 批量删除车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsxc/deleteCMSXCByIds [delete]
func (cmsxcApi *CMSXCApi) DeleteCMSXCByIds(c *gin.Context) {
	XC001s := c.QueryArray("XC001s[]")
	if err := cmsxcService.DeleteCMSXCByIds(XC001s); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCMSXC 更新车厂规范信息
// @Tags CMSXC
// @Summary 更新车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.CMSXC true "更新车厂规范信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsxc/updateCMSXC [put]
func (cmsxcApi *CMSXCApi) UpdateCMSXC(c *gin.Context) {
	var cmsxc alpha.CMSXC
	err := c.ShouldBindJSON(&cmsxc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := cmsxcService.UpdateCMSXC(cmsxc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCMSXC 用id查询车厂规范信息
// @Tags CMSXC
// @Summary 用id查询车厂规范信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.CMSXC true "用id查询车厂规范信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsxc/findCMSXC [get]
func (cmsxcApi *CMSXCApi) FindCMSXC(c *gin.Context) {
	XC001 := c.Query("XC001")
	if recmsxc, err := cmsxcService.GetCMSXC(XC001); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsxc": recmsxc}, c)
	}
}

// GetCMSXCList 分页获取车厂规范信息列表
// @Tags CMSXC
// @Summary 分页获取车厂规范信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.CMSXCSearch true "分页获取车厂规范信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsxc/getCMSXCList [get]
func (cmsxcApi *CMSXCApi) GetCMSXCList(c *gin.Context) {
	var pageInfo alphaReq.CMSXCSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmsxcService.GetCMSXCInfoList(pageInfo); err != nil {
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
