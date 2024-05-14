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

type CostCollectionApi struct {
}

var costCollectionService = service.ServiceGroupApp.AlphaServiceGroup.CostCollectionService

// CreateCostCollection 创建成本信息收集信息
// @Tags CostCollection
// @Summary 创建成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.CostCollection true "创建成本信息收集信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /costCollection/createCostCollection [post]
func (costCollectionApi *CostCollectionApi) CreateCostCollection(c *gin.Context) {
	var costCollection alpha.CostCollection
	err := c.ShouldBindJSON(&costCollection)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	costCollection.CreatedBy = utils.GetUserID(c)

	if err := costCollectionService.CreateCostCollection(&costCollection); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCostCollection 删除成本信息收集信息
// @Tags CostCollection
// @Summary 删除成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.CostCollection true "删除成本信息收集信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /costCollection/deleteCostCollection [delete]
func (costCollectionApi *CostCollectionApi) DeleteCostCollection(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := costCollectionService.DeleteCostCollection(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCostCollectionByIds 批量删除成本信息收集信息
// @Tags CostCollection
// @Summary 批量删除成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /costCollection/deleteCostCollectionByIds [delete]
func (costCollectionApi *CostCollectionApi) DeleteCostCollectionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := costCollectionService.DeleteCostCollectionByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCostCollection 更新成本信息收集信息
// @Tags CostCollection
// @Summary 更新成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.CostCollection true "更新成本信息收集信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /costCollection/updateCostCollection [put]
func (costCollectionApi *CostCollectionApi) UpdateCostCollection(c *gin.Context) {
	var costCollection alpha.CostCollection
	err := c.ShouldBindJSON(&costCollection)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	costCollection.UpdatedBy = utils.GetUserID(c)

	if err := costCollectionService.UpdateCostCollection(costCollection); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCostCollection 用id查询成本信息收集信息
// @Tags CostCollection
// @Summary 用id查询成本信息收集信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.CostCollection true "用id查询成本信息收集信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /costCollection/findCostCollection [get]
func (costCollectionApi *CostCollectionApi) FindCostCollection(c *gin.Context) {
	ID := c.Query("ID")
	if recostCollection, err := costCollectionService.GetCostCollection(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recostCollection": recostCollection}, c)
	}
}

// GetCostCollectionList 分页获取成本信息收集信息列表
// @Tags CostCollection
// @Summary 分页获取成本信息收集信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.CostCollectionSearch true "分页获取成本信息收集信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /costCollection/getCostCollectionList [get]
func (costCollectionApi *CostCollectionApi) GetCostCollectionList(c *gin.Context) {
	var pageInfo alphaReq.CostCollectionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := costCollectionService.GetCostCollectionInfoList(pageInfo); err != nil {
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

// GetCostCollectionNumer 获取流水号
// @Tags CostCollection
// @Summary 获取流水号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.CostCollectionSearch true "获取流水号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /costCollection/getCostCollectionList [get]
func (costCollectionApi *CostCollectionApi) GetCostCollectionNumer(c *gin.Context) {
	if number, err := costCollectionService.GetCostCollectionNumber(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"number": number}, c)
	}
}
