package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
)

type ItemTypeService struct {
}

// CreateItemType 创建零件类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemtypeService *ItemTypeService) CreateItemType(itemtype *alpha.ItemType) (err error) {
	err = global.GVA_DB.Create(itemtype).Error
	return err
}

// DeleteItemType 删除零件类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemtypeService *ItemTypeService) DeleteItemType(ID string) (err error) {
	err = global.GVA_DB.Delete(&alpha.ItemType{}, "id = ?", ID).Error
	return err
}

// DeleteItemTypeByIds 批量删除零件类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemtypeService *ItemTypeService) DeleteItemTypeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]alpha.ItemType{}, "id in ?", IDs).Error
	return err
}

// UpdateItemType 更新零件类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemtypeService *ItemTypeService) UpdateItemType(itemtype alpha.ItemType) (err error) {
	err = global.GVA_DB.Save(&itemtype).Error
	return err
}

// GetItemType 根据ID获取零件类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemtypeService *ItemTypeService) GetItemType(ID string) (itemtype alpha.ItemType, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&itemtype).Error
	return
}

// GetItemTypeInfoList 分页获取零件类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemtypeService *ItemTypeService) GetItemTypeInfoList(info alphaReq.ItemTypeSearch) (list []alpha.ItemType, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.ItemType{})
	var itemtypes []alpha.ItemType
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

	err = db.Find(&itemtypes).Error
	return itemtypes, total, err
}

// GetItemTypeInfoList 分页获取零件类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (itemtypeService *ItemTypeService) GetItemTypeInfoParentList(info alphaReq.ItemTypeSearch) (list []alpha.ItemType, total int64, err error) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.ItemType{})
	var itemtypes []alpha.ItemType
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TypeParent != nil {
		db = db.Where("type_parent=?", info.TypeParent)
	} else {
		//if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("type_parent=?", 0)
	}
	//}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	//if limit != 0 {
	//	db = db.Limit(limit).Offset(offset)
	//}

	err = db.Find(&itemtypes).Error
	return itemtypes, total, err
}
