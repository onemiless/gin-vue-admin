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

type COPMAApi struct {
}

var copmaService = service.ServiceGroupApp.AlphaServiceGroup.COPMAService

// CreateCOPMA 创建ERP客户信息
// @Tags COPMA
// @Summary 创建ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.COPMA true "创建ERP客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /copma/createCOPMA [post]
func (copmaApi *COPMAApi) CreateCOPMA(c *gin.Context) {
	var copma alpha.COPMA
	err := c.ShouldBindJSON(&copma)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := copmaService.CreateCOPMA(&copma); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCOPMA 删除ERP客户信息
// @Tags COPMA
// @Summary 删除ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.COPMA true "删除ERP客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /copma/deleteCOPMA [delete]
func (copmaApi *COPMAApi) DeleteCOPMA(c *gin.Context) {
	MA001 := c.Query("MA001")
	if err := copmaService.DeleteCOPMA(MA001); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCOPMAByIds 批量删除ERP客户信息
// @Tags COPMA
// @Summary 批量删除ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /copma/deleteCOPMAByIds [delete]
func (copmaApi *COPMAApi) DeleteCOPMAByIds(c *gin.Context) {
	MA001s := c.QueryArray("MA001s[]")
	if err := copmaService.DeleteCOPMAByIds(MA001s); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCOPMA 更新ERP客户信息
// @Tags COPMA
// @Summary 更新ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.COPMA true "更新ERP客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /copma/updateCOPMA [put]
func (copmaApi *COPMAApi) UpdateCOPMA(c *gin.Context) {
	var copma alpha.COPMA
	err := c.ShouldBindJSON(&copma)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := copmaService.UpdateCOPMA(copma); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCOPMA 用id查询ERP客户信息
// @Tags COPMA
// @Summary 用id查询ERP客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.COPMA true "用id查询ERP客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /copma/findCOPMA [get]
func (copmaApi *COPMAApi) FindCOPMA(c *gin.Context) {
	MA001 := c.Query("MA001")
	if recopma, err := copmaService.GetCOPMA(MA001); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recopma": recopma}, c)
	}
}

// GetCOPMAList 分页获取ERP客户信息列表
// @Tags COPMA
// @Summary 分页获取ERP客户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.COPMASearch true "分页获取ERP客户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /copma/getCOPMAList [get]
func (copmaApi *COPMAApi) GetCOPMAList(c *gin.Context) {
	var pageInfo alphaReq.COPMASearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := copmaService.GetCOPMAInfoList(pageInfo); err != nil {
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

// GetCOPMAPublic 不需要鉴权的ERP客户信息接口
// @Tags COPMA
// @Summary 不需要鉴权的ERP客户信息接口
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.COPMASearch true "分页获取ERP客户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /copma/getCOPMAPublic [get]
func (copmaApi *COPMAApi) GetCOPMAPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的ERP客户信息接口信息",
	}, "获取成功", c)
}
