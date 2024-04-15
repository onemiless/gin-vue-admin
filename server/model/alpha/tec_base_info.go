// 自动生成模板TecBaseInfo
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 技术部基础信息 结构体  TecBaseInfo
type TecBaseInfo struct {
	global.GVA_MODEL
	ME002               string   `json:"ME002" form:"ME002" gorm:"column:ME002;comment:;size:255;" binding:"required"`                    //分部
	Type                *int     `json:"type" form:"type" gorm:"column:type;comment:;" binding:"required"`                                //零件类型1
	Typeext             *int     `json:"typeext" form:"typeext" gorm:"column:typeext;comment:;"`                                          //零件类型2
	MB201               string   `json:"MB201" form:"MB201" gorm:"column:MB201;comment:;size:255;" binding:"required"`                    //客户名称
	MB002               string   `json:"MB002" form:"MB002" gorm:"column:MB002;comment:;size:255;" binding:"required"`                    //零件名称
	MB202               string   `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:255;" binding:"required"`                    //客户品号
	MB003               string   `json:"MB003" form:"MB003" gorm:"column:MB003;comment:;size:255;" binding:"required"`                    //零件规格
	FigureNumber        string   `json:"figureNumber" form:"figureNumber" gorm:"column:figure_number;comment:;size:255;"`                 //零件图号
	MB004               string   `json:"MB004" form:"MB004" gorm:"column:MB004;comment:;" binding:"required"`                             //计价单位
	UTN                 string   `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:255;" binding:"required"`                          //唯一追踪号
	Textrue             string   `json:"textrue" form:"textrue" gorm:"column:textrue;comment:;size:255;"`                                 //材质
	Level               string   `json:"level" form:"level" gorm:"column:level;comment:;size:255;"`                                       //硬度/机械等级
	Length              *float64 `json:"length" form:"length" gorm:"column:length;comment:;"`                                             //长
	Width               *float64 `json:"width" form:"width" gorm:"column:width;comment:;"`                                                //宽
	Height              *float64 `json:"height" form:"height" gorm:"column:height;comment:;"`                                             //高
	Diameter            *float64 `json:"diameter" form:"diameter" gorm:"column:diameter;comment:;"`                                       //最大直径
	Thick               *float64 `json:"thick" form:"thick" gorm:"column:thick;comment:;"`                                                //料厚
	WireDiameter        *float64 `json:"wireDiameter" form:"wireDiameter" gorm:"column:wire_diameter;comment:;"`                          //线径
	PieceWeight         *float64 `json:"pieceWeight" form:"pieceWeight" gorm:"column:piece_weight;comment:;"`                             //单重
	BET                 *float64 `json:"BET" form:"BET" gorm:"column:BET;comment:;"`                                                      //表面积
	GraphPaper          *float64 `json:"graphPaper" form:"graphPaper" gorm:"column:graph_paper;comment:;"`                                //图纸
	GPDateOrVersion     string   `json:"GPDateOrVersion" form:"GPDateOrVersion" gorm:"column:GP_date_or_version;comment:;size:255;"`      //图纸日期或版本
	GPAudit             string   `json:"GPAudit" form:"GPAudit" gorm:"column:gp_audit;comment:;size:255;"`                                //图纸评审
	OE                  string   `json:"OE" form:"OE" gorm:"column:OE;comment:;size:255;"`                                                //主机厂
	OEMStandCode        string   `json:"OEMStandCode" form:"OEMStandCode" gorm:"column:OEM_stand_code;comment:;size:255;"`                //车厂标准代号
	OEMStand            string   `json:"OEMStand" form:"OEMStand" gorm:"column:OEM_stand;comment:;size:255;"`                             //车厂标准
	OEMStandardReview   string   `json:"OEMStandardReview" form:"OEMStandardReview" gorm:"column:OEM_standard_review;comment:;size:255;"` //车厂标准评审
	Technique           *int     `json:"technique" form:"technique" gorm:"column:technique;comment:;size:10;"`                            //工艺、设备及治具基本信息
	SurfaceInfo         *int     `json:"surfaceInfo" form:"surfaceInfo" gorm:"column:surface_info;comment:;"`                             //表面处理信息
	QualityInfo         *int     `json:"qualityInfo" form:"qualityInfo" gorm:"column:quality_info;comment:;"`                             //质量要求信息
	CoastInfo           *int     `json:"coastInfo" form:"coastInfo" gorm:"column:coast_info;comment:;"`                                   //成本信息收集信息
	ImportInfo          *int     `json:"importInfo" form:"importInfo" gorm:"column:import_info;comment:;"`                                //产品导入信息
	ProcessInfo         *int     `json:"processInfo" form:"processInfo" gorm:"column:process_info;comment:;"`                             //工艺文件信息
	TransferInformation *int     `json:"transferInformation" form:"transferInformation" gorm:"column:transfer_information;comment:;"`     //量产转移信息
	ChangeInfo          *int     `json:"changeInfo" form:"changeInfo" gorm:"column:change_info;comment:;"`                                //工程变更信息
	PackageInfo         *int     `json:"packageInfo" form:"packageInfo" gorm:"column:package_info;comment:;"`                             //包装要求
	ProductivityInfo    *int     `json:"productivityInfo" form:"productivityInfo" gorm:"column:productivity_info;comment:;"`              //效率相关信息
	CreatedBy           uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy           uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy           uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 技术部基础信息 TecBaseInfo自定义表名 tec_base_info
func (TecBaseInfo) TableName() string {
	return "tec_base_info"
}
