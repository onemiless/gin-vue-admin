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

type MdSecondLevelApi struct {
}

var mdSecondLevelService = service.ServiceGroupApp.AlphaServiceGroup.MdSecondLevelService

// CreateMdSecondLevel 创建层级二
// @Tags MdSecondLevel
// @Summary 创建层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdSecondLevel true "创建层级二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdSecondLevel/createMdSecondLevel [post]
func (mdSecondLevelApi *MdSecondLevelApi) CreateMdSecondLevel(c *gin.Context) {
	var mdSecondLevel alpha.MdSecondLevel
	err := c.ShouldBindJSON(&mdSecondLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mdSecondLevel.CreatedBy = utils.GetUserID(c)

	if err := mdSecondLevelService.CreateMdSecondLevel(&mdSecondLevel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMdSecondLevel 删除层级二
// @Tags MdSecondLevel
// @Summary 删除层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdSecondLevel true "删除层级二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdSecondLevel/deleteMdSecondLevel [delete]
func (mdSecondLevelApi *MdSecondLevelApi) DeleteMdSecondLevel(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := mdSecondLevelService.DeleteMdSecondLevel(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMdSecondLevelByIds 批量删除层级二
// @Tags MdSecondLevel
// @Summary 批量删除层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mdSecondLevel/deleteMdSecondLevelByIds [delete]
func (mdSecondLevelApi *MdSecondLevelApi) DeleteMdSecondLevelByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := mdSecondLevelService.DeleteMdSecondLevelByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMdSecondLevel 更新层级二
// @Tags MdSecondLevel
// @Summary 更新层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdSecondLevel true "更新层级二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdSecondLevel/updateMdSecondLevel [put]
func (mdSecondLevelApi *MdSecondLevelApi) UpdateMdSecondLevel(c *gin.Context) {
	var mdSecondLevel alpha.MdSecondLevel
	err := c.ShouldBindJSON(&mdSecondLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mdSecondLevel.UpdatedBy = utils.GetUserID(c)

	if err := mdSecondLevelService.UpdateMdSecondLevel(mdSecondLevel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMdSecondLevel 用id查询层级二
// @Tags MdSecondLevel
// @Summary 用id查询层级二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.MdSecondLevel true "用id查询层级二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdSecondLevel/findMdSecondLevel [get]
func (mdSecondLevelApi *MdSecondLevelApi) FindMdSecondLevel(c *gin.Context) {
	ID := c.Query("ID")
	if remdSecondLevel, err := mdSecondLevelService.GetMdSecondLevel(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remdSecondLevel": remdSecondLevel}, c)
	}
}

// GetMdSecondLevelList 分页获取层级二列表
// @Tags MdSecondLevel
// @Summary 分页获取层级二列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.MdSecondLevelSearch true "分页获取层级二列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdSecondLevel/getMdSecondLevelList [get]
func (mdSecondLevelApi *MdSecondLevelApi) GetMdSecondLevelList(c *gin.Context) {
	var pageInfo alphaReq.MdSecondLevelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mdSecondLevelService.GetMdSecondLevelInfoList(pageInfo); err != nil {
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
