// 自动生成模板QualityBaseInfo
package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 质量要求信息 结构体  QualityBaseInfo
type QualityBaseInfo struct {
	global.GVA_MODEL
	ThicknessBaseCoat               string `json:"thicknessBaseCoat" form:"thicknessBaseCoat" gorm:"column:thickness_base_coat;comment:;size:255;"`                                           //底涂
	ThicknessTopCoat                string `json:"thicknessTopCoat" form:"thicknessTopCoat" gorm:"column:thickness_top_coat;comment:;size:255;"`                                              //面涂
	ThicknessTotal                  string `json:"thicknessTotal" form:"thicknessTotal" gorm:"column:thickness_total;comment:;size:255;"`                                                     //总
	ThicknessTeststandard           string `json:"thicknessTeststandard" form:"thicknessTeststandard" gorm:"column:thickness_teststandard;comment:;size:255;"`                                //测试标准
	CoatweightBaseCoat              string `json:"coatweightBaseCoat" form:"coatweightBaseCoat" gorm:"column:coatweight_base_coat;comment:;size:255;"`                                        //底涂
	CoatweightTopcoat               string `json:"coatweightTopcoat" form:"coatweightTopcoat" gorm:"column:coatweight_topcoat;comment:;size:255;"`                                            //面涂
	CoatweightTeststandard          string `json:"coatweightTeststandard" form:"coatweightTeststandard" gorm:"column:coatweight_teststandard;comment:;size:255;"`                             //测试标准
	AdhesionRequirement             string `json:"adhesionRequirement" form:"adhesionRequirement" gorm:"column:adhesion_requirement;comment:;size:255;"`                                      //附着力要求
	AdhesionTeststandard            string `json:"adhesionTeststandard" form:"adhesionTeststandard" gorm:"column:adhesion_teststandard;comment:;size:255;"`                                   //附着力测试标准
	SizeWorkblank                   string `json:"sizeWorkblank" form:"sizeWorkblank" gorm:"column:size_workblank;comment:;size:255;"`                                                        //毛坯
	SizeFinshedproduct              string `json:"sizeFinshedproduct" form:"sizeFinshedproduct" gorm:"column:size_finshedproduct;comment:;size:255;"`                                         //成品
	NSSRedrust                      string `json:"NSSRedrust" form:"NSSRedrust" gorm:"column:NSS_redrust;comment:;size:255;"`                                                                 //红锈
	NSSWhiterust                    string `json:"NSSWhiterust" form:"NSSWhiterust" gorm:"column:NSS_whiterust;comment:;size:255;"`                                                           //白锈
	NSSTeststandard                 string `json:"NSSTeststanddard" form:"NSSTeststanddard" gorm:"column:NSS_teststanddard;comment:;size:255;"`                                               //中性盐雾测试标准
	HotstorageRequirement           string `json:"hotstorageRequirement" form:"hotstorageRequirement" gorm:"column:hotstorage_requirement;comment:;size:255;"`                                //热存放标准
	HotstorageTeststandard          string `json:"hotstorageTeststandard" form:"hotstorageTeststandard" gorm:"column:hotstorage_teststandard;comment:;size:255;"`                             //测试标准
	HotstorageCondensedWaterTest    string `json:"hotstorageCondensedWaterTest" form:"hotstorageCondensedWaterTest" gorm:"column:hotstorage_condensed_water_test;comment:;size:255;"`         //耐水测试
	HotstorageWaterTestStandard     string `json:"hotstorageWaterTestStandard" form:"hotstorageWaterTestStandard" gorm:"column:hotstorage_water_test_standard;comment:;size:255;"`            //耐水测试标准
	CyclicCorrosion                 string `json:"cyclicCorrosion" form:"cyclicCorrosion" gorm:"column:cyclic_corrosion;comment:;size:255;"`                                                  //循环腐蚀
	CyclicCorrosionTeststandard     string `json:"cyclicCorrosionTeststandard" form:"cyclicCorrosionTeststandard" gorm:"column:cyclic_corrosion_teststandard;comment:;size:255;"`             //循环腐蚀测试标准
	FrictionCoefficient             string `json:"frictionCoefficient" form:"frictionCoefficient" gorm:"column:friction_coefficient;comment:;size:255;"`                                      //摩擦系数
	FrictionCoefficientTeststandard string `json:"frictionCoefficientTeststandard" form:"frictionCoefficientTeststandard" gorm:"column:friction_coefficient_teststandard;comment:;size:255;"` //摩擦系数测试标准
	Torque                          string `json:"torque" form:"torque" gorm:"column:torque;comment:;size:255;"`                                                                              //扭矩
	TorqueTestStandard              string `json:"torqueTestStandard" form:"torqueTestStandard" gorm:"column:torque_test_standard;comment:;size:255;"`                                        //扭矩测试标准
	DehydroDetection                string `json:"dehydroDetection" form:"dehydroDetection" gorm:"column:dehydro_detection;comment:;size:255;"`                                               //氢脆检测
	DehydroDetectionStandard        string `json:"dehydroDetectionStandard" form:"dehydroDetectionStandard" gorm:"column:dehydro_detection_standard;comment:;size:255;"`                      //氢脆检测标准
	BlankPassGauge                  string `json:"blankPassGauge" form:"blankPassGauge" gorm:"column:blank_pass_gauge;comment:;size:255;"`                                                    //毛坯通规
	BlankStopGauge                  string `json:"blankStopGauge" form:"blankStopGauge" gorm:"column:blank_stop_gauge;comment:;size:255;"`                                                    //毛坯止规
	FinishedGauge                   string `json:"finishedGauge" form:"finishedGauge" gorm:"column:finished_gauge;comment:;size:255;"`                                                        //成品通规
	FinishedStopGauge               string `json:"finishedStopGauge" form:"finishedStopGauge" gorm:"column:finished_stop_gauge;comment:;size:255;"`                                           //成品止规
	Cleanliness                     string `json:"cleanliness" form:"cleanliness" gorm:"column:cleanliness;comment:;size:255;"`                                                               //清洁度
	CleanlinessTestStandrad         string `json:"cleanlinessTestStandrad" form:"cleanlinessTestStandrad" gorm:"column:cleanliness_test_standrad;comment:;size:255;"`                         //测试标准
	RoughnessRough                  string `json:"roughnessRough" form:"roughnessRough" gorm:"column:roughness_rough;comment:;size:255;"`                                                     //毛坯
	RoughnessFinish                 string `json:"roughnessFinish" form:"roughnessFinish" gorm:"column:roughness_finish;comment:;size:255;"`                                                  //成品
	ChemicalResistance              string `json:"chemicalResistance" form:"chemicalResistance" gorm:"column:chemical_resistance;comment:;size:255;"`                                         //耐化学抵抗
	ChemicalResistanceTestStandard  string `json:"chemicalResistanceTestStandard" form:"chemicalResistanceTestStandard" gorm:"column:chemical_resistance_test_standard;comment:;size:255;"`   //耐化学抵抗测试标准
	RubbleImpact                    string `json:"rubbleImpact" form:"rubbleImpact" gorm:"column:rubble_impact;comment:;size:255;"`                                                           //碎石冲击
	RubbleImpactTestStandard        string `json:"rubbleImpactTestStandard" form:"rubbleImpactTestStandard" gorm:"column:rubble_impact_test_standard;comment:;size:255;"`                     //测试标准
	ParentId                        *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;" binding:"required"`                                                             //基础信息id
	UTN                             string `json:"UTN" form:"UTN" gorm:"column:UTN;comment:;size:255;" binding:"required"`                                                                    //唯一追踪号
	MB202                           string `json:"MB202" form:"MB202" gorm:"column:MB202;comment:;size:255;"`                                                                                 //客户品号
	CreatedBy                       uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy                       uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy                       uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 质量要求信息 QualityBaseInfo自定义表名 quality_base_info
func (QualityBaseInfo) TableName() string {
	return "quality_base_info"
}
