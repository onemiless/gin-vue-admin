package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MdSecondLevelSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name         string `json:"name" form:"name" `
	Remark       string `json:"remark" form:"remark" `
	FirstLevelId uint   `json:"firstLevelId" form:"firstLevelId" `
	request.PageInfo
}