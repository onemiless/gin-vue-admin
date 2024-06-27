package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProductLineSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	ProductNumber string `json:"productNumber" form:"productNumber" `
	ProductName   string `json:"productName" form:"productName" `
	request.PageInfo
}
