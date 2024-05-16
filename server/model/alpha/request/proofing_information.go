package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProofingInformationSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	UTN            string     `json:"UTN" form:"UTN" `
	MB202          string     `json:"MB202" form:"MB202" `
	request.PageInfo
}
