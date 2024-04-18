// 自动生成模板TecBaseProcess
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 工艺、设备及治具基本信息 结构体  TecBaseProcess
type TecBaseProcess struct {
	global.GVA_MODEL
	ParenId      *int   `json:"parenId" form:"parenId" gorm:"column:paren_id;comment:;"`                     //主ID
	ProcessType  string `json:"processType" form:"processType" gorm:"column:process_type;comment:;size:20;"` //工艺方式
	Unoil        string `json:"unoil" form:"unoil" gorm:"column:unoil;comment:;size:20;"`                    //除油
	ShotBlasting string `json:"shotBlasting" form:"shotBlasting" gorm:"column:shot_blasting;comment:;"`      //抛丸
	Phosphat     string `json:"phosphat" form:"phosphat" gorm:"column:phosphat;comment:;"`                   //磷化
	Electroplate string `json:"electroplate" form:"electroplate" gorm:"column:electroplate;comment:;"`       //电镀
	Remark       string `json:"remark" form:"remark" gorm:"column:remark;comment:;size:255;"`                //备注
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 工艺、设备及治具基本信息 TecBaseProcess自定义表名 tec_base_process
func (TecBaseProcess) TableName() string {
	return "tec_base_process"
}
