// 自动生成模板ItemType
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 零件类型 结构体  ItemType
type ItemType struct {
	global.GVA_MODEL
	TypeName   string `json:"typeName" form:"typeName" gorm:"column:type_name;comment:零件类型;size:50;"binding:"required"` //零件类型
	TypeParent *int   `json:"typeParent" form:"typeParent" gorm:"column:type_parent;comment:;"`                         //上级类型
	IsBase     *int   `json:"isBase" form:"isBase" gorm:"column:is_base;comment:基础类型;"binding:"required"`               //基础类型
}

// TableName 零件类型 ItemType自定义表名 item_type
func (ItemType) TableName() string {
	return "item_type"
}
