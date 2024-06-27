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

type TecBaseCraftsmanshipApi struct{}

var tecBaseCraftsmanshipService = service.ServiceGroupApp.AlphaServiceGroup.TecBaseCraftsmanshipService

// CreateTecBaseCraftsmanship 创建入篮量和程序号
// @Tags TecBaseCraftsmanship
// @Summary 创建入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseCraftsmanship true "创建入篮量和程序号"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /tecBaseCraftsmanship/createTecBaseCraftsmanship [post]
func (tecBaseCraftsmanshipApi *TecBaseCraftsmanshipApi) CreateTecBaseCraftsmanship(c *gin.Context) {
	var tecBaseCraftsmanship []alpha.TecBaseCraftsmanship
	err := c.ShouldBindJSON(&tecBaseCraftsmanship)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	for _, craftsmanship := range tecBaseCraftsmanship {
		//if i == 0 {
		craftsmanship.CreatedBy = utils.GetUserID(c)
		//}

	}

	//tecBaseCraftsmanship.CreatedBy = utils.GetUserID(c)

	if err := tecBaseCraftsmanshipService.CreateTecBaseCraftsmanship(&tecBaseCraftsmanship); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTecBaseCraftsmanship 删除入篮量和程序号
// @Tags TecBaseCraftsmanship
// @Summary 删除入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseCraftsmanship true "删除入篮量和程序号"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /tecBaseCraftsmanship/deleteTecBaseCraftsmanship [delete]
func (tecBaseCraftsmanshipApi *TecBaseCraftsmanshipApi) DeleteTecBaseCraftsmanship(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := tecBaseCraftsmanshipService.DeleteTecBaseCraftsmanship(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTecBaseCraftsmanshipByIds 批量删除入篮量和程序号
// @Tags TecBaseCraftsmanship
// @Summary 批量删除入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /tecBaseCraftsmanship/deleteTecBaseCraftsmanshipByIds [delete]
func (tecBaseCraftsmanshipApi *TecBaseCraftsmanshipApi) DeleteTecBaseCraftsmanshipByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := tecBaseCraftsmanshipService.DeleteTecBaseCraftsmanshipByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTecBaseCraftsmanship 更新入篮量和程序号
// @Tags TecBaseCraftsmanship
// @Summary 更新入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body alpha.TecBaseCraftsmanship true "更新入篮量和程序号"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /tecBaseCraftsmanship/updateTecBaseCraftsmanship [put]
func (tecBaseCraftsmanshipApi *TecBaseCraftsmanshipApi) UpdateTecBaseCraftsmanship(c *gin.Context) {
	var tecBaseCraftsmanship []alpha.TecBaseCraftsmanship
	err := c.ShouldBindJSON(&tecBaseCraftsmanship)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	for _, craftsmanship := range tecBaseCraftsmanship {
		//if i == 0 {
		craftsmanship.UpdatedBy = utils.GetUserID(c)
		//}
	}
	//tecBaseCraftsmanship.UpdatedBy = utils.GetUserID(c)

	if err := tecBaseCraftsmanshipService.UpdateTecBaseCraftsmanship(tecBaseCraftsmanship); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTecBaseCraftsmanship 用id查询入篮量和程序号
// @Tags TecBaseCraftsmanship
// @Summary 用id查询入篮量和程序号
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.TecBaseCraftsmanship true "用id查询入篮量和程序号"
// @Success 200 {object} response.Response{data=object{retecBaseCraftsmanship=alpha.TecBaseCraftsmanship},msg=string} "查询成功"
// @Router /tecBaseCraftsmanship/findTecBaseCraftsmanship [get]
func (tecBaseCraftsmanshipApi *TecBaseCraftsmanshipApi) FindTecBaseCraftsmanship(c *gin.Context) {
	parentId := c.Query("parentId")
	if retecBaseCraftsmanship, err := tecBaseCraftsmanshipService.GetTecBaseCraftsmanship(parentId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(retecBaseCraftsmanship, c)
	}
}

// GetTecBaseCraftsmanshipList 分页获取入篮量和程序号列表
// @Tags TecBaseCraftsmanship
// @Summary 分页获取入篮量和程序号列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.TecBaseCraftsmanshipSearch true "分页获取入篮量和程序号列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /tecBaseCraftsmanship/getTecBaseCraftsmanshipList [get]
func (tecBaseCraftsmanshipApi *TecBaseCraftsmanshipApi) GetTecBaseCraftsmanshipList(c *gin.Context) {
	var pageInfo alphaReq.TecBaseCraftsmanshipSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tecBaseCraftsmanshipService.GetTecBaseCraftsmanshipInfoList(pageInfo); err != nil {
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

// GetTecBaseCraftsmanshipDataSource 获取TecBaseCraftsmanship的数据源
// @Tags TecBaseCraftsmanship
// @Summary 获取TecBaseCraftsmanship的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /tecBaseCraftsmanship/getTecBaseCraftsmanshipDataSource [get]
func (tecBaseCraftsmanshipApi *TecBaseCraftsmanshipApi) GetTecBaseCraftsmanshipDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := tecBaseCraftsmanshipService.GetTecBaseCraftsmanshipDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}

// GetTecBaseCraftsmanshipPublic 不需要鉴权的入篮量和程序号接口
// @Tags TecBaseCraftsmanship
// @Summary 不需要鉴权的入篮量和程序号接口
// @accept application/json
// @Produce application/json
// @Param data query alphaReq.TecBaseCraftsmanshipSearch true "分页获取入篮量和程序号列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /tecBaseCraftsmanship/getTecBaseCraftsmanshipPublic [get]
func (tecBaseCraftsmanshipApi *TecBaseCraftsmanshipApi) GetTecBaseCraftsmanshipPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的入篮量和程序号接口信息",
	}, "获取成功", c)
}
