package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
)

type CMSMEService struct {
}

// CreateCMSME 创建ERP部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsmeService *CMSMEService) CreateCMSME(cmsme *alpha.CMSME) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Create(cmsme).Error
	return err
}

// DeleteCMSME 删除ERP部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsmeService *CMSMEService) DeleteCMSME(ME001 string) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Delete(&alpha.CMSME{}, "ME001 = ?", ME001).Error
	return err
}

// DeleteCMSMEByIds 批量删除ERP部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsmeService *CMSMEService) DeleteCMSMEByIds(ME001s []string) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Delete(&[]alpha.CMSME{}, "ME001 in ?", ME001s).Error
	return err
}

// UpdateCMSME 更新ERP部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsmeService *CMSMEService) UpdateCMSME(cmsme alpha.CMSME) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Save(&cmsme).Error
	return err
}

// GetCMSME 根据ME001获取ERP部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsmeService *CMSMEService) GetCMSME(ME001 string) (cmsme alpha.CMSME, err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Where("ME001 = ?", ME001).First(&cmsme).Error
	return
}

// GetCMSMEInfoList 分页获取ERP部门记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsmeService *CMSMEService) GetCMSMEInfoList(info alphaReq.CMSMESearch) (list []alpha.CMSME, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("A1111").Model(&alpha.CMSME{})
	var cmsmes []alpha.CMSME
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cmsmes).Error
	return cmsmes, total, err
}
