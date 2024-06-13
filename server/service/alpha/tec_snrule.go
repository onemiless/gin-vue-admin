package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
)

//var (
//	mdClient = service.ServiceGroupApp.AlphaServiceGroup.MdClientService
//)

type TecSNRuleService struct {
}

// GetTecSNRule 获取编码规则
// Author [piexlmax](https://github.com/piexlmax)
func (tecSNRuleService *TecSNRuleService) GetTecSNRule(info alphaReq.TecSNRuleSearch) (result string, err error) {
	level := alpha.MdThirdLevel{}
	mdClient := MdClientService{}
	//根据名称查询层级2规则对应名称
	secondLevel := alpha.MdSecondLevel{}
	err = global.GVA_DB.Where("name = ?", info.Name).First(&secondLevel).Error
	if err == nil {
		err = global.GVA_DB.Where("branch = ? and secondLevel_id = ?", info.Branch, secondLevel.ID).First(&level).Error
		if err == nil {
			if level.IsCustomer == "1" {
				//查询客户信息表
				client, err := mdClient.GetMdClient(info.Customer)
				if err == nil {
					//获取客户首字母
					clientType := client.ClientType
					//拼接字符串
					result = level.Name + "-" + clientType + "-"
				}

			} else {
				result = level.Name
			}
		}
	}
	return
}

//
//// GetMdSecondLevelInfoList 分页获取层级二记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (mdSecondLevelService *MdSecondLevelService) GetMdSecondLevelInfoList(info alphaReq.MdSecondLevelSearch) (list []alpha.MdSecondLevel, total int64, err error) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.GVA_DB.Model(&alpha.MdSecondLevel{})
//	var mdSecondLevels []alpha.MdSecondLevel
//	// 如果有条件搜索 下方会自动创建搜索语句
//	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
//		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
//	}
//	if info.FirstLevelId != 0 {
//		db = db.Where("firstLevel_id = ?", info.FirstLevelId)
//	}
//	if info.Name != "" {
//		db = db.Or("name LIKE ?", "%"+info.Name+"%")
//	}
//	if info.Remark != "" {
//		db = db.Or("remark LIKE ?", "%"+info.Remark+"%")
//	}
//	if info.Query != "" {
//		db = db.Where("name LIKE ?", "%"+info.Query+"%")
//
//	}
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//
//	if limit != 0 {
//		db = db.Limit(limit).Offset(offset)
//	}
//
//	err = db.Find(&mdSecondLevels).Error
//	return mdSecondLevels, total, err
//}
