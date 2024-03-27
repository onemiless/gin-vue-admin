package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
)

type MdClientService struct {
}

// CreateMdClient 创建客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdClientService *MdClientService) CreateMdClient(mdClient *alpha.MdClient) (err error) {
	err = global.GVA_DB.Create(mdClient).Error
	return err
}

// DeleteMdClient 删除客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdClientService *MdClientService) DeleteMdClient(ID string) (err error) {
	err = global.GVA_DB.Delete(&alpha.MdClient{}, "id = ?", ID).Error
	return err
}

// DeleteMdClientByIds 批量删除客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdClientService *MdClientService) DeleteMdClientByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]alpha.MdClient{}, "id in ?", IDs).Error
	return err
}

// UpdateMdClient 更新客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdClientService *MdClientService) UpdateMdClient(mdClient alpha.MdClient) (err error) {
	err = global.GVA_DB.Save(&mdClient).Error
	return err
}

// GetMdClient 根据ID获取客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdClientService *MdClientService) GetMdClient(ID string) (mdClient alpha.MdClient, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mdClient).Error
	return
}

// GetMdClientInfoList 分页获取客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdClientService *MdClientService) GetMdClientInfoList(info alphaReq.MdClientSearch) (list []alpha.MdClient, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.MdClient{})
	var mdClients []alpha.MdClient
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

	err = db.Find(&mdClients).Error
	return mdClients, total, err
}
