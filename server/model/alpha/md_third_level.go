// 自动生成模板MdThirdLevel
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 层级三 结构体  MdThirdLevel
type MdThirdLevel struct {
	global.GVA_MODEL
	SecondLevelId uint   `json:"secondLevelId" form:"secondLevelId" gorm:"column:secondLevel_id;comment:" binding:"required"` //层级二
	Name          string `json:"name" form:"name" gorm:"column:name;comment:;size:255;" binding:"required"`                   //名称
	IsEnable      string `json:"isEnable" form:"isEnable" gorm:"column:is_enable;comment:;" binding:"required"`               //是否启用
	Remark        string `json:"remark" form:"remark" gorm:"column:remark;comment:;"`                                         //备注
	CreatedBy     uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 层级三 MdThirdLevel自定义表名 md_thirdLevel
func (MdThirdLevel) TableName() string {
	return "md_thirdLevel"
}
