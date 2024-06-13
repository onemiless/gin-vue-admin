package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
)

type MultiSelectService struct {
}

// CreateMutiSelect 创建多级选择信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mutiSelectService *MultiSelectService) CreateMutiSelect(mutiSelect *alpha.MutiSelect) (err error) {
	err = global.GVA_DB.Create(mutiSelect).Error
	return err
}

// DeleteMutiSelect 删除多级选择信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mutiSelectService *MultiSelectService) DeleteMutiSelect(ID string) (err error) {
	err = global.GVA_DB.Delete(&alpha.MutiSelect{}, "id = ?", ID).Error
	return err
}

// DeleteMutiSelectByIds 批量删除多级选择信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mutiSelectService *MultiSelectService) DeleteMutiSelectByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]alpha.MutiSelect{}, "id in ?", IDs).Error
	return err
}

// UpdateMutiSelect 更新多级选择信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mutiSelectService *MultiSelectService) UpdateMutiSelect(mutiSelect alpha.MutiSelect) (err error) {
	err = global.GVA_DB.Save(&mutiSelect).Error
	return err
}

// GetMutiSelect 根据ID获取多级选择信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mutiSelectService *MultiSelectService) GetMutiSelect(ID string) (mutiSelect alpha.MutiSelect, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mutiSelect).Error
	return
}

// GetMutiSelectInfoList 分页获取多级选择信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mutiSelectService *MultiSelectService) GetMutiSelectInfoList(info alphaReq.MutiSelectSearch) (list []alpha.MutiSelect, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.MutiSelect{})
	var mutiSelects []alpha.MutiSelect
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

	err = db.Find(&mutiSelects).Error
	return mutiSelects, total, err
}
