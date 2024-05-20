package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EntryNumberApi struct {
}

var entryNumberService = service.ServiceGroupApp.AlphaServiceGroup.IEntryNumberService

// 获取流水号
func (entryNumberApi *EntryNumberApi) FindEntryNumber(c *gin.Context) {
	//tableName := c.Query("tableName")
	//fmt.Println(value)
	var entryNumber alpha.EntryNumber
	err := c.ShouldBindJSON(&entryNumber)
	if err != nil {
		return
	}
	if number, err := entryNumberService.GetEntryNumber(entryNumber); err != nil {
		global.GVA_LOG.Error("获取流水号失败!", zap.Error(err))
		response.FailWithMessage("获取流水号失败", c)
	} else {
		response.OkWithData(gin.H{"number": number}, c)
	}
}
