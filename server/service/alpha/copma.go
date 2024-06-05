package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
)

type COPMAService struct {
}

// CreateCOPMA 创建ERP客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (copmaService *COPMAService) CreateCOPMA(copma *alpha.COPMA) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Create(copma).Error
	return err
}

// DeleteCOPMA 删除ERP客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (copmaService *COPMAService) DeleteCOPMA(MA001 string) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Delete(&alpha.COPMA{}, "MA001 = ?", MA001).Error
	return err
}

// DeleteCOPMAByIds 批量删除ERP客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (copmaService *COPMAService) DeleteCOPMAByIds(MA001s []string) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Delete(&[]alpha.COPMA{}, "MA001 in ?", MA001s).Error
	return err
}

// UpdateCOPMA 更新ERP客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (copmaService *COPMAService) UpdateCOPMA(copma alpha.COPMA) (err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Model(&alpha.COPMA{}).Where("MA001 = ?", copma.MA001).Updates(&copma).Error
	return err
}

// GetCOPMA 根据MA001获取ERP客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (copmaService *COPMAService) GetCOPMA(MA001 string) (copma alpha.COPMA, err error) {
	err = global.MustGetGlobalDBByDBName("A1111").Where("MA001 = ?", MA001).First(&copma).Error
	return
}

// GetCOPMAInfoList 分页获取ERP客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (copmaService *COPMAService) GetCOPMAInfoList(info alphaReq.COPMASearch) (list []alpha.COPMA, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("A1111").Model(&alpha.COPMA{})
	var copmas []alpha.COPMA
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&copmas).Error
	return copmas, total, err
}

// GetCOPMAInfoAllList 获取ERP客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (copmaService *COPMAService) GetCOPMAInfoAllList() (list []alpha.COPMA, err error) {

	// 创建db
	db := global.MustGetGlobalDBByDBName("A1111").Model(&alpha.COPMA{})
	var copmas []alpha.COPMA
	err = db.Find(&copmas).Error
	return copmas, err
}
