package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CMSXCSearch struct {
	request.PageInfo
	Query string `json:"query" form:"query"`
}
