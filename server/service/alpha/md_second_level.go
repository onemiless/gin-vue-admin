package alpha

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type MdSecondLevelService struct {
}

// CreateMdSecondLevel 创建层级二记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdSecondLevelService *MdSecondLevelService) CreateMdSecondLevel(mdSecondLevel *alpha.MdSecondLevel) (err error) {
	err = global.GVA_DB.Create(mdSecondLevel).Error
	return err
}

// DeleteMdSecondLevel 删除层级二记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdSecondLevelService *MdSecondLevelService) DeleteMdSecondLevel(ID string, userID uint) (err error) {
	//检查是否有下级记录，有则无法删除
	checkErr := mdSecondLevelService.checkHasThirdLevel([]string{ID})
	if checkErr != nil {
		return checkErr
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.MdSecondLevel{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.MdSecondLevel{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
func (mdSecondLevelService *MdSecondLevelService) checkHasThirdLevel(IDs any) (err error) {
	var thirdLevels []alpha.MdThirdLevel

	err = global.GVA_DB.Model(&thirdLevels).Where("secondLevel_id in (?)", IDs).Find(&thirdLevels).Error
	if err != nil {
		return err
	} else if len(thirdLevels) > 0 {
		//返回有下级记录，无法删除
		return errors.New("存在下级记录，无法删除")
	}
	return
}

// DeleteMdSecondLevelByIds 批量删除层级二记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdSecondLevelService *MdSecondLevelService) DeleteMdSecondLevelByIds(IDs []string, deleted_by uint) (err error) {

	//检查是否有下级记录，有则无法删除
	checkErr := mdSecondLevelService.checkHasThirdLevel(IDs)
	if checkErr != nil {
		return checkErr
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.MdSecondLevel{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.MdSecondLevel{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMdSecondLevel 更新层级二记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdSecondLevelService *MdSecondLevelService) UpdateMdSecondLevel(mdSecondLevel alpha.MdSecondLevel) (err error) {
	err = global.GVA_DB.Save(&mdSecondLevel).Error
	return err
}

// GetMdSecondLevel 根据ID获取层级二记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdSecondLevelService *MdSecondLevelService) GetMdSecondLevel(ID string) (mdSecondLevel alpha.MdSecondLevel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mdSecondLevel).Error
	return
}

// GetMdSecondLevelInfoList 分页获取层级二记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdSecondLevelService *MdSecondLevelService) GetMdSecondLevelInfoList(info alphaReq.MdSecondLevelSearch) (list []alpha.MdSecondLevel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.MdSecondLevel{})
	var mdSecondLevels []alpha.MdSecondLevel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FirstLevelId != 0 {
		db = db.Where("firstLevel_id = ?", info.FirstLevelId)
	}
	if info.Name != "" {
		db = db.Or("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Remark != "" {
		db = db.Or("remark LIKE ?", "%"+info.Remark+"%")
	}
	if info.Query != "" {
		db = db.Where("name LIKE ?", "%"+info.Query+"%")

	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&mdSecondLevels).Error
	return mdSecondLevels, total, err
}
