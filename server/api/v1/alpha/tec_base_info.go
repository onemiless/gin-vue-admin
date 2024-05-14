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

type TecBaseInfoApi struct {
}

var tecBaseInfoService = service.ServiceGroupApp.AlphaServiceGroup.TecBaseInfoService

// CreateTecBaseInfo 创建技术部基础信息
// @Tags TecBaseInfo
// @Summary 创建技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseInfo true "创建技术部基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tecBaseInfo/createTecBaseInfo [post]
func (tecBaseInfoApi *TecBaseInfoApi) CreateTecBaseInfo(c *gin.Context) {
	var tecBaseInfo alpha.TecBaseInfo
	err := c.ShouldBindJSON(&tecBaseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tecBaseInfo.CreatedBy = utils.GetUserID(c)
	id, err := tecBaseInfoService.CreateTecBaseInfo(&tecBaseInfo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(gin.H{"ID": id}, "创建成功", c)
	}
}

// DeleteTecBaseInfo 删除技术部基础信息
// @Tags TecBaseInfo
// @Summary 删除技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseInfo true "删除技术部基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tecBaseInfo/deleteTecBaseInfo [delete]
func (tecBaseInfoApi *TecBaseInfoApi) DeleteTecBaseInfo(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := tecBaseInfoService.DeleteTecBaseInfo(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTecBaseInfoByIds 批量删除技术部基础信息
// @Tags TecBaseInfo
// @Summary 批量删除技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tecBaseInfo/deleteTecBaseInfoByIds [delete]
func (tecBaseInfoApi *TecBaseInfoApi) DeleteTecBaseInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := tecBaseInfoService.DeleteTecBaseInfoByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTecBaseInfo 更新技术部基础信息
// @Tags TecBaseInfo
// @Summary 更新技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseInfo true "更新技术部基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tecBaseInfo/updateTecBaseInfo [put]
func (tecBaseInfoApi *TecBaseInfoApi) UpdateTecBaseInfo(c *gin.Context) {
	var tecBaseInfo alpha.TecBaseInfo
	err := c.ShouldBindJSON(&tecBaseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	tecBaseInfo.UpdatedBy = utils.GetUserID(c)

	if err := tecBaseInfoService.UpdateTecBaseInfo(tecBaseInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTecBaseInfo 用id查询技术部基础信息
// @Tags TecBaseInfo
// @Summary 用id查询技术部基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.TecBaseInfo true "用id查询技术部基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tecBaseInfo/findTecBaseInfo [get]
func (tecBaseInfoApi *TecBaseInfoApi) FindTecBaseInfo(c *gin.Context) {
	ID := c.Query("ID")
	UTN := c.Query("UTN")
	if UTN != "" {
		if retecBaseInfo, err := tecBaseInfoService.GetTecBaseInfoByUTN(UTN); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"retecBaseInfo": retecBaseInfo}, c)
		}
		return
	}
	if ID != "" {
		if retecBaseInfo, err := tecBaseInfoService.GetTecBaseInfo(ID); err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			response.OkWithData(gin.H{"retecBaseInfo": retecBaseInfo}, c)
		}
	}
}

// GetTecBaseInfoList 分页获取技术部基础信息列表
// @Tags TecBaseInfo
// @Summary 分页获取技术部基础信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.TecBaseInfoSearch true "分页获取技术部基础信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tecBaseInfo/getTecBaseInfoList [get]
func (tecBaseInfoApi *TecBaseInfoApi) GetTecBaseInfoList(c *gin.Context) {
	var pageInfo alphaReq.TecBaseInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tecBaseInfoService.GetTecBaseInfoInfoList(pageInfo); err != nil {
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
