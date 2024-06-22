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

type EngineeringChangeApi struct {
}

var engineeringChangeService = service.ServiceGroupApp.AlphaServiceGroup.EngineeringChangeService

// CreateEngineeringChange 创建工程变更信息
// @Tags EngineeringChange
// @Summary 创建工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.EngineeringChange true "创建工程变更信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /engineeringChange/createEngineeringChange [post]
func (engineeringChangeApi *EngineeringChangeApi) CreateEngineeringChange(c *gin.Context) {
	var engineeringChange alpha.EngineeringChange
	err := c.ShouldBindJSON(&engineeringChange)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	engineeringChange.CreatedBy = utils.GetUserID(c)

	if err := engineeringChangeService.CreateEngineeringChange(&engineeringChange); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEngineeringChange 删除工程变更信息
// @Tags EngineeringChange
// @Summary 删除工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.EngineeringChange true "删除工程变更信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /engineeringChange/deleteEngineeringChange [delete]
func (engineeringChangeApi *EngineeringChangeApi) DeleteEngineeringChange(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := engineeringChangeService.DeleteEngineeringChange(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEngineeringChangeByIds 批量删除工程变更信息
// @Tags EngineeringChange
// @Summary 批量删除工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /engineeringChange/deleteEngineeringChangeByIds [delete]
func (engineeringChangeApi *EngineeringChangeApi) DeleteEngineeringChangeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := engineeringChangeService.DeleteEngineeringChangeByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEngineeringChange 更新工程变更信息
// @Tags EngineeringChange
// @Summary 更新工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.EngineeringChange true "更新工程变更信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /engineeringChange/updateEngineeringChange [put]
func (engineeringChangeApi *EngineeringChangeApi) UpdateEngineeringChange(c *gin.Context) {
	var engineeringChange alpha.EngineeringChange
	err := c.ShouldBindJSON(&engineeringChange)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	engineeringChange.UpdatedBy = utils.GetUserID(c)

	if err := engineeringChangeService.UpdateEngineeringChange(engineeringChange); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEngineeringChange 用id查询工程变更信息
// @Tags EngineeringChange
// @Summary 用id查询工程变更信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.EngineeringChange true "用id查询工程变更信息"
// @Success 200 {object} response.Response{data=object{reengineeringChange=alpha.EngineeringChange},msg=string} "查询成功"
// @Router /engineeringChange/findEngineeringChange [get]
func (engineeringChangeApi *EngineeringChangeApi) FindEngineeringChange(c *gin.Context) {
	ID := c.Query("ID")
	if reengineeringChange, err := engineeringChangeService.GetEngineeringChange(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reengineeringChange, c)
	}
}

// GetEngineeringChangeList 分页获取工程变更信息列表
// @Tags EngineeringChange
// @Summary 分页获取工程变更信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.EngineeringChangeSearch true "分页获取工程变更信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /engineeringChange/getEngineeringChangeList [get]
func (engineeringChangeApi *EngineeringChangeApi) GetEngineeringChangeList(c *gin.Context) {
	var pageInfo alphaReq.EngineeringChangeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := engineeringChangeService.GetEngineeringChangeInfoList(pageInfo); err != nil {
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

// GetEngineeringChangePublic 不需要鉴权的工程变更信息接口
// @Tags EngineeringChange
// @Summary 不需要鉴权的工程变更信息接口
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.EngineeringChangeSearch true "分页获取工程变更信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /engineeringChange/getEngineeringChangePublic [get]
func (engineeringChangeApi *EngineeringChangeApi) GetEngineeringChangePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的工程变更信息接口信息",
	}, "获取成功", c)
}
