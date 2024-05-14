package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MdThirdLevelSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name          string `json:"name" form:"name" `
	SecondLevelId uint   `json:"secondLevelId" form:"secondLevelId" `
	Query         string `json:"query" form:"query"`
	request.PageInfo
}
