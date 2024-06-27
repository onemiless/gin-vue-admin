// 自动生成模板TecBaseCraftsmanship
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 入篮量和程序号 结构体  TecBaseCraftsmanship
type TecBaseCraftsmanship struct {
	global.GVA_MODEL
	ParentId           *int     `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;" binding:"required"`                               //主ID
	UTN                string   `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:100;" binding:"required"`                                      //唯一追踪号
	MB202              string   `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:100;" binding:"required"`                                //客户品名
	Process            string   `json:"process" form:"process" gorm:"column:process;comment:;size:20;" binding:"required"`                           //工艺
	Coating            string   `json:"coating" form:"coating" gorm:"column:coating;comment:;size:100;"`                                             //涂料/溶剂/其他
	ProductLine        *int     `json:"productLine" form:"productLine" gorm:"column:product_line;comment:;" binding:"required"`                      //产线
	BasketQuantity     *float64 `json:"basketQuantity" form:"basketQuantity" gorm:"column:basket_quantity;comment:;" binding:"required"`             //入篮量
	ProcedureNumber    string   `json:"procedureNumber" form:"procedureNumber" gorm:"column:procedure_number;comment:;size:100;" binding:"required"` //程序号
	Beat               *float64 `json:"beat" form:"beat" gorm:"column:beat;comment:;" binding:"required"`                                            //节拍
	ProductionCapacity *float64 `json:"productionCapacity" form:"productionCapacity" gorm:"column:production_capacity;comment:;"`                    //产能
	CreatedBy          uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 入篮量和程序号 TecBaseCraftsmanship自定义表名 tec_base_craftsmanship
func (TecBaseCraftsmanship) TableName() string {
	return "tec_base_craftsmanship"
}
