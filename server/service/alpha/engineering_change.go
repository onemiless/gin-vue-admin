package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type EngineeringChangeService struct {
}

// CreateEngineeringChange 创建工程变更信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (engineeringChangeService *EngineeringChangeService) CreateEngineeringChange(engineeringChange *alpha.EngineeringChange) (err error) {
	err = global.GVA_DB.Create(engineeringChange).Error
	return err
}

// DeleteEngineeringChange 删除工程变更信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (engineeringChangeService *EngineeringChangeService) DeleteEngineeringChange(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.EngineeringChange{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.EngineeringChange{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteEngineeringChangeByIds 批量删除工程变更信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (engineeringChangeService *EngineeringChangeService) DeleteEngineeringChangeByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.EngineeringChange{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.EngineeringChange{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateEngineeringChange 更新工程变更信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (engineeringChangeService *EngineeringChangeService) UpdateEngineeringChange(engineeringChange alpha.EngineeringChange) (err error) {
	err = global.GVA_DB.Model(&alpha.EngineeringChange{}).Where("id = ?", engineeringChange.ID).Updates(&engineeringChange).Error
	return err
}

// GetEngineeringChange 根据ID获取工程变更信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (engineeringChangeService *EngineeringChangeService) GetEngineeringChange(ID string) (engineeringChange alpha.EngineeringChange, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&engineeringChange).Error
	return
}

// GetEngineeringChangeInfoList 分页获取工程变更信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (engineeringChangeService *EngineeringChangeService) GetEngineeringChangeInfoList(info alphaReq.EngineeringChangeSearch) (list []alpha.EngineeringChange, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.EngineeringChange{})
	var engineeringChanges []alpha.EngineeringChange
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
	if info.SN != "" {
		db = db.Where("SN LIKE ?", "%"+info.SN+"%")
	}
	if info.StartChangeDate != nil && info.EndChangeDate != nil {
		db = db.Where("change_date BETWEEN ? AND ? ", info.StartChangeDate, info.EndChangeDate)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&engineeringChanges).Error
	return engineeringChanges, total, err
}
