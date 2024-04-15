// 自动生成模板MdUnitMeasure
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 品号单位设置 结构体  MdUnitMeasure
type MdUnitMeasure struct {
	global.GVA_MODEL
	Measure_code string `json:"measureCode" form:"measureCode" gorm:"column:measure_code;comment:单位编码;size:64;"binding:"required"`  //单位
	MeasureName  string `json:"measureName" form:"measureName" gorm:"column:measure_name;comment:单位名称;size:255;"binding:"required"` //单位名称
	EnableFlag   int    `json:"enableFlag" form:"enableFlag" gorm:"column:enable_flag;comment:是否启用;"binding:"required"`             //是否启用
	Remark       string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:500;"`                                     //备注
	Attr1        string `json:"attr1" form:"attr1" gorm:"column:attr1;comment:;size:64;"`                                           //属性1
	Attr2        string `json:"attr2" form:"attr2" gorm:"column:attr2;comment:;size:255;"`                                          //属性2
	Attr3        *int   `json:"attr3" form:"attr3" gorm:"column:attr3;comment:;"`                                                   //属性3
	Attr4        *int   `json:"attr4" form:"attr4" gorm:"column:attr4;comment:;"`                                                   //属性4
}

// TableName 品号单位设置 MdUnitMeasure自定义表名 md_unit_measure
func (MdUnitMeasure) TableName() string {
	return "md_unit_measure"
}
