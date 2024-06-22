// 自动生成模板PackageRequirement
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 包装要求信息 结构体  PackageRequirement
type PackageRequirement struct {
	global.GVA_MODEL
	ParentId      *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;" binding:"required"`                         //父ID
	UTN           string `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:100;" binding:"required"`                                //唯一追踪号
	MB202         string `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:100;" binding:"required"`                          //客户品号
	SN            string `json:"SN" form:"SN" gorm:"column:SN;comment:;size:100;" binding:"required"`                                   //包装方案编号
	PackageMethod string `json:"packageMethod" form:"packageMethod" gorm:"column:package_method;comment:;size:999;" binding:"required"` //包装方式
	CreatedBy     uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 包装要求信息 PackageRequirement自定义表名 package_requirement
func (PackageRequirement) TableName() string {
	return "package_requirement"
}
