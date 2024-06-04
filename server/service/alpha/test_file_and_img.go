package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type TestFileAndImgService struct {
}

// CreateTestFileAndImg 创建TestFileAndImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (testFileAndImgService *TestFileAndImgService) CreateTestFileAndImg(testFileAndImg *alpha.TestFileAndImg) (err error) {
	err = global.GVA_DB.Create(testFileAndImg).Error
	return err
}

// DeleteTestFileAndImg 删除TestFileAndImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (testFileAndImgService *TestFileAndImgService) DeleteTestFileAndImg(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TestFileAndImg{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.TestFileAndImg{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTestFileAndImgByIds 批量删除TestFileAndImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (testFileAndImgService *TestFileAndImgService) DeleteTestFileAndImgByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TestFileAndImg{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.TestFileAndImg{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTestFileAndImg 更新TestFileAndImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (testFileAndImgService *TestFileAndImgService) UpdateTestFileAndImg(testFileAndImg alpha.TestFileAndImg) (err error) {
	err = global.GVA_DB.Model(&alpha.TestFileAndImg{}).Where("id = ?", testFileAndImg.ID).Updates(&testFileAndImg).Error
	return err
}

// GetTestFileAndImg 根据ID获取TestFileAndImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (testFileAndImgService *TestFileAndImgService) GetTestFileAndImg(ID string) (testFileAndImg alpha.TestFileAndImg, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&testFileAndImg).Error
	return
}

// GetTestFileAndImgInfoList 分页获取TestFileAndImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (testFileAndImgService *TestFileAndImgService) GetTestFileAndImgInfoList(info alphaReq.TestFileAndImgSearch) (list []alpha.TestFileAndImg, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.TestFileAndImg{})
	var testFileAndImgs []alpha.TestFileAndImg
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

	err = db.Find(&testFileAndImgs).Error
	return testFileAndImgs, total, err
}
