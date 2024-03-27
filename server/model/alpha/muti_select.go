// 自动生成模板MutiSelect
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 多级选择信息 结构体  MutiSelect
type MutiSelect struct {
	global.GVA_MODEL
	Lable       string `json:"typeName" form:"typeName" gorm:"column:lable;comment:显示内容;size:255;"`  //显示内容
	ParentValue *int   `json:"typeParent" form:"typeParent" gorm:"column:parent_value;comment:上层值;"` //上层值
	Tier        *int   `json:"tier" form:"tier" gorm:"column:tier;comment:层级;"`                      //层级
	Value       *int   `json:"value" form:"value" gorm:"column:value;comment:值;"`                    //值
	Status      *int   `json:"status" form:"status" gorm:"column:status;comment:是否启用;"`              //是否启用
}

// TableName 多级选择信息 MutiSelect自定义表名 muti_select
func (MutiSelect) TableName() string {
	return "muti_select"
}
