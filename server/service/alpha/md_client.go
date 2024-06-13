package alpha

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alpha"
	alphaReq "github.com/flipped-aurora/gin-vue-admin/server/model/alpha/request"
)

type MdClientService struct {
}

// CreateMdClient 创建客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdClientService *MdClientService) CreateMdClient(mdClient *alpha.MdClient) (err error) {
	copma := COPMAService{}
	//获取ERP客户信息
	list, err := copma.GetCOPMAInfoAllList()
	if err != nil {
		return err
	}
	//获取本地客户信息
	infoList, err := mdClientService.GetMdClientAllInfoList()

	//insertlist和updatelist
	//insertList := []alpha.MdClient{}
	//updateList := []alpha.MdClient{}
	//

	//比较infoList和list中的数据，多的添加，少的删除，需要更新的更新
	enableFlag := 1
	fmt.Println(enableFlag)
	for _, v := range list {
		flag := false
		updateFlag := false
		if v.MA097 != "1" {
			enableFlag = 0
		}
		for _, v1 := range infoList {
			if v.MA001 == v1.ClientCode && v.MA002 == v1.ClientName {
				flag = true
				break
			}
			if v.MA001 == v1.ClientCode && (v.MA002 != v1.ClientName || v.MA003 != v1.ClientDes || v.UDF01 != v1.ClientEn || v.MA102 != v1.ClientType) {
				updateFlag = true
				break
			}

		}
		if !flag {
			//添加

			err = global.GVA_DB.Create(&alpha.MdClient{
				ClientCode: v.MA001,
				ClientName: v.MA002,
				ClientType: v.MA102,
				ClientEn:   v.UDF01,
				ClientDes:  v.MA003,
				EnableFlag: &enableFlag,
			}).Error
			if err != nil {
				return err
			}
		}
		if updateFlag {
			//更新
			err = mdClientService.UpdateMdClient(alpha.MdClient{
				ClientCode: v.MA001,
				ClientName: v.MA002,
				ClientType: v.MA102,
				ClientEn:   v.UDF01,
				ClientDes:  v.MA003,
				EnableFlag: &enableFlag,
			})
			if err != nil {
				return err
			}
		}
	}

	//err = global.GVA_DB.Create(mdClient).Error
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
func (mdClientService *MdClientService) GetMdClient(client_code string) (mdClient alpha.MdClient, err error) {
	err = global.GVA_DB.Where("client_code = ?", client_code).First(&mdClient).Error
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
	if info.Search != "" {
		//like
		db = db.Or("client_code LIKE ?", "%"+info.Search+"%").Or("client_name LIKE ?", "%"+info.Search+"%").Or("client_en LIKE ?", "%"+info.Search+"%")

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

// GetMdClientInfoAllList 分页获取客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (mdClientService *MdClientService) GetMdClientAllInfoList() (list []alpha.MdClient, err error) {
	fmt.Println("get all copma info ")
	// 创建db
	db := global.GVA_DB.Model(&alpha.MdClient{})
	var mdClients []alpha.MdClient

	if err != nil {
		return
	}

	err = db.Find(&mdClients).Error
	return mdClients, err
}
