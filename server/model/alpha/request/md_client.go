package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MdClientSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	ClientCode     string     `json:"clientCode" form:"clientCode"`  //客户编码
	ClientName     string     `json:"clientName" form:"clientName"`  //客户名称
	ClientEn       string     `json:"clientEn" form:"clientEn"`      //英文名称
	ClientDes      string     `json:"clientDes" form:"clientDes" `   //描述
	ClientType     string     `json:"clientType" form:"clientType" ` //拼音首字母
	Search         string     `json:"search" form:"search"`          //搜索条件
	EnableFlag     *int       `json:"enableFlag" form:"enableFlag"`  //启用状态

	request.PageInfo
}
