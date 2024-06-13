package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TecSNRuleApi struct {
}

var tecSNRuleService = service.ServiceGroupApp.AlphaServiceGroup.TecSNRuleService

// FindSNRuler 获取编码规则编号
// @Tags MdFirstLevel
// @Summary 根据规则名称获取对应的编号规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query alpha.FindSNRuler true "用于获取技术部基础资料的编码规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FindSNRuler/FindSNRuler [post]
func (tecSNRuleApi *TecSNRuleApi) FindSNRuler(c *gin.Context) {
	//ID := c.Query("ID")
	var pageInfo alphaReq.TecSNRuleSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if result, err := tecSNRuleService.GetTecSNRule(pageInfo); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(result, c)
	}
}

//
//// GetSNRulerLevelList 分页获取层级一列表
//// @Tags MdFirstLevel
//// @Summary 分页获取层级一列表
//// @Security ApiKeyAuth
//// @accept application/json
//// @Produce application/json
//// @Param data query alphaReq.MdFirstLevelSearch true "分页获取层级一列表"
//// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
//// @Router /mdFirstLevel/getMdFirstLevelList [get]
//func (tecSNRuleApi *TecSNRuleApi) GetSNRulerLevelList(c *gin.Context) {
//	var pageInfo alphaReq.MdFirstLevelSearch
//	err := c.ShouldBindQuery(&pageInfo)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	if list, total, err := mdFirstLevelService.GetMdFirstLevelInfoList(pageInfo); err != nil {
//		global.GVA_LOG.Error("获取失败!", zap.Error(err))
//		response.FailWithMessage("获取失败", c)
//	} else {
//		response.OkWithDetailed(response.PageResult{
//			List:     list,
//			Total:    total,
//			Page:     pageInfo.Page,
//			PageSize: pageInfo.PageSize,
//		}, "获取成功", c)
//	}
//}
