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

type ProofingInformationApi struct {
}

var proofingInformationService = service.ServiceGroupApp.AlphaServiceGroup.ProofingInformationService

// CreateProofingInformation 创建打样信息
// @Tags ProofingInformation
// @Summary 创建打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ProofingInformation true "创建打样信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /proofingInformation/createProofingInformation [post]
func (proofingInformationApi *ProofingInformationApi) CreateProofingInformation(c *gin.Context) {
	var proofingInformation alpha.ProofingInformation
	err := c.ShouldBindJSON(&proofingInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	proofingInformation.CreatedBy = utils.GetUserID(c)

	if err := proofingInformationService.CreateProofingInformation(&proofingInformation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProofingInformation 删除打样信息
// @Tags ProofingInformation
// @Summary 删除打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ProofingInformation true "删除打样信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /proofingInformation/deleteProofingInformation [delete]
func (proofingInformationApi *ProofingInformationApi) DeleteProofingInformation(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := proofingInformationService.DeleteProofingInformation(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProofingInformationByIds 批量删除打样信息
// @Tags ProofingInformation
// @Summary 批量删除打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /proofingInformation/deleteProofingInformationByIds [delete]
func (proofingInformationApi *ProofingInformationApi) DeleteProofingInformationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := proofingInformationService.DeleteProofingInformationByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProofingInformation 更新打样信息
// @Tags ProofingInformation
// @Summary 更新打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.ProofingInformation true "更新打样信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /proofingInformation/updateProofingInformation [put]
func (proofingInformationApi *ProofingInformationApi) UpdateProofingInformation(c *gin.Context) {
	var proofingInformation alpha.ProofingInformation
	err := c.ShouldBindJSON(&proofingInformation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	proofingInformation.UpdatedBy = utils.GetUserID(c)

	if err := proofingInformationService.UpdateProofingInformation(proofingInformation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProofingInformation 用id查询打样信息
// @Tags ProofingInformation
// @Summary 用id查询打样信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.ProofingInformation true "用id查询打样信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /proofingInformation/findProofingInformation [get]
func (proofingInformationApi *ProofingInformationApi) FindProofingInformation(c *gin.Context) {
	ID := c.Query("ID")
	if reproofingInformation, err := proofingInformationService.GetProofingInformation(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproofingInformation": reproofingInformation}, c)
	}
}

// GetProofingInformationList 分页获取打样信息列表
// @Tags ProofingInformation
// @Summary 分页获取打样信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.ProofingInformationSearch true "分页获取打样信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /proofingInformation/getProofingInformationList [get]
func (proofingInformationApi *ProofingInformationApi) GetProofingInformationList(c *gin.Context) {
	var pageInfo alphaReq.ProofingInformationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := proofingInformationService.GetProofingInformationInfoList(pageInfo); err != nil {
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
