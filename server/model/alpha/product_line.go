// 自动生成模板ProductLine
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 产线信息 结构体  ProductLine
type ProductLine struct {
	global.GVA_MODEL
	ProductNumber string `json:"productNumber" form:"productNumber" gorm:"column:product_number;comment:;size:100;" binding:"required"` //产线编号
	ProductName   string `json:"productName" form:"productName" gorm:"column:product_name;comment:;size:100;" binding:"required"`       //产线名称
	Subdivision   string `json:"subdivision" form:"subdivision" gorm:"column:subdivision;comment:;size:20;" binding:"required"`         //所属分部
	Status        string `json:"status" form:"status" gorm:"column:status;comment:;size:20;" binding:"required"`                        //是否启用
	Remark        string `json:"remark" form:"remark" gorm:"column:remark;comment:;size:500;"`                                          //描述
	CreatedBy     uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 产线信息 ProductLine自定义表名 product_line
func (ProductLine) TableName() string {
	return "product_line"
}
