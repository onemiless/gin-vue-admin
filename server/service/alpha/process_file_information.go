package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type ProcessFileInformationService struct {
}

// CreateProcessFileInformation 创建工艺文件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (processFileInformationService *ProcessFileInformationService) CreateProcessFileInformation(processFileInformation *alpha.ProcessFileInformation) (err error) {
	err = global.GVA_DB.Create(processFileInformation).Error
	return err
}

// DeleteProcessFileInformation 删除工艺文件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (processFileInformationService *ProcessFileInformationService) DeleteProcessFileInformation(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.ProcessFileInformation{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.ProcessFileInformation{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteProcessFileInformationByIds 批量删除工艺文件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (processFileInformationService *ProcessFileInformationService) DeleteProcessFileInformationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.ProcessFileInformation{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.ProcessFileInformation{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateProcessFileInformation 更新工艺文件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (processFileInformationService *ProcessFileInformationService) UpdateProcessFileInformation(processFileInformation alpha.ProcessFileInformation) (err error) {
	err = global.GVA_DB.Model(&alpha.ProcessFileInformation{}).Where("id = ?", processFileInformation.ID).Updates(&processFileInformation).Error
	return err
}

// GetProcessFileInformation 根据ID获取工艺文件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (processFileInformationService *ProcessFileInformationService) GetProcessFileInformation(ID string) (processFileInformation alpha.ProcessFileInformation, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&processFileInformation).Error
	return
}

// GetProcessFileInformationInfoList 分页获取工艺文件信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (processFileInformationService *ProcessFileInformationService) GetProcessFileInformationInfoList(info alphaReq.ProcessFileInformationSearch) (list []alpha.ProcessFileInformation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.ProcessFileInformation{})
	var processFileInformations []alpha.ProcessFileInformation
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
	if info.ProcessCardNumber != "" {
		db = db.Where("process_card_number LIKE ?", "%"+info.ProcessCardNumber+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&processFileInformations).Error
	return processFileInformations, total, err
}
