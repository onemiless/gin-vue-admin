package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PackageRequirementSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	UTN string `json:"UTN" form:"UTN" `
	SN  string `json:"SN" form:"SN" `
	request.PageInfo
}
