package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type QualityBaseInfoService struct {
}

// CreateQualityBaseInfo 创建质量要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (qualityBaseInfoService *QualityBaseInfoService) CreateQualityBaseInfo(qualityBaseInfo *alpha.QualityBaseInfo) (err error) {
	err = global.GVA_DB.Create(qualityBaseInfo).Error
	return err
}

// DeleteQualityBaseInfo 删除质量要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (qualityBaseInfoService *QualityBaseInfoService) DeleteQualityBaseInfo(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.QualityBaseInfo{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.QualityBaseInfo{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteQualityBaseInfoByIds 批量删除质量要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (qualityBaseInfoService *QualityBaseInfoService) DeleteQualityBaseInfoByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.QualityBaseInfo{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.QualityBaseInfo{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateQualityBaseInfo 更新质量要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (qualityBaseInfoService *QualityBaseInfoService) UpdateQualityBaseInfo(qualityBaseInfo alpha.QualityBaseInfo) (err error) {
	err = global.GVA_DB.Save(&qualityBaseInfo).Error
	return err
}

// GetQualityBaseInfo 根据ID获取质量要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (qualityBaseInfoService *QualityBaseInfoService) GetQualityBaseInfo(ID string) (qualityBaseInfo alpha.QualityBaseInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&qualityBaseInfo).Error
	return
}

// GetQualityBaseInfoInfoList 分页获取质量要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (qualityBaseInfoService *QualityBaseInfoService) GetQualityBaseInfoInfoList(info alphaReq.QualityBaseInfoSearch) (list []alpha.QualityBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.QualityBaseInfo{})
	var qualityBaseInfos []alpha.QualityBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UTN != "" {
		db = db.Where("UTN LIKE ?", "%"+info.UTN+"%")
	}
	if info.MB202 != "" {
		db = db.Where("MB001 LIKE ?", "%"+info.MB202+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&qualityBaseInfos).Error
	return qualityBaseInfos, total, err
}
