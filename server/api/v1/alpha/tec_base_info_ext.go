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

type TecBaseInfoExtApi struct {
}

var tecBaseInfoExtService = service.ServiceGroupApp.AlphaServiceGroup.TecBaseInfoExtService

// CreateTecBaseInfoExt 创建零件基础信息扩展
// @Tags TecBaseInfoExt
// @Summary 创建零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseInfoExt true "创建零件基础信息扩展"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tecBaseInfoExt/createTecBaseInfoExt [post]
func (tecBaseInfoExtApi *TecBaseInfoExtApi) CreateTecBaseInfoExt(c *gin.Context) {
	var tecBaseInfoExt alpha.TecBaseInfoExt
	err := c.ShouldBindJSON(&tecBaseInfoExt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tecBaseInfoExt.CreatedBy = utils.GetUserID(c)

	if err := tecBaseInfoExtService.CreateTecBaseInfoExt(&tecBaseInfoExt); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTecBaseInfoExt 删除零件基础信息扩展
// @Tags TecBaseInfoExt
// @Summary 删除零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseInfoExt true "删除零件基础信息扩展"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseInfoExt/deleteTecBaseInfoExt [delete]
func (tecBaseInfoExtApi *TecBaseInfoExtApi) DeleteTecBaseInfoExt(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := tecBaseInfoExtService.DeleteTecBaseInfoExt(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTecBaseInfoExtByIds 批量删除零件基础信息扩展
// @Tags TecBaseInfoExt
// @Summary 批量删除零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tecBaseInfoExt/deleteTecBaseInfoExtByIds [delete]
func (tecBaseInfoExtApi *TecBaseInfoExtApi) DeleteTecBaseInfoExtByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := tecBaseInfoExtService.DeleteTecBaseInfoExtByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTecBaseInfoExt 更新零件基础信息扩展
// @Tags TecBaseInfoExt
// @Summary 更新零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseInfoExt true "更新零件基础信息扩展"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tecBaseInfoExt/updateTecBaseInfoExt [put]
func (tecBaseInfoExtApi *TecBaseInfoExtApi) UpdateTecBaseInfoExt(c *gin.Context) {
	var tecBaseInfoExt alpha.TecBaseInfoExt
	err := c.ShouldBindJSON(&tecBaseInfoExt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tecBaseInfoExt.UpdatedBy = utils.GetUserID(c)

	if err := tecBaseInfoExtService.UpdateTecBaseInfoExt(tecBaseInfoExt); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTecBaseInfoExt 用id查询零件基础信息扩展
// @Tags TecBaseInfoExt
// @Summary 用id查询零件基础信息扩展
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.TecBaseInfoExt true "用id查询零件基础信息扩展"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tecBaseInfoExt/findTecBaseInfoExt [get]
func (tecBaseInfoExtApi *TecBaseInfoExtApi) FindTecBaseInfoExt(c *gin.Context) {
	ID := c.Query("ID")
	if retecBaseInfoExt, err := tecBaseInfoExtService.GetTecBaseInfoExt(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retecBaseInfoExt": retecBaseInfoExt}, c)
	}
}

// GetTecBaseInfoExtList 分页获取零件基础信息扩展列表
// @Tags TecBaseInfoExt
// @Summary 分页获取零件基础信息扩展列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.TecBaseInfoExtSearch true "分页获取零件基础信息扩展列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tecBaseInfoExt/getTecBaseInfoExtList [get]
func (tecBaseInfoExtApi *TecBaseInfoExtApi) GetTecBaseInfoExtList(c *gin.Context) {
	var pageInfo alphaReq.TecBaseInfoExtSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tecBaseInfoExtService.GetTecBaseInfoExtInfoList(pageInfo); err != nil {
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

// CheckIsDuplicate 检查是否有重复
// @Tags TecBaseInfoExt
// @Summary 检查是否有重复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.TecBaseInfoExtSearch true "分页获取零件基础信息扩展列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tecBaseInfoExt/checkIsDuplicate [post]
func (tecBaseInfoExtApi *TecBaseInfoExtApi) CheckIsDuplicate(c *gin.Context) {
	var pageInfo alphaReq.TecBaseInfoExtSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if result, err := tecBaseInfoExtService.CheckIsDuplicate(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(result, "获取成功", c)
	}

}
