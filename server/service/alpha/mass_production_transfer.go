package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
)

type MassProductionTransferService struct {
}

// CreateMassProductionTransfer 创建量产转移记录
// Author [piexlmax](https://github.com/piexlmax)
func (massProductionTransferService *MassProductionTransferService) CreateMassProductionTransfer(massProductionTransfer *alpha.MassProductionTransfer) (err error) {
	err = global.GVA_DB.Create(massProductionTransfer).Error
	return err
}

// DeleteMassProductionTransfer 删除量产转移记录
// Author [piexlmax](https://github.com/piexlmax)
func (massProductionTransferService *MassProductionTransferService) DeleteMassProductionTransfer(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.MassProductionTransfer{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.MassProductionTransfer{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMassProductionTransferByIds 批量删除量产转移记录
// Author [piexlmax](https://github.com/piexlmax)
func (massProductionTransferService *MassProductionTransferService) DeleteMassProductionTransferByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.MassProductionTransfer{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.MassProductionTransfer{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMassProductionTransfer 更新量产转移记录
// Author [piexlmax](https://github.com/piexlmax)
func (massProductionTransferService *MassProductionTransferService) UpdateMassProductionTransfer(massProductionTransfer alpha.MassProductionTransfer) (err error) {
	err = global.GVA_DB.Model(&alpha.MassProductionTransfer{}).Where("id = ?", massProductionTransfer.ID).Updates(&massProductionTransfer).Error
	return err
}

// GetMassProductionTransfer 根据ID获取量产转移记录
// Author [piexlmax](https://github.com/piexlmax)
func (massProductionTransferService *MassProductionTransferService) GetMassProductionTransfer(ID string) (massProductionTransfer alpha.MassProductionTransfer, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&massProductionTransfer).Error
	return
}

// GetMassProductionTransferInfoList 分页获取量产转移记录
// Author [piexlmax](https://github.com/piexlmax)
func (massProductionTransferService *MassProductionTransferService) GetMassProductionTransferInfoList(info alphaReq.MassProductionTransferSearch) (list []alpha.MassProductionTransfer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.MassProductionTransfer{})
	var massProductionTransfers []alpha.MassProductionTransfer
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UTN != "" {
		db = db.Where("UTN LIKE ?", "%"+info.UTN+"%")
	}
	if info.MB202 != "" {
		db = db.Where("MB202 LIKE ?", "%"+info.MB202+"%")
	}
	if info.SN != "" {
		db = db.Where("SN LIKE ?", "%"+info.SN+"%")
	}
	if info.AuditSN != "" {
		db = db.Where("audit_sn LIKE ?", "%"+info.AuditSN+"%")
	}
	if info.StartTransferDate != nil && info.EndTransferDate != nil {
		db = db.Where("transfer_date BETWEEN ? AND ? ", info.StartTransferDate, info.EndTransferDate)
	}
	if info.AuditStatus != "" {
		db = db.Where("audit_status = ?", info.AuditStatus)
	}
	if info.StartSucessDate != nil && info.EndSucessDate != nil {
		db = db.Where("sucess_date BETWEEN ? AND ? ", info.StartSucessDate, info.EndSucessDate)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&massProductionTransfers).Error
	return massProductionTransfers, total, err
}
