// 自动生成模板ProofingInformation
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 打样信息 结构体  ProofingInformation
type ProofingInformation struct {
	global.GVA_MODEL
	ParentId                       *int       `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;" binding:"required"`                                                           //父ID
	UTN                            string     `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:100;" binding:"required"`                                                                  //唯一追踪号
	MB202                          string     `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:100;" binding:"required"`                                                            //客户品号
	Prefix                         string     `json:"prefix" form:"prefix" gorm:"column:prefix;comment:;size:100;" binding:"required"`                                                         //前缀
	ProofingCount                  *int       `json:"proofingCount" form:"proofingCount" gorm:"column:proofing_count;comment:;"`                                                               //打样次数
	NumberOfProofing               string     `json:"numberOfProofing" form:"numberOfProofing" gorm:"column:number_of_proofing;comment:;size:100;" binding:"required"`                         //打样单号
	ReciveDate                     *time.Time `json:"reciveDate" form:"reciveDate" gorm:"column:recive_date;comment:;"`                                                                        //接收日期
	SN                             string     `json:"SN" form:"SN" gorm:"column:SN;comment:;size:10;"`                                                                                         //流水号
	EstimatedDateOfCompletion      *time.Time `json:"estimatedDateOfCompletion" form:"estimatedDateOfCompletion" gorm:"column:estimated_date_of_completion;comment:;"`                         //预计完成日期
	ProofingType                   string     `json:"proofingType" form:"proofingType" gorm:"column:proofing_type;comment:;"`                                                                  //打样类型
	ProofingMethod                 string     `json:"proofingMethod" form:"proofingMethod" gorm:"column:proofing_method;comment:;"`                                                            //打样方式
	ProofingCondition              string     `json:"proofingCondition" form:"proofingCondition" gorm:"column:proofing_condition;comment:;"`                                                   //打样状态
	SpecialInstructionsForProofing string     `json:"specialInstructionsForProofing" form:"specialInstructionsForProofing" gorm:"column:special_instructions_for_proofing;comment:;size:255;"` //打样特殊说明
	ProofDrawing                   string     `json:"proofDrawing" form:"proofDrawing" gorm:"column:proof_drawing;comment:;" binding:"required"`                                               //打样图纸
	IsReservedSample               string     `json:"isReservedSample" form:"isReservedSample" gorm:"column:is_reserved_samplecomment:;"`                                                      //是否留样
	GoodsReceived                  *float64   `json:"goodsReceived" form:"goodsReceived" gorm:"column:goods_received;comment:;"`                                                               //入货量
	IncomingUnit                   string     `json:"incomingUnit" form:"incomingUnit" gorm:"column:incoming_unit;comment:;size:20;"`                                                          //入货单位
	QualityJudgment                string     `json:"qualityJudgment" form:"qualityJudgment" gorm:"column:quality_judgment;comment:;"`                                                         //佳马质量判定
	SaltSprayDetermination         string     `json:"saltSprayDetermination" form:"saltSprayDetermination" gorm:"column:salt_spray_determination;comment:;size:100;"`                          //佳马盐雾判定
	ShippingStatus                 string     `json:"shippingStatus" form:"shippingStatus" gorm:"column:shipping_status;comment:;"`                                                            //出货状态
	DateIssued                     *time.Time `json:"dateIssued" form:"dateIssued" gorm:"column:date_issued;comment:;"`                                                                        //发出日期
	CustomerQualityJudgment        string     `json:"customerQualityJudgment" form:"customerQualityJudgment" gorm:"column:customer_quality_judgment;comment:;"`                                //客户质量判定
	CustomerQualityStatement       string     `json:"customerQualityStatement" form:"customerQualityStatement" gorm:"column:customer_quality_statement;comment:;size:255;"`                    //客户质量说明
	ProofingStatus                 string     `json:"proofingCondition" form:"proofingCondition" gorm:"column:proofing_condition;comment:;size:255;"`                                          //打样状态
	DeliveryJudgment               string     `json:"deliveryJudgment" form:"deliveryJudgment" gorm:"column:delivery_judgment;comment:;size:100;"`                                             //交期判断
	CreatedBy                      uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy                      uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy                      uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 打样信息 ProofingInformation自定义表名 proofing_information
func (ProofingInformation) TableName() string {
	return "proofing_information"
}
