package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type MdThirdLevelService struct {
}

// CreateMdThirdLevel 创建层级三记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdThirdLevelService *MdThirdLevelService) CreateMdThirdLevel(mdThirdLevel *alpha.MdThirdLevel) (err error) {
	err = global.GVA_DB.Create(mdThirdLevel).Error
	return err
}

// DeleteMdThirdLevel 删除层级三记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdThirdLevelService *MdThirdLevelService) DeleteMdThirdLevel(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.MdThirdLevel{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.MdThirdLevel{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMdThirdLevelByIds 批量删除层级三记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdThirdLevelService *MdThirdLevelService) DeleteMdThirdLevelByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.MdThirdLevel{}).Where("id in (?)", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.MdThirdLevel{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMdThirdLevel 更新层级三记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdThirdLevelService *MdThirdLevelService) UpdateMdThirdLevel(mdThirdLevel alpha.MdThirdLevel) (err error) {
	err = global.GVA_DB.Save(&mdThirdLevel).Error
	return err
}

// GetMdThirdLevel 根据ID获取层级三记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdThirdLevelService *MdThirdLevelService) GetMdThirdLevel(ID string) (mdThirdLevel alpha.MdThirdLevel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mdThirdLevel).Error
	return
}

// GetMdThirdLevelInfoList 分页获取层级三记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdThirdLevelService *MdThirdLevelService) GetMdThirdLevelInfoList(info alphaReq.MdThirdLevelSearch) (list []alpha.MdThirdLevel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.MdThirdLevel{})
	var mdThirdLevels []alpha.MdThirdLevel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.SecondLevelId != 0 {
		db = db.Where("secondLevel_id = ?", info.SecondLevelId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Query != "" {
		db = db.Or("name LIKE ?", "%"+info.Query+"%").Or("remark LIKE ?", "%"+info.Query+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mdThirdLevels).Error
	return mdThirdLevels, total, err
}
