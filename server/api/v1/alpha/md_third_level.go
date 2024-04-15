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

type MdThirdLevelApi struct {
}

var mdThirdLevelService = service.ServiceGroupApp.AlphaServiceGroup.MdThirdLevelService

// CreateMdThirdLevel 创建层级三
// @Tags MdThirdLevel
// @Summary 创建层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdThirdLevel true "创建层级三"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mdThirdLevel/createMdThirdLevel [post]
func (mdThirdLevelApi *MdThirdLevelApi) CreateMdThirdLevel(c *gin.Context) {
	var mdThirdLevel alpha.MdThirdLevel
	err := c.ShouldBindJSON(&mdThirdLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mdThirdLevel.CreatedBy = utils.GetUserID(c)

	if err := mdThirdLevelService.CreateMdThirdLevel(&mdThirdLevel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMdThirdLevel 删除层级三
// @Tags MdThirdLevel
// @Summary 删除层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdThirdLevel true "删除层级三"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mdThirdLevel/deleteMdThirdLevel [delete]
func (mdThirdLevelApi *MdThirdLevelApi) DeleteMdThirdLevel(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := mdThirdLevelService.DeleteMdThirdLevel(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMdThirdLevelByIds 批量删除层级三
// @Tags MdThirdLevel
// @Summary 批量删除层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mdThirdLevel/deleteMdThirdLevelByIds [delete]
func (mdThirdLevelApi *MdThirdLevelApi) DeleteMdThirdLevelByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := mdThirdLevelService.DeleteMdThirdLevelByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMdThirdLevel 更新层级三
// @Tags MdThirdLevel
// @Summary 更新层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.MdThirdLevel true "更新层级三"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mdThirdLevel/updateMdThirdLevel [put]
func (mdThirdLevelApi *MdThirdLevelApi) UpdateMdThirdLevel(c *gin.Context) {
	var mdThirdLevel alpha.MdThirdLevel
	err := c.ShouldBindJSON(&mdThirdLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mdThirdLevel.UpdatedBy = utils.GetUserID(c)

	if err := mdThirdLevelService.UpdateMdThirdLevel(mdThirdLevel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMdThirdLevel 用id查询层级三
// @Tags MdThirdLevel
// @Summary 用id查询层级三
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.MdThirdLevel true "用id查询层级三"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mdThirdLevel/findMdThirdLevel [get]
func (mdThirdLevelApi *MdThirdLevelApi) FindMdThirdLevel(c *gin.Context) {
	ID := c.Query("ID")
	if remdThirdLevel, err := mdThirdLevelService.GetMdThirdLevel(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remdThirdLevel": remdThirdLevel}, c)
	}
}

// GetMdThirdLevelList 分页获取层级三列表
// @Tags MdThirdLevel
// @Summary 分页获取层级三列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.MdThirdLevelSearch true "分页获取层级三列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mdThirdLevel/getMdThirdLevelList [get]
func (mdThirdLevelApi *MdThirdLevelApi) GetMdThirdLevelList(c *gin.Context) {
	var pageInfo alphaReq.MdThirdLevelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mdThirdLevelService.GetMdThirdLevelInfoList(pageInfo); err != nil {
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
