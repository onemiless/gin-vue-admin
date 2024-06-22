package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type PackageRequirementService struct {
}

// CreatePackageRequirement 创建包装要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (packageRequirementService *PackageRequirementService) CreatePackageRequirement(packageRequirement *alpha.PackageRequirement) (err error) {
	err = global.GVA_DB.Create(packageRequirement).Error
	return err
}

// DeletePackageRequirement 删除包装要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (packageRequirementService *PackageRequirementService) DeletePackageRequirement(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.PackageRequirement{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.PackageRequirement{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePackageRequirementByIds 批量删除包装要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (packageRequirementService *PackageRequirementService) DeletePackageRequirementByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.PackageRequirement{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.PackageRequirement{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePackageRequirement 更新包装要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (packageRequirementService *PackageRequirementService) UpdatePackageRequirement(packageRequirement alpha.PackageRequirement) (err error) {
	err = global.GVA_DB.Model(&alpha.PackageRequirement{}).Where("id = ?", packageRequirement.ID).Updates(&packageRequirement).Error
	return err
}

// GetPackageRequirement 根据ID获取包装要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (packageRequirementService *PackageRequirementService) GetPackageRequirement(ID string) (packageRequirement alpha.PackageRequirement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&packageRequirement).Error
	return
}

// GetPackageRequirementInfoList 分页获取包装要求信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (packageRequirementService *PackageRequirementService) GetPackageRequirementInfoList(info alphaReq.PackageRequirementSearch) (list []alpha.PackageRequirement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.PackageRequirement{})
	var packageRequirements []alpha.PackageRequirement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UTN != "" {
		db = db.Where("UTN LIKE ?", "%"+info.UTN+"%")
	}
	if info.SN != "" {
		db = db.Where("SN LIKE ?", "%"+info.SN+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&packageRequirements).Error
	return packageRequirements, total, err
}
