package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MassProductionTransferSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	UTN               string     `json:"UTN" form:"UTN" `
	MB202             string     `json:"MB202" form:"MB202" `
	SN                string     `json:"SN" form:"SN" `
	AuditSN           string     `json:"auditSN" form:"auditSN" `
	StartTransferDate *time.Time `json:"startTransferDate" form:"startTransferDate"`
	EndTransferDate   *time.Time `json:"endTransferDate" form:"endTransferDate"`
	AuditStatus       string     `json:"auditStatus" form:"auditStatus" `
	StartSucessDate   *time.Time `json:"startSucessDate" form:"startSucessDate"`
	EndSucessDate     *time.Time `json:"endSucessDate" form:"endSucessDate"`
	request.PageInfo
}
