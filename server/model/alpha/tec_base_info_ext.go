// 自动生成模板TecBaseInfoExt
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 零件基础信息扩展 结构体  TecBaseInfoExt
type TecBaseInfoExt struct {
	global.GVA_MODEL
	ME002                    string   `json:"ME002" form:"ME002" gorm:"column:ME002;comment:;size:20;" binding:"required"`                                 //分部
	Type                     *int     `json:"type" form:"type" gorm:"column:type;comment:;" binding:"required"`                                            //零件类型1
	Typeext                  *int     `json:"typeext" form:"typeext" gorm:"column:typeext;comment:;"`                                                      //零件类型2
	MB201                    string   `json:"MB201" form:"MB201" gorm:"column:MB201;comment:;size:255;" binding:"required"`                                //客户名称
	MB002                    string   `json:"MB002" form:"MB002" gorm:"column:MB002;comment:;size:255;" binding:"required"`                                //零件名称
	MB202                    string   `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:255;" binding:"required"`                                //客户品号
	MB003                    string   `json:"MB003" form:"MB003" gorm:"column:MB003;comment:;size:255;" binding:"required"`                                //零件规格
	FigureNumber             string   `json:"figureNumber" form:"figureNumber" gorm:"column:figure_number;comment:;size:255;"`                             //零件图号
	MB004                    string   `json:"MB004" form:"MB004" gorm:"column:MB004;comment:;" binding:"required"`                                         //计价单位
	UTN                      string   `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:255;" binding:"required"`                                      //唯一追踪号
	Textrue                  string   `json:"textrue" form:"textrue" gorm:"column:textrue;comment:;size:255;"`                                             //材质
	Level                    string   `json:"level" form:"level" gorm:"column:level;comment:;size:255;"`                                                   //硬度/机械等级
	Length                   *float64 `json:"length" form:"length" gorm:"column:length;comment:;"`                                                         //长
	Width                    *float64 `json:"width" form:"width" gorm:"column:width;comment:;"`                                                            //宽
	Height                   *float64 `json:"height" form:"height" gorm:"column:height;comment:;"`                                                         //高
	Diameter                 *float64 `json:"diameter" form:"diameter" gorm:"column:diameter;comment:;"`                                                   //最大直径
	Thick                    *float64 `json:"thick" form:"thick" gorm:"column:thick;comment:;"`                                                            //料厚
	WireDiameter             *float64 `json:"wireDiameter" form:"wireDiameter" gorm:"column:wire_diameter;comment:;"`                                      //线径
	PieceWeight              *float64 `json:"pieceWeight" form:"pieceWeight" gorm:"column:piece_weight;comment:;"`                                         //单重
	DrawWeight               *float64 `json:"drawWeight" form:"drawWeight" gorm:"column:draw_weight;comment:;"`                                            //单重
	BET                      *float64 `json:"BET" form:"BET" gorm:"column:BET;comment:;"`                                                                  //表面积
	GraphPaper               *float64 `json:"graphPaper" form:"graphPaper" gorm:"column:graph_paper;comment:;"`                                            //图纸
	GPDateOrVersion          string   `json:"GPDateOrVersion" form:"GPDateOrVersion" gorm:"column:GP_date_or_version;comment:;size:255;"`                  //图纸日期或版本
	GPAudit                  string   `json:"GPAudit" form:"GPAudit" gorm:"column:gp_audit;comment:;size:255;"`                                            //图纸评审
	OE                       *int     `json:"OE" form:"OE" gorm:"column:OE;comment:;size:255;"`                                                            //主机厂
	OEMStandCode             string   `json:"OEMStandCode" form:"OEMStandCode" gorm:"column:OEM_stand_code;comment:;size:255;"`                            //车厂标准代号
	OEMStand                 string   `json:"OEMStand" form:"OEMStand" gorm:"column:OEM_stand;comment:;size:255;"`                                         //车厂标准
	OEMStandardReview        string   `json:"OEMStandardReview" form:"OEMStandardReview" gorm:"column:OEM_standard_review;comment:;size:255;"`             //车厂标准评审
	ProcessMode              string   `json:"processMode" form:"processMode" gorm:"column:process_mode;comment:;size:20;" binding:"required"`              //工艺方式
	Unoil                    string   `json:"unoil" form:"unoil" gorm:"column:unoil;comment:;size:20;"`                                                    //除油
	ShotBlasting             string   `json:"shotBlasting" form:"shotBlasting" gorm:"column:shot_blasting;comment:;"`                                      //抛丸
	Phosphating              string   `json:"phosphating" form:"phosphating" gorm:"column:phosphating;comment:;size:20;"`                                  //磷化
	Electroplate             string   `json:"electroplate" form:"electroplate" gorm:"column:electroplate;comment:;size:20;"`                               //电镀
	CoatingSpecification     string   `json:"coatingSpecification" form:"coatingSpecification" gorm:"column:coating_specification;comment:;size:255;"`     //涂覆规范
	BaseCoat                 string   `json:"baseCoat" form:"baseCoat" gorm:"column:base_coat;comment:;size:255;"`                                         //底涂
	NumberOfUndercoats       string   `json:"numberOfUndercoats" form:"numberOfUndercoats" gorm:"column:number_of_undercoats;comment:;size:20;"`           //底涂次数
	Topcoat                  string   `json:"topcoat" form:"topcoat" gorm:"column:topcoat;comment:;size:255;"`                                             //面涂
	NumberOfTopcoats         string   `json:"numberOfTopcoats" form:"numberOfTopcoats" gorm:"column:number_of_topcoats;comment:;size:20;"`                 //面涂次数
	SpecialProcess           *int     `json:"specialProcess" form:"specialProcess" gorm:"column:special_process;comment:;"`                                //特殊工序
	SpecialProcessParameters *int     `json:"specialProcessParameters" form:"specialProcessParameters" gorm:"column:special_process_parameters;comment:;"` //特殊工序工艺参数
	CreatedBy                uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy                uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy                uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 零件基础信息扩展 TecBaseInfoExt自定义表名 tec_base_info_ext
func (TecBaseInfoExt) TableName() string {
	return "tec_base_info_ext"
}
