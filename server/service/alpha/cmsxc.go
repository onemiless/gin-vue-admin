package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
)

type CMSXCService struct {
}

// CreateCMSXC 创建车厂规范信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsxcService *CMSXCService) CreateCMSXC(cmsxc *alpha.CMSXC) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Create(cmsxc).Error
	return err
}

// DeleteCMSXC 删除车厂规范信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsxcService *CMSXCService) DeleteCMSXC(XC001 string) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Delete(&alpha.CMSXC{}, "XC001 = ?", XC001).Error
	return err
}

// DeleteCMSXCByIds 批量删除车厂规范信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsxcService *CMSXCService) DeleteCMSXCByIds(XC001s []string) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Delete(&[]alpha.CMSXC{}, "XC001 in ?", XC001s).Error
	return err
}

// UpdateCMSXC 更新车厂规范信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsxcService *CMSXCService) UpdateCMSXC(cmsxc alpha.CMSXC) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Save(&cmsxc).Error
	return err
}

// GetCMSXC 根据XC001获取车厂规范信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsxcService *CMSXCService) GetCMSXC(XC001 string) (cmsxc alpha.CMSXC, err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Where("XC001 = ?", XC001).First(&cmsxc).Error
	return
}

// GetCMSXCInfoList 分页获取车厂规范信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsxcService *CMSXCService) GetCMSXCInfoList(info alphaReq.CMSXCSearch) (list []alpha.CMSXC, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("A1111").Model(&alpha.CMSXC{})
	var cmsxcs []alpha.CMSXC
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Query != "" {
		db = db.Where("XC001 LIKE ?", "%"+info.Query+"%").Or("XC002 LIKE ?", "%"+info.Query+"%")
		//db = db.Where()
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cmsxcs).Error
	return cmsxcs, total, err
}
