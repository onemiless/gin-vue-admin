// 自动生成模板EngineeringChange
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 工程变更信息 结构体  EngineeringChange
type EngineeringChange struct {
	global.GVA_MODEL
	ParentId     *int       `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;" binding:"required"`                       //父ID
	UTN          string     `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:100;" binding:"required"`                              //唯一追踪号
	MB202        string     `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:100;"`                                           //客户品名
	SN           string     `json:"SN" form:"SN" gorm:"column:SN;comment:;size:100;" binding:"required"`                                 //变更单号
	ChangeDate   *time.Time `json:"changeDate" form:"changeDate" gorm:"column:change_date;comment:;" binding:"required"`                 //变更日期
	ChangeDetail string     `json:"changeDetail" form:"changeDetail" gorm:"column:change_detail;comment:;type:text;" binding:"required"` //变更内容
	CreatedBy    uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 工程变更信息 EngineeringChange自定义表名 engineering_change
func (EngineeringChange) TableName() string {
	return "engineering_change"
}
