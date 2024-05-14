package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type ProofingInformationService struct {
}

// CreateProofingInformation 创建打样信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (proofingInformationService *ProofingInformationService) CreateProofingInformation(proofingInformation *alpha.ProofingInformation) (err error) {
	err = global.GVA_DB.Create(proofingInformation).Error
	return err
}

// DeleteProofingInformation 删除打样信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (proofingInformationService *ProofingInformationService) DeleteProofingInformation(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.ProofingInformation{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.ProofingInformation{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteProofingInformationByIds 批量删除打样信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (proofingInformationService *ProofingInformationService) DeleteProofingInformationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.ProofingInformation{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.ProofingInformation{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateProofingInformation 更新打样信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (proofingInformationService *ProofingInformationService) UpdateProofingInformation(proofingInformation alpha.ProofingInformation) (err error) {
	err = global.GVA_DB.Save(&proofingInformation).Error
	return err
}

// GetProofingInformation 根据ID获取打样信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (proofingInformationService *ProofingInformationService) GetProofingInformation(ID string) (proofingInformation alpha.ProofingInformation, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&proofingInformation).Error
	return
}

// GetProofingInformationInfoList 分页获取打样信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (proofingInformationService *ProofingInformationService) GetProofingInformationInfoList(info alphaReq.ProofingInformationSearch) (list []alpha.ProofingInformation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.ProofingInformation{})
	var proofingInformations []alpha.ProofingInformation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&proofingInformations).Error
	return proofingInformations, total, err
}
