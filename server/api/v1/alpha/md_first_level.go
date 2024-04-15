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

type MdFirstLevelApi struct {
}

var mdFirstLevelService = service.ServiceGroupApp.AlphaServiceGroup.MdFirstLevelService

// CreateMdFirstLevel 创建层级一
// @Tags MdFirstLevel
// @Summary 创建层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdFirstLevel true "创建层级一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdFirstLevel/createMdFirstLevel [post]
func (mdFirstLevelApi *MdFirstLevelApi) CreateMdFirstLevel(c *gin.Context) {
	var mdFirstLevel alpha.MdFirstLevel
	err := c.ShouldBindJSON(&mdFirstLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mdFirstLevel.CreatedBy = utils.GetUserID(c)

	if err := mdFirstLevelService.CreateMdFirstLevel(&mdFirstLevel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMdFirstLevel 删除层级一
// @Tags MdFirstLevel
// @Summary 删除层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdFirstLevel true "删除层级一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdFirstLevel/deleteMdFirstLevel [delete]
func (mdFirstLevelApi *MdFirstLevelApi) DeleteMdFirstLevel(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := mdFirstLevelService.DeleteMdFirstLevel(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMdFirstLevelByIds 批量删除层级一
// @Tags MdFirstLevel
// @Summary 批量删除层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mdFirstLevel/deleteMdFirstLevelByIds [delete]
func (mdFirstLevelApi *MdFirstLevelApi) DeleteMdFirstLevelByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := mdFirstLevelService.DeleteMdFirstLevelByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMdFirstLevel 更新层级一
// @Tags MdFirstLevel
// @Summary 更新层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdFirstLevel true "更新层级一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdFirstLevel/updateMdFirstLevel [put]
func (mdFirstLevelApi *MdFirstLevelApi) UpdateMdFirstLevel(c *gin.Context) {
	var mdFirstLevel alpha.MdFirstLevel
	err := c.ShouldBindJSON(&mdFirstLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mdFirstLevel.UpdatedBy = utils.GetUserID(c)

	if err := mdFirstLevelService.UpdateMdFirstLevel(mdFirstLevel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMdFirstLevel 用id查询层级一
// @Tags MdFirstLevel
// @Summary 用id查询层级一
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.MdFirstLevel true "用id查询层级一"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdFirstLevel/findMdFirstLevel [get]
func (mdFirstLevelApi *MdFirstLevelApi) FindMdFirstLevel(c *gin.Context) {
	ID := c.Query("ID")
	if remdFirstLevel, err := mdFirstLevelService.GetMdFirstLevel(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remdFirstLevel": remdFirstLevel}, c)
	}
}

// GetMdFirstLevelList 分页获取层级一列表
// @Tags MdFirstLevel
// @Summary 分页获取层级一列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.MdFirstLevelSearch true "分页获取层级一列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdFirstLevel/getMdFirstLevelList [get]
func (mdFirstLevelApi *MdFirstLevelApi) GetMdFirstLevelList(c *gin.Context) {
	var pageInfo alphaReq.MdFirstLevelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mdFirstLevelService.GetMdFirstLevelInfoList(pageInfo); err != nil {
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
