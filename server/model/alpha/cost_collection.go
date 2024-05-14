// 自动生成模板CostCollection
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 成本信息收集信息 结构体  CostCollection
type CostCollection struct {
	global.GVA_MODEL
	ParentId           *int       `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;"`                                          //父ID
	TypeOfPlan         string     `json:"typeOfPlan" form:"typeOfPlan" gorm:"column:type_of_plan;comment:;" binding:"required"`                //策划类型
	CountOfPlan        *int       `json:"countOfPlan" form:"countOfPlan" gorm:"column:count_of_plan;comment:;" binding:"required"`             //策划次数
	BasketUnit         string     `json:"basketUnit" form:"basketUnit" gorm:"column:basket_unit;comment:;size:20;" binding:"required"`         //入篮单位
	Prefix             string     `json:"prefix" form:"prefix" gorm:"column:prefix;comment:;size:255;"`                                        //前缀
	NumberOfPlan       string     `json:"numberOfPlan" form:"numberOfPlan" gorm:"column:number_of_plan;comment:;size:255;" binding:"required"` //策划单号
	DateReciveOfPlan   *time.Time `json:"dateReciveOfPlan" form:"dateReciveOfPlan" gorm:"column:date_recive_of_plan;comment:;"`                //策划接收日期
	SN                 string     `json:"SN" form:"SN" gorm:"column:SN;comment:;size:10;" binding:"required"`                                  //流水号
	DateIssuedOfPlan   *time.Time `json:"dateIssuedOfPlan" form:"dateIssuedOfPlan" gorm:"column:date_issued_of_plan;comment:;"`                //策划发出日期
	PlanBasis          string     `json:"planBasis" form:"planBasis" gorm:"column:plan_basis;comment:;" binding:"required"`                    //策划依据
	MainPlanLine       string     `json:"mainPlanLine" form:"mainPlanLine" gorm:"column:main_plan_line;comment:;size:20;" binding:"required"`  //主要策划线别
	Baskets            *float64   `json:"baskets" form:"baskets" gorm:"column:baskets;comment:;"`                                              //入篮量
	Beat               *float64   `json:"beat" form:"beat" gorm:"column:beat;comment:;"`                                                       //节拍
	ProductionCapacity *float64   `json:"productionCapacity" form:"productionCapacity" gorm:"column:production_capacity;comment:;"`            //小时产能
	Remark             string     `json:"remark" form:"remark" gorm:"column:remark;comment:;type:text;"`                                       //策划的特殊说明
	UTN                string     `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:100;" binding:"required"`                              //唯一追踪号
	MB202              string     `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:100;" binding:"required"`                        //客户品号
	CreatedBy          uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 成本信息收集信息 CostCollection自定义表名 cost_collection
func (CostCollection) TableName() string {
	return "cost_collection"
}
