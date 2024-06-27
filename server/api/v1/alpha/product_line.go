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

type ProductLineApi struct{}

var productLineService = service.ServiceGroupApp.AlphaServiceGroup.ProductLineService

// CreateProductLine 创建产线信息
// @Tags ProductLine
// @Summary 创建产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ProductLine true "创建产线信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /productLine/createProductLine [post]
func (productLineApi *ProductLineApi) CreateProductLine(c *gin.Context) {
	var productLine alpha.ProductLine
	err := c.ShouldBindJSON(&productLine)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	productLine.CreatedBy = utils.GetUserID(c)

	if err := productLineService.CreateProductLine(&productLine); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProductLine 删除产线信息
// @Tags ProductLine
// @Summary 删除产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ProductLine true "删除产线信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /productLine/deleteProductLine [delete]
func (productLineApi *ProductLineApi) DeleteProductLine(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := productLineService.DeleteProductLine(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProductLineByIds 批量删除产线信息
// @Tags ProductLine
// @Summary 批量删除产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /productLine/deleteProductLineByIds [delete]
func (productLineApi *ProductLineApi) DeleteProductLineByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := productLineService.DeleteProductLineByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProductLine 更新产线信息
// @Tags ProductLine
// @Summary 更新产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ProductLine true "更新产线信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /productLine/updateProductLine [put]
func (productLineApi *ProductLineApi) UpdateProductLine(c *gin.Context) {
	var productLine alpha.ProductLine
	err := c.ShouldBindJSON(&productLine)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	productLine.UpdatedBy = utils.GetUserID(c)

	if err := productLineService.UpdateProductLine(productLine); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProductLine 用id查询产线信息
// @Tags ProductLine
// @Summary 用id查询产线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.ProductLine true "用id查询产线信息"
// @Success 200 {object} response.Response{data=object{reproductLine=alpha.ProductLine},msg=string} "查询成功"
// @Router /productLine/findProductLine [get]
func (productLineApi *ProductLineApi) FindProductLine(c *gin.Context) {
	ID := c.Query("ID")
	if reproductLine, err := productLineService.GetProductLine(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reproductLine, c)
	}
}

// GetProductLineList 分页获取产线信息列表
// @Tags ProductLine
// @Summary 分页获取产线信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.ProductLineSearch true "分页获取产线信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /productLine/getProductLineList [get]
func (productLineApi *ProductLineApi) GetProductLineList(c *gin.Context) {
	var pageInfo alphaReq.ProductLineSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := productLineService.GetProductLineInfoList(pageInfo); err != nil {
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

// GetProductLinePublic 不需要鉴权的产线信息接口
// @Tags ProductLine
// @Summary 不需要鉴权的产线信息接口
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.ProductLineSearch true "分页获取产线信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /productLine/getProductLinePublic [get]
func (productLineApi *ProductLineApi) GetProductLinePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的产线信息接口信息",
	}, "获取成功", c)
}
