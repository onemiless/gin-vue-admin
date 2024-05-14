package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type TecBaseProcessService struct {
}

// CreateTecBaseProcess 创建工艺、设备及治具基本信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseProcessService *TecBaseProcessService) CreateTecBaseProcess(tecBaseProcess *alpha.TecBaseProcess) (err error) {
	err = global.GVA_DB.Create(tecBaseProcess).Error
	return err
}

// DeleteTecBaseProcess 删除工艺、设备及治具基本信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseProcessService *TecBaseProcessService) DeleteTecBaseProcess(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TecBaseProcess{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.TecBaseProcess{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTecBaseProcessByIds 批量删除工艺、设备及治具基本信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseProcessService *TecBaseProcessService) DeleteTecBaseProcessByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.TecBaseProcess{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.TecBaseProcess{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTecBaseProcess 更新工艺、设备及治具基本信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseProcessService *TecBaseProcessService) UpdateTecBaseProcess(tecBaseProcess alpha.TecBaseProcess) (err error) {
	err = global.GVA_DB.Save(&tecBaseProcess).Error
	return err
}

// GetTecBaseProcess 根据ID获取工艺、设备及治具基本信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseProcessService *TecBaseProcessService) GetTecBaseProcess(ID string) (tecBaseProcess alpha.TecBaseProcess, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&tecBaseProcess).Error
	return
}

// GetTecBaseProcessInfoList 分页获取工艺、设备及治具基本信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (tecBaseProcessService *TecBaseProcessService) GetTecBaseProcessInfoList(info alphaReq.TecBaseProcessSearch) (list []alpha.TecBaseProcess, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.TecBaseProcess{})
	var tecBaseProcesss []alpha.TecBaseProcess
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ParentID > 0 {
		db = db.Where("parent_id = ?", info.ParentID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tecBaseProcesss).Error
	return tecBaseProcesss, total, err
}

func (tecBaseProcessService *TecBaseProcessService) GetTecBaseProcessParentID(parentId string) (tecBaseProcess alpha.TecBaseProcess, err error) {
	err = global.GVA_DB.Where("parent_id = ?", parentId).First(&tecBaseProcess).Error
	return

}
