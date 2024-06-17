// 自动生成模板MassProductionTransfer
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 量产转移 结构体  MassProductionTransfer
type MassProductionTransfer struct {
	global.GVA_MODEL
	ParentId       *int       `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;" binding:"required"`                            //父ID
	UTN            string     `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:100;"`                                                      //唯一追踪号
	MB202          string     `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:100;"`                                                //客户品号
	SN             string     `json:"SN" form:"SN" gorm:"column:SN;comment:;size:100;"`                                                         //量产转移单号
	AuditSN        string     `json:"auditSN" form:"auditSN" gorm:"column:audit_sn;comment:;" binding:"required"`                               //量产转移审批号
	TransferDate   *time.Time `json:"transferDate" form:"transferDate" gorm:"column:transfer_date;comment:;" binding:"required"`                //量产转移日期
	AuditStatus    string     `json:"auditStatus" form:"auditStatus" gorm:"column:audit_status;comment:;size:20;" binding:"required"`           //量产转移审批状态
	TransferStatus string     `json:"transferStatus" form:"transferStatus" gorm:"column:transfer_status;comment:;size:100;" binding:"required"` //量产转移状态
	SucessDate     *time.Time `json:"sucessDate" form:"sucessDate" gorm:"column:sucess_date;comment:;"`                                         //量产成功转移日期
	CreatedBy      uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 量产转移 MassProductionTransfer自定义表名 mass_production_transfer
func (MassProductionTransfer) TableName() string {
	return "mass_production_transfer"
}
