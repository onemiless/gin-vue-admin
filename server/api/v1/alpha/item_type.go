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

type ItemTypeApi struct {
}

var itemtypeService = service.ServiceGroupApp.AlphaServiceGroup.ItemTypeService

// CreateItemType 创建零件类型
// @Tags ItemType
// @Summary 创建零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ItemType true "创建零件类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /itemtype/createItemType [post]
func (itemtypeApi *ItemTypeApi) CreateItemType(c *gin.Context) {
	var itemtype alpha.ItemType
	err := c.ShouldBindJSON(&itemtype)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := itemtypeService.CreateItemType(&itemtype); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteItemType 删除零件类型
// @Tags ItemType
// @Summary 删除零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ItemType true "删除零件类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /itemtype/deleteItemType [delete]
func (itemtypeApi *ItemTypeApi) DeleteItemType(c *gin.Context) {
	ID := c.Query("ID")
	if err := itemtypeService.DeleteItemType(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteItemTypeByIds 批量删除零件类型
// @Tags ItemType
// @Summary 批量删除零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /itemtype/deleteItemTypeByIds [delete]
func (itemtypeApi *ItemTypeApi) DeleteItemTypeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := itemtypeService.DeleteItemTypeByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateItemType 更新零件类型
// @Tags ItemType
// @Summary 更新零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ItemType true "更新零件类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /itemtype/updateItemType [put]
func (itemtypeApi *ItemTypeApi) UpdateItemType(c *gin.Context) {
	var itemtype alpha.ItemType
	err := c.ShouldBindJSON(&itemtype)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := itemtypeService.UpdateItemType(itemtype); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindItemType 用id查询零件类型
// @Tags ItemType
// @Summary 用id查询零件类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.ItemType true "用id查询零件类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /itemtype/findItemType [get]
func (itemtypeApi *ItemTypeApi) FindItemType(c *gin.Context) {
	ID := c.Query("ID")
	if reitemtype, err := itemtypeService.GetItemType(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reitemtype": reitemtype}, c)
	}
}

// GetItemTypeList 分页获取零件类型列表
// @Tags ItemType
// @Summary 分页获取零件类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.ItemTypeSearch true "分页获取零件类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /itemtype/getItemTypeList [get]
func (itemtypeApi *ItemTypeApi) GetItemTypeList(c *gin.Context) {
	var pageInfo alphaReq.ItemTypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := itemtypeService.GetItemTypeInfoList(pageInfo); err != nil {
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

// GetItemTypeListParent 分页获取零件类型列表
// @Tags ItemType
// @Summary 分页获取零件类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.ItemTypeSearch true "分页获取零件类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /itemtype/getItemTypeList [get]
func (itemtypeApi *ItemTypeApi) GetItemTypeListParent(c *gin.Context) {
	var pageInfo alphaReq.ItemTypeSearch

	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := itemtypeService.GetItemTypeInfoParentList(pageInfo); err != nil {
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
