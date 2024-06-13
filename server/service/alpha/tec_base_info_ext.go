package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type TecBaseInfoExtService struct {
}

// CreateTecBaseInfoExt 创建零件基础信息扩展记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoExtService *TecBaseInfoExtService) CreateTecBaseInfoExt(tecBaseInfoExt *alpha.TecBaseInfoExt) (err error) {
	err = global.GVA_DB.Create(tecBaseInfoExt).Error
	return err
}

// DeleteTecBaseInfoExt 删除零件基础信息扩展记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoExtService *TecBaseInfoExtService) DeleteTecBaseInfoExt(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TecBaseInfoExt{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.TecBaseInfoExt{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTecBaseInfoExtByIds 批量删除零件基础信息扩展记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoExtService *TecBaseInfoExtService) DeleteTecBaseInfoExtByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TecBaseInfoExt{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.TecBaseInfoExt{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTecBaseInfoExt 更新零件基础信息扩展记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoExtService *TecBaseInfoExtService) UpdateTecBaseInfoExt(tecBaseInfoExt alpha.TecBaseInfoExt) (err error) {
	err = global.GVA_DB.Save(&tecBaseInfoExt).Error
	return err
}

// GetTecBaseInfoExt 根据ID获取零件基础信息扩展记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoExtService *TecBaseInfoExtService) GetTecBaseInfoExt(ID string) (tecBaseInfoExt alpha.TecBaseInfoExt, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&tecBaseInfoExt).Error
	return
}

// GetTecBaseInfoExtInfoList 分页获取零件基础信息扩展记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoExtService *TecBaseInfoExtService) GetTecBaseInfoExtInfoList(info alphaReq.TecBaseInfoExtSearch) (list []alpha.TecBaseInfoExt, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.TecBaseInfoExt{})
	var tecBaseInfoExts []alpha.TecBaseInfoExt
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ME002 != "" {
		db = db.Where("ME002 LIKE ?", "%"+info.ME002+"%")
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Typeext != nil {
		db = db.Where("typeext = ?", info.Typeext)
	}
	if info.MB201 != "" {
		db = db.Where("MB201 LIKE ?", "%"+info.MB201+"%")
	}
	if info.MB002 != "" {
		db = db.Where("MB002 LIKE ?", "%"+info.MB002+"%")
	}
	if info.MB202 != "" {
		db = db.Where("MB202 LIKE ?", "%"+info.MB202+"%")
	}
	if info.MB003 != "" {
		db = db.Where("MB003 LIKE ?", "%"+info.MB003+"%")
	}
	if info.Query != "" {
		db = db.Or("UTN LIKE ?", "%"+info.Query+"%").
			Or("MB201 LIKE ?", "%"+info.Query+"%").Or("MB002 LIKE ?", "%"+info.Query+"%").
			Or("MB202 LIKE ?", "%"+info.Query+"%").Or("MB003 LIKE ?", "%"+info.Query+"%")
	}
	//if info.OE != nil {
	//   db = db.Where("OE LIKE ?","%"+info.OE+"%")
	//}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tecBaseInfoExts).Error
	return tecBaseInfoExts, total, err
}

// CheckIsDuplicate 检查是否有重复录入
func (tecBaseInfoExtService *TecBaseInfoExtService) CheckIsDuplicate(info alphaReq.TecBaseInfoExtSearch) (result bool, err error) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	//// 创建db
	db := global.GVA_DB.Model(&alpha.TecBaseInfoExt{})
	//var tecBaseInfoExts []alpha.TecBaseInfoExt
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.MB201 != "" {
		db = db.Where("MB201 = ?", info.MB201)
	}

	if info.MB202 != "" {
		db = db.Where("MB202 = ?", info.MB202)
	}

	//if info.OE != nil {
	//   db = db.Where("OE LIKE ?","%"+info.OE+"%")
	//}
	var total int64
	db.Count(&total)

	if total > 0 {
		return true, err
	} else {
		return false, err
	}
}
