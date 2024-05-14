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

type QualityBaseInfoApi struct {
}

var qualityBaseInfoService = service.ServiceGroupApp.AlphaServiceGroup.QualityBaseInfoService

// CreateQualityBaseInfo 创建质量要求信息
// @Tags QualityBaseInfo
// @Summary 创建质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.QualityBaseInfo true "创建质量要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /qualityBaseInfo/createQualityBaseInfo [post]
func (qualityBaseInfoApi *QualityBaseInfoApi) CreateQualityBaseInfo(c *gin.Context) {
	var qualityBaseInfo alpha.QualityBaseInfo
	err := c.ShouldBindJSON(&qualityBaseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	qualityBaseInfo.CreatedBy = utils.GetUserID(c)

	if err := qualityBaseInfoService.CreateQualityBaseInfo(&qualityBaseInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteQualityBaseInfo 删除质量要求信息
// @Tags QualityBaseInfo
// @Summary 删除质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.QualityBaseInfo true "删除质量要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qualityBaseInfo/deleteQualityBaseInfo [delete]
func (qualityBaseInfoApi *QualityBaseInfoApi) DeleteQualityBaseInfo(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := qualityBaseInfoService.DeleteQualityBaseInfo(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteQualityBaseInfoByIds 批量删除质量要求信息
// @Tags QualityBaseInfo
// @Summary 批量删除质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /qualityBaseInfo/deleteQualityBaseInfoByIds [delete]
func (qualityBaseInfoApi *QualityBaseInfoApi) DeleteQualityBaseInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := qualityBaseInfoService.DeleteQualityBaseInfoByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateQualityBaseInfo 更新质量要求信息
// @Tags QualityBaseInfo
// @Summary 更新质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.QualityBaseInfo true "更新质量要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /qualityBaseInfo/updateQualityBaseInfo [put]
func (qualityBaseInfoApi *QualityBaseInfoApi) UpdateQualityBaseInfo(c *gin.Context) {
	var qualityBaseInfo alpha.QualityBaseInfo
	err := c.ShouldBindJSON(&qualityBaseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	qualityBaseInfo.UpdatedBy = utils.GetUserID(c)

	if err := qualityBaseInfoService.UpdateQualityBaseInfo(qualityBaseInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindQualityBaseInfo 用id查询质量要求信息
// @Tags QualityBaseInfo
// @Summary 用id查询质量要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.QualityBaseInfo true "用id查询质量要求信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /qualityBaseInfo/findQualityBaseInfo [get]
func (qualityBaseInfoApi *QualityBaseInfoApi) FindQualityBaseInfo(c *gin.Context) {
	ID := c.Query("ID")
	if requalityBaseInfo, err := qualityBaseInfoService.GetQualityBaseInfo(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requalityBaseInfo": requalityBaseInfo}, c)
	}
}

// GetQualityBaseInfoList 分页获取质量要求信息列表
// @Tags QualityBaseInfo
// @Summary 分页获取质量要求信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.QualityBaseInfoSearch true "分页获取质量要求信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /qualityBaseInfo/getQualityBaseInfoList [get]
func (qualityBaseInfoApi *QualityBaseInfoApi) GetQualityBaseInfoList(c *gin.Context) {
	var pageInfo alphaReq.QualityBaseInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := qualityBaseInfoService.GetQualityBaseInfoInfoList(pageInfo); err != nil {
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
