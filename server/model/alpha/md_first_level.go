// 自动生成模板MdFirstLevel
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 层级一 结构体  MdFirstLevel
type MdFirstLevel struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:;size:255;" binding:"required"`     //名称
	IsEnable  string `json:"isEnable" form:"isEnable" gorm:"column:is_enable;comment:;" binding:"required"` //是否启用
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`                           //备注
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 层级一 MdFirstLevel自定义表名 md_firstLevel
func (MdFirstLevel) TableName() string {
	return "md_firstLevel"
}
