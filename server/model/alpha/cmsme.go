// 自动生成模板CMSME
package alpha

// ERP部门 结构体  CMSME
type CMSME struct {
	COMPANY    string   `json:"COMPANY" form:"COMPANY" gorm:"column:COMPANY;comment:;size:10;"`           //COMPANY字段
	CREATOR    string   `json:"CREATOR" form:"CREATOR" gorm:"column:CREATOR;comment:;size:10;"`           //CREATOR字段
	USRGROUP   string   `json:"USRGROUP" form:"USRGROUP" gorm:"column:USR_GROUP;comment:;size:10;"`       //USRGROUP字段
	CREATEDATE string   `json:"CREATEDATE" form:"CREATEDATE" gorm:"column:CREATE_DATE;comment:;size:17;"` //CREATEDATE字段
	MODIFIER   string   `json:"MODIFIER" form:"MODIFIER" gorm:"column:MODIFIER;comment:;size:10;"`        //MODIFIER字段
	MODIDATE   string   `json:"MODIDATE" form:"MODIDATE" gorm:"column:MODI_DATE;comment:;size:17;"`       //MODIDATE字段
	FLAG       *float64 `json:"FLAG" form:"FLAG" gorm:"column:FLAG;comment:;"`                            //FLAG字段
	ME001      string   `json:"ME001" form:"ME001" gorm:"primarykey;column:ME001;comment:;size:6;"`       //ME001字段
	ME002      string   `json:"ME002" form:"ME002" gorm:"column:ME002;comment:;size:20;"`                 //ME002字段
	ME003      string   `json:"ME003" form:"ME003" gorm:"column:ME003;comment:;size:255;"`                //ME003字段
	ME004      string   `json:"ME004" form:"ME004" gorm:"column:ME004;comment:;size:20;"`                 //ME004字段
	ME005      string   `json:"ME005" form:"ME005" gorm:"column:ME005;comment:;size:1;"`                  //ME005字段
	ME006      string   `json:"ME006" form:"ME006" gorm:"column:ME006;comment:;size:1;"`                  //ME006字段
	ME007      string   `json:"ME007" form:"ME007" gorm:"column:ME007;comment:;size:12;"`                 //ME007字段
	ME008      string   `json:"ME008" form:"ME008" gorm:"column:ME008;comment:;size:1;"`                  //ME008字段
	ME009      string   `json:"ME009" form:"ME009" gorm:"column:ME009;comment:;size:1;"`                  //ME009字段
	ME010      string   `json:"ME010" form:"ME010" gorm:"column:ME010;comment:;size:8;"`                  //ME010字段
	ME011      string   `json:"ME011" form:"ME011" gorm:"column:ME011;comment:;size:30;"`                 //ME011字段
	ME012      *float64 `json:"ME012" form:"ME012" gorm:"column:ME012;comment:;"`                         //ME012字段
	ME013      *float64 `json:"ME013" form:"ME013" gorm:"column:ME013;comment:;"`                         //ME013字段
	ME014      *float64 `json:"ME014" form:"ME014" gorm:"column:ME014;comment:;"`                         //ME014字段
	ME015      string   `json:"ME015" form:"ME015" gorm:"column:ME015;comment:;size:1;"`                  //ME015字段
	UDF01      string   `json:"UDF01" form:"UDF01" gorm:"column:UDF01;comment:;size:255;"`                //UDF01字段
	UDF02      string   `json:"UDF02" form:"UDF02" gorm:"column:UDF02;comment:;size:255;"`                //UDF02字段
	UDF03      string   `json:"UDF03" form:"UDF03" gorm:"column:UDF03;comment:;size:255;"`                //UDF03字段
	UDF04      string   `json:"UDF04" form:"UDF04" gorm:"column:UDF04;comment:;size:255;"`                //UDF04字段
	UDF05      string   `json:"UDF05" form:"UDF05" gorm:"column:UDF05;comment:;size:255;"`                //UDF05字段
	UDF06      string   `json:"UDF06" form:"UDF06" gorm:"column:UDF06;comment:;size:255;"`                //UDF06字段
	UDF51      *float64 `json:"UDF51" form:"UDF51" gorm:"column:UDF51;comment:;"`                         //UDF51字段
	UDF52      *float64 `json:"UDF52" form:"UDF52" gorm:"column:UDF52;comment:;"`                         //UDF52字段
	UDF53      *float64 `json:"UDF53" form:"UDF53" gorm:"column:UDF53;comment:;"`                         //UDF53字段
	UDF54      *float64 `json:"UDF54" form:"UDF54" gorm:"column:UDF54;comment:;"`                         //UDF54字段
	UDF55      *float64 `json:"UDF55" form:"UDF55" gorm:"column:UDF55;comment:;"`                         //UDF55字段
	UDF56      *float64 `json:"UDF56" form:"UDF56" gorm:"column:UDF56;comment:;"`                         //UDF56字段
	UDF07      string   `json:"UDF07" form:"UDF07" gorm:"column:UDF07;comment:;size:255;"`                //UDF07字段
	UDF08      string   `json:"UDF08" form:"UDF08" gorm:"column:UDF08;comment:;size:255;"`                //UDF08字段
	UDF09      string   `json:"UDF09" form:"UDF09" gorm:"column:UDF09;comment:;size:255;"`                //UDF09字段
	UDF10      string   `json:"UDF10" form:"UDF10" gorm:"column:UDF10;comment:;size:255;"`                //UDF10字段
	UDF11      string   `json:"UDF11" form:"UDF11" gorm:"column:UDF11;comment:;size:255;"`                //UDF11字段
	UDF12      string   `json:"UDF12" form:"UDF12" gorm:"column:UDF12;comment:;size:255;"`                //UDF12字段
	UDF57      *float64 `json:"UDF57" form:"UDF57" gorm:"column:UDF57;comment:;"`                         //UDF57字段
	UDF58      *float64 `json:"UDF58" form:"UDF58" gorm:"column:UDF58;comment:;"`                         //UDF58字段
	UDF59      *float64 `json:"UDF59" form:"UDF59" gorm:"column:UDF59;comment:;"`                         //UDF59字段
	UDF60      *float64 `json:"UDF60" form:"UDF60" gorm:"column:UDF60;comment:;"`                         //UDF60字段
	UDF61      *float64 `json:"UDF61" form:"UDF61" gorm:"column:UDF61;comment:;"`                         //UDF61字段
	UDF62      *float64 `json:"UDF62" form:"UDF62" gorm:"column:UDF62;comment:;"`                         //UDF62字段
}

// TableName ERP部门 CMSME自定义表名 CMSME
func (CMSME) TableName() string {
	return "CMSME"
}
