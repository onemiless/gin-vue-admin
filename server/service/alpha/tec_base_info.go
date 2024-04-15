package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type TecBaseInfoService struct {
}

// CreateTecBaseInfo 创建技术部基础信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoService *TecBaseInfoService) CreateTecBaseInfo(tecBaseInfo *alpha.TecBaseInfo) (err error) {
	err = global.GVA_DB.Create(tecBaseInfo).Error
	return err
}

// DeleteTecBaseInfo 删除技术部基础信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoService *TecBaseInfoService) DeleteTecBaseInfo(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TecBaseInfo{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.TecBaseInfo{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTecBaseInfoByIds 批量删除技术部基础信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoService *TecBaseInfoService) DeleteTecBaseInfoByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TecBaseInfo{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.TecBaseInfo{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTecBaseInfo 更新技术部基础信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoService *TecBaseInfoService) UpdateTecBaseInfo(tecBaseInfo alpha.TecBaseInfo) (err error) {
	err = global.GVA_DB.Save(&tecBaseInfo).Error
	return err
}

// GetTecBaseInfo 根据ID获取技术部基础信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoService *TecBaseInfoService) GetTecBaseInfo(ID string) (tecBaseInfo alpha.TecBaseInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&tecBaseInfo).Error
	return
}

// GetTecBaseInfoInfoList 分页获取技术部基础信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseInfoService *TecBaseInfoService) GetTecBaseInfoInfoList(info alphaReq.TecBaseInfoSearch) (list []alpha.TecBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.TecBaseInfo{})
	var tecBaseInfos []alpha.TecBaseInfo
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
	if info.OE != "" {
		db = db.Where("OE LIKE ?", "%"+info.OE+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tecBaseInfos).Error
	return tecBaseInfos, total, err
}
