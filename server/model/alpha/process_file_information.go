// 自动生成模板ProcessFileInformation
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 工艺文件信息 结构体  ProcessFileInformation
type ProcessFileInformation struct {
	global.GVA_MODEL
	ParentId           *int       `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;" binding:"required"`                        //父ID
	UTN                string     `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:100;" binding:"required"`                               //唯一追踪号
	MB202              string     `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:100;"`                                            //客户品号
	ProcessCardNumber  string     `json:"processCardNumber" form:"processCardNumber" gorm:"column:process_card_number;comment:;size:100;"`      //工艺卡编号
	ProcessCardSN      *int       `json:"processCardSN" form:"processCardSN" gorm:"column:process_card_SN;comment:;" binding:"required"`        //工艺卡流水号
	LatestDate         *time.Time `json:"latestDate" form:"latestDate" gorm:"column:latest_date;comment:;" binding:"required"`                  //工艺卡最新日期
	VersionNumber      string     `json:"versionNumber" form:"versionNumber" gorm:"column:version_number;comment:;size:20;" binding:"required"` //工艺卡版本号
	Baskets            *float64   `json:"baskets" form:"baskets" gorm:"column:baskets;comment:;" binding:"required"`                            //工艺卡入篮量
	Beat               *float64   `json:"beat" form:"beat" gorm:"column:beat;comment:;" binding:"required"`                                     //工艺卡节拍
	ProductionCapacity *float64   `json:"productionCapacity" form:"productionCapacity" gorm:"column:production_capacity;comment:;"`             //工艺卡产能
	PDFSN              string     `json:"PDFSN" form:"PDFSN" gorm:"column:PDFSN;comment:;size:100;" binding:"required"`                         //PDF编号
	PDFCreateDate      *time.Time `json:"PDFCreateDate" form:"PDFCreateDate" gorm:"column:PDF_create_date;comment:;" binding:"required"`        //PFD建立日期
	PDFLatestDate      *time.Time `json:"PDFLatestDate" form:"PDFLatestDate" gorm:"column:PDF_latest_date;comment:;"`                           //PDF最新日期
	PDFVN              string     `json:"PDFVN" form:"PDFVN" gorm:"column:PDFVN;comment:;size:100;" binding:"required"`                         //PDF版本号
	PFMEASN            string     `json:"PFMEASN" form:"PFMEASN" gorm:"column:PFMEASN;comment:;size:100;" binding:"required"`                   //PFMEA编号
	PFMEACreateDate    *time.Time `json:"PFMEACreateDate" form:"PFMEACreateDate" gorm:"column:PFMEA_create_date;comment:;" binding:"required"`  //PFMEA建立日期
	PFMEALatestDate    *time.Time `json:"PFMEALatestDate" form:"PFMEALatestDate" gorm:"column:PFMEA_latest_date;comment:;"`                     //PFMEA最新日期
	PFMEAVN            string     `json:"PFMEAVN" form:"PFMEAVN" gorm:"column:PFMEAVN;comment:;size:100;" binding:"required"`                   //PFMEA版本号
	CPSN               string     `json:"CPSN" form:"CPSN" gorm:"column:CPSN;comment:;size:100;" binding:"required"`                            //CP编号
	CPCreateDate       *time.Time `json:"CPCreateDate" form:"CPCreateDate" gorm:"column:CP_create_date;comment:;" binding:"required"`           //CP建立日期
	CPLatestDate       *time.Time `json:"CPLatestDate" form:"CPLatestDate" gorm:"column:CP_latest_date;comment:;"`                              //CP最新日期
	CPSV               string     `json:"CPSV" form:"CPSV" gorm:"column:CPSV;comment:;size:100;" binding:"required"`                            //CP版本号
	CreatedBy          uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 工艺文件信息 ProcessFileInformation自定义表名 process_file_information
func (ProcessFileInformation) TableName() string {
	return "process_file_information"
}
