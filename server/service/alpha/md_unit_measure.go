package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
)

type MdUnitMeasureService struct {
}

// CreateMdUnitMeasure 创建品号单位设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdUnitMeasureService *MdUnitMeasureService) CreateMdUnitMeasure(mdUnitMeasure *alpha.MdUnitMeasure) (err error) {
	err = global.GVA_DB.Create(mdUnitMeasure).Error
	return err
}

// DeleteMdUnitMeasure 删除品号单位设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdUnitMeasureService *MdUnitMeasureService) DeleteMdUnitMeasure(ID string) (err error) {
	err = global.GVA_DB.Delete(&alpha.MdUnitMeasure{}, "id = ?", ID).Error
	return err
}

// DeleteMdUnitMeasureByIds 批量删除品号单位设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdUnitMeasureService *MdUnitMeasureService) DeleteMdUnitMeasureByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]alpha.MdUnitMeasure{}, "id in ?", IDs).Error
	return err
}

// UpdateMdUnitMeasure 更新品号单位设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdUnitMeasureService *MdUnitMeasureService) UpdateMdUnitMeasure(mdUnitMeasure alpha.MdUnitMeasure) (err error) {
	err = global.GVA_DB.Save(&mdUnitMeasure).Error
	return err
}

// GetMdUnitMeasure 根据ID获取品号单位设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdUnitMeasureService *MdUnitMeasureService) GetMdUnitMeasure(ID string) (mdUnitMeasure alpha.MdUnitMeasure, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mdUnitMeasure).Error
	return
}

// GetMdUnitMeasureInfoList 分页获取品号单位设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdUnitMeasureService *MdUnitMeasureService) GetMdUnitMeasureInfoList(info alphaReq.MdUnitMeasureSearch) (list []alpha.MdUnitMeasure, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.MdUnitMeasure{})
	var mdUnitMeasures []alpha.MdUnitMeasure
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

	err = db.Find(&mdUnitMeasures).Error
	return mdUnitMeasures, total, err
}
