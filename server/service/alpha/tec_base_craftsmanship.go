package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type TecBaseCraftsmanshipService struct{}

// CreateTecBaseCraftsmanship 创建入篮量和程序号记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseCraftsmanshipService *TecBaseCraftsmanshipService) CreateTecBaseCraftsmanship(tecBaseCraftsmanship *[]alpha.TecBaseCraftsmanship) (err error) {
	//err = global.GVA_DB.Create(tecBaseCraftsmanship).Error
	//批量新增数据
	err = global.GVA_DB.Create(tecBaseCraftsmanship).Error
	return err
}

// DeleteTecBaseCraftsmanship 删除入篮量和程序号记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseCraftsmanshipService *TecBaseCraftsmanshipService) DeleteTecBaseCraftsmanship(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TecBaseCraftsmanship{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.TecBaseCraftsmanship{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTecBaseCraftsmanshipByIds 批量删除入篮量和程序号记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseCraftsmanshipService *TecBaseCraftsmanshipService) DeleteTecBaseCraftsmanshipByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TecBaseCraftsmanship{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.TecBaseCraftsmanship{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTecBaseCraftsmanship 更新入篮量和程序号记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseCraftsmanshipService *TecBaseCraftsmanshipService) UpdateTecBaseCraftsmanship(tecBaseCraftsmanship []alpha.TecBaseCraftsmanship) (err error) {
	for _, craftsmanship := range tecBaseCraftsmanship {
		err = global.GVA_DB.Model(&alpha.TecBaseCraftsmanship{}).Where("id = ?", craftsmanship.ID).Updates(&craftsmanship).Error
	}

	return err
}

// GetTecBaseCraftsmanship 根据ID获取入篮量和程序号记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseCraftsmanshipService *TecBaseCraftsmanshipService) GetTecBaseCraftsmanship(parentId string) (tecBaseCraftsmanship []alpha.TecBaseCraftsmanship, err error) {
	err = global.GVA_DB.Where("parent_id = ?", parentId).Find(&tecBaseCraftsmanship).Error
	return
}

// GetTecBaseCraftsmanshipInfoList 分页获取入篮量和程序号记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseCraftsmanshipService *TecBaseCraftsmanshipService) GetTecBaseCraftsmanshipInfoList(info alphaReq.TecBaseCraftsmanshipSearch) (list []alpha.TecBaseCraftsmanship, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.TecBaseCraftsmanship{})
	var tecBaseCraftsmanships []alpha.TecBaseCraftsmanship
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UTN != "" {
		db = db.Where("UTN LIKE ?", "%"+info.UTN+"%")
	}
	if info.MB202 != "" {
		db = db.Where("MB202 LIKE ?", "%"+info.MB202+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tecBaseCraftsmanships).Error
	return tecBaseCraftsmanships, total, err
}
func (tecBaseCraftsmanshipService *TecBaseCraftsmanshipService) GetTecBaseCraftsmanshipDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	productLine := make([]map[string]any, 0)
	global.GVA_DB.Table("product_line").Select("product_name as label,id as value").Scan(&productLine)
	res["productLine"] = productLine
	return
}
