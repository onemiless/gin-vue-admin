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

type PackageRequirementApi struct {
}

var packageRequirementService = service.ServiceGroupApp.AlphaServiceGroup.PackageRequirementService

// CreatePackageRequirement 创建包装要求信息
// @Tags PackageRequirement
// @Summary 创建包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.PackageRequirement true "创建包装要求信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /packageRequirement/createPackageRequirement [post]
func (packageRequirementApi *PackageRequirementApi) CreatePackageRequirement(c *gin.Context) {
	var packageRequirement alpha.PackageRequirement
	err := c.ShouldBindJSON(&packageRequirement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	packageRequirement.CreatedBy = utils.GetUserID(c)

	if err := packageRequirementService.CreatePackageRequirement(&packageRequirement); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePackageRequirement 删除包装要求信息
// @Tags PackageRequirement
// @Summary 删除包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.PackageRequirement true "删除包装要求信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /packageRequirement/deletePackageRequirement [delete]
func (packageRequirementApi *PackageRequirementApi) DeletePackageRequirement(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := packageRequirementService.DeletePackageRequirement(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePackageRequirementByIds 批量删除包装要求信息
// @Tags PackageRequirement
// @Summary 批量删除包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /packageRequirement/deletePackageRequirementByIds [delete]
func (packageRequirementApi *PackageRequirementApi) DeletePackageRequirementByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := packageRequirementService.DeletePackageRequirementByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePackageRequirement 更新包装要求信息
// @Tags PackageRequirement
// @Summary 更新包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.PackageRequirement true "更新包装要求信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /packageRequirement/updatePackageRequirement [put]
func (packageRequirementApi *PackageRequirementApi) UpdatePackageRequirement(c *gin.Context) {
	var packageRequirement alpha.PackageRequirement
	err := c.ShouldBindJSON(&packageRequirement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	packageRequirement.UpdatedBy = utils.GetUserID(c)

	if err := packageRequirementService.UpdatePackageRequirement(packageRequirement); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPackageRequirement 用id查询包装要求信息
// @Tags PackageRequirement
// @Summary 用id查询包装要求信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.PackageRequirement true "用id查询包装要求信息"
// @Success 200 {object} response.Response{data=object{repackageRequirement=alpha.PackageRequirement},msg=string} "查询成功"
// @Router /packageRequirement/findPackageRequirement [get]
func (packageRequirementApi *PackageRequirementApi) FindPackageRequirement(c *gin.Context) {
	ID := c.Query("ID")
	if repackageRequirement, err := packageRequirementService.GetPackageRequirement(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(repackageRequirement, c)
	}
}

// GetPackageRequirementList 分页获取包装要求信息列表
// @Tags PackageRequirement
// @Summary 分页获取包装要求信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.PackageRequirementSearch true "分页获取包装要求信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /packageRequirement/getPackageRequirementList [get]
func (packageRequirementApi *PackageRequirementApi) GetPackageRequirementList(c *gin.Context) {
	var pageInfo alphaReq.PackageRequirementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := packageRequirementService.GetPackageRequirementInfoList(pageInfo); err != nil {
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

// GetPackageRequirementPublic 不需要鉴权的包装要求信息接口
// @Tags PackageRequirement
// @Summary 不需要鉴权的包装要求信息接口
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.PackageRequirementSearch true "分页获取包装要求信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /packageRequirement/getPackageRequirementPublic [get]
func (packageRequirementApi *PackageRequirementApi) GetPackageRequirementPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的包装要求信息接口信息",
	}, "获取成功", c)
}
