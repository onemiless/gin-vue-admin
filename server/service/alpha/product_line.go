package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type ProductLineService struct{}

// CreateProductLine 创建产线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (productLineService *ProductLineService) CreateProductLine(productLine *alpha.ProductLine) (err error) {
	err = global.GVA_DB.Create(productLine).Error
	return err
}

// DeleteProductLine 删除产线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (productLineService *ProductLineService) DeleteProductLine(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.ProductLine{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.ProductLine{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteProductLineByIds 批量删除产线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (productLineService *ProductLineService) DeleteProductLineByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.ProductLine{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.ProductLine{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateProductLine 更新产线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (productLineService *ProductLineService) UpdateProductLine(productLine alpha.ProductLine) (err error) {
	err = global.GVA_DB.Model(&alpha.ProductLine{}).Where("id = ?", productLine.ID).Updates(&productLine).Error
	return err
}

// GetProductLine 根据ID获取产线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (productLineService *ProductLineService) GetProductLine(ID string) (productLine alpha.ProductLine, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&productLine).Error
	return
}

// GetProductLineInfoList 分页获取产线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (productLineService *ProductLineService) GetProductLineInfoList(info alphaReq.ProductLineSearch) (list []alpha.ProductLine, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.ProductLine{})
	var productLines []alpha.ProductLine
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ProductNumber != "" {
		db = db.Where("product_number LIKE ?", "%"+info.ProductNumber+"%")
	}
	if info.ProductName != "" {
		db = db.Where("product_name LIKE ?", "%"+info.ProductName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&productLines).Error
	return productLines, total, err
}
