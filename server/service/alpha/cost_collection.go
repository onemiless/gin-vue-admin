package alpha

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type CostCollectionService struct {
}

// CreateCostCollection 创建成本信息收集信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (costCollectionService *CostCollectionService) CreateCostCollection(costCollection *alpha.CostCollection) (err error) {
	err = global.GVA_DB.Create(costCollection).Error
	return err
}

// DeleteCostCollection 删除成本信息收集信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (costCollectionService *CostCollectionService) DeleteCostCollection(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.CostCollection{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alpha.CostCollection{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCostCollectionByIds 批量删除成本信息收集信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (costCollectionService *CostCollectionService) DeleteCostCollectionByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alpha.CostCollection{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alpha.CostCollection{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCostCollection 更新成本信息收集信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (costCollectionService *CostCollectionService) UpdateCostCollection(costCollection alpha.CostCollection) (err error) {
	err = global.GVA_DB.Save(&costCollection).Error
	return err
}

// GetCostCollection 根据ID获取成本信息收集信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (costCollectionService *CostCollectionService) GetCostCollection(ID string) (costCollection alpha.CostCollection, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&costCollection).Error
	return
}

// GetCostCollectionInfoList 分页获取成本信息收集信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (costCollectionService *CostCollectionService) GetCostCollectionInfoList(info alphaReq.CostCollectionSearch) (list []alpha.CostCollection, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alpha.CostCollection{})
	var costCollections []alpha.CostCollection
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
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&costCollections).Error
	return costCollections, total, err
}

func (costCollectionService *CostCollectionService) GetCostCollectionNumber() (string, error) {
	var costCollection alpha.CostCollection
	//获取当天的日期
	date := time.Now().Format("2006-01-02")
	//获取当天的最后一条记录
	err := global.GVA_DB.Where("created_at BETWEEN ? AND ?", date+" 00:00:00", date+" 23:59:59").Last(&costCollection).Error
	//流水号转为数字后加1，再转换为流水号
	if err == nil {
		//字符串转数字
		number, err := strconv.Atoi(costCollection.SN)
		if err != nil {
			return "001", err
		}
		number = number + 1
		//数字转字符串,并填充流水号格式为001
		number = number + 1000
		number = number % 10000
		//截取后三位
		costCollection.SN = strconv.Itoa(number)[1:]

	} else {
		costCollection.SN = "001"
		err = nil
	}
	return costCollection.SN, err

}
