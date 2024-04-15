package alpha

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type MdFirstLevelService struct {
}

// CreateMdFirstLevel 创建层级一记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdFirstLevelService *MdFirstLevelService) CreateMdFirstLevel(mdFirstLevel *alpha.MdFirstLevel) (err error) {
	err = global.GVA_DB.Create(mdFirstLevel).Error
	return err
}

func (mdFirstLevelService *MdFirstLevelService) checkHasSecondLevel(IDs any) (err error) {
	var secondLevels []alpha.MdSecondLevel

	err = global.GVA_DB.Model(&secondLevels).Where("firstLevel_id in (?)", IDs).Find(&secondLevels).Error
	if err != nil {
		return err
	} else if len(secondLevels) > 0 {
		//返回有下级记录，无法删除
		return errors.New("存在下级记录，无法删除")
	}
	return
}

// DeleteMdFirstLevel 删除层级一记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdFirstLevelService *MdFirstLevelService) DeleteMdFirstLevel(ID string, userID uint) (err error) {
	// 检查是否有下级记录
	err = mdFirstLevelService.checkHasSecondLevel(ID)
	if err != nil {
		return err
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.MdFirstLevel{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.MdFirstLevel{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMdFirstLevelByIds 批量删除层级一记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdFirstLevelService *MdFirstLevelService) DeleteMdFirstLevelByIds(IDs []string, deleted_by uint) (err error) {
	// 检查是否有下级记录
	err = mdFirstLevelService.checkHasSecondLevel(IDs)
	if err != nil {
		return err
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.MdFirstLevel{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.MdFirstLevel{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMdFirstLevel 更新层级一记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdFirstLevelService *MdFirstLevelService) UpdateMdFirstLevel(mdFirstLevel alpha.MdFirstLevel) (err error) {
	err = global.GVA_DB.Save(&mdFirstLevel).Error
	return err
}

// GetMdFirstLevel 根据ID获取层级一记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdFirstLevelService *MdFirstLevelService) GetMdFirstLevel(ID string) (mdFirstLevel alpha.MdFirstLevel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mdFirstLevel).Error
	return
}

// GetMdFirstLevelInfoList 分页获取层级一记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdFirstLevelService *MdFirstLevelService) GetMdFirstLevelInfoList(info alphaReq.MdFirstLevelSearch) (list []alpha.MdFirstLevel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.MdFirstLevel{})
	var mdFirstLevels []alpha.MdFirstLevel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}
	if info.IsEnable != "" {
		db = db.Where("is_enable = ?", info.IsEnable)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mdFirstLevels).Error
	return mdFirstLevels, total, err
}
