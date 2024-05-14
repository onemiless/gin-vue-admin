package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TecBaseInfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	ID             uint       `json:"ID" form:"ID" ` // 主键ID
	ME002          string     `json:"ME002" form:"ME002" `
	Type           *int       `json:"type" form:"type" `
	Typeext        *int       `json:"typeext" form:"typeext" `
	MB201          string     `json:"MB201" form:"MB201" `
	MB002          string     `json:"MB002" form:"MB002" `
	MB202          string     `json:"MB202" form:"MB202" `
	MB003          string     `json:"MB003" form:"MB003" `
	OE             string     `json:"OE" form:"OE" `
	UTN            string     `json:"UTN" form:"UTN" `
	request.PageInfo
}
