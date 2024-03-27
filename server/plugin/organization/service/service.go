package service

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	organization "github.com/flipped-aurora/gin-vue-admin/server/plugin/organization/model"
	organizationReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/organization/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/organization/utils"
	utils2 "github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type OrganizationService struct {
}

// CreateOrganization 创建Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (orgService *OrganizationService) CreateOrganization(org organization.Organization) (err error) {
	err = global.GVA_DB.Create(&org).Error
	return err
}

// DeleteOrganization 删除Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (orgService *OrganizationService) DeleteOrganization(org organization.Organization) (err error) {
	err = global.GVA_DB.Where("parent_id = ?", org.ID).First(&organization.Organization{}).Error
	if err == nil {
		return errors.New("该组织有子组织，不能删除")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Delete(&org, "id = ?", org.ID).Error
	}
	return err
}

// DeleteOrganizationByIds 批量删除Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (orgService *OrganizationService) DeleteOrganizationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]organization.Organization{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrganization 更新Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (orgService *OrganizationService) UpdateOrganization(org organization.Organization) (err error) {
	var updatesmap = make(map[string]interface{})
	updatesmap["Name"] = org.Name
	updatesmap["ParentID"] = org.ParentID
	err = global.GVA_DB.Model(&organization.Organization{}).Where("id = ?", org.ID).Updates(updatesmap).Error
	return err
}

// GetOrganization 根据id获取Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (orgService *OrganizationService) GetOrganization(id uint) (org organization.Organization, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&org).Error
	return
}

// GetOrganizationInfoList 分页获取Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (orgService *OrganizationService) GetOrganizationInfoList(info organizationReq.OrganizationSearch) (list interface{}, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&organization.Organization{})
	var orgs []organization.Organization
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("parent_id = ?", info.ParentID)
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&orgs).Error
	return orgs, total, err
}

func (orgService *OrganizationService) CreateOrgUser(orgUser organization.OrgUserReq) error {
	var Users []organization.OrgUser
	var CUsers []organization.OrgUser
	err := global.GVA_DB.Find(&Users, "organization_id = ?", orgUser.OrganizationID).Error
	if err != nil {
		return err
	}
	var UserIdMap = make(map[uint]bool)
	for i := range Users {
		UserIdMap[Users[i].SysUserID] = true
	}

	for i := range orgUser.SysUserIDS {
		if !UserIdMap[orgUser.SysUserIDS[i]] {
			CUsers = append(CUsers, organization.OrgUser{SysUserID: orgUser.SysUserIDS[i], OrganizationID: orgUser.OrganizationID})
		}
	}
	err = global.GVA_DB.Create(&CUsers).Error
	return err
}

func (orgService *OrganizationService) FindOrgUserAll(orgID string) ([]uint, error) {
	var Users []organization.OrgUser
	var ids []uint
	err := global.GVA_DB.Find(&Users, "organization_id = ?", orgID).Error
	if err != nil {
		return ids, err
	}
	for i := range Users {
		ids = append(ids, Users[i].SysUserID)
	}
	return ids, err
}

// GetOrganizationInfoList 分页获取当前组织下用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (orgService *OrganizationService) GetOrgUserList(info organizationReq.OrgUserSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&organization.OrgUser{}).Joins("SysUser").Preload("SysUser.Authority")
	var orgs []organization.OrgUser
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("organization_id = ?", info.OrganizationID)
	if info.UserName != "" {
		db = db.Where("SysUser.nick_name LIKE ?", "%"+info.UserName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&orgs).Error
	return orgs, total, err
}

func (orgService *OrganizationService) SetOrgUserAdmin(id uint, flag bool) (err error) {
	return global.GVA_DB.Model(&organization.OrgUser{}).Where("sys_user_id = ?", id).Update("is_admin", flag).Error
}

func (orgService *OrganizationService) SetOrgAuthority(authID uint, authorityType int) (err error) {
	return global.GVA_DB.Model(&organization.DataAuthority{}).Where("authority_id = ?", authID).Update("authority_type", authorityType).Error
}

func (orgService *OrganizationService) GetOrgAuthority() (authorityData []organization.DataAuthority, err error) {
	err = global.GVA_DB.Preload("Authority").Find(&authorityData).Error
	return authorityData, err
}

func (orgService *OrganizationService) SyncAuthority() (err error) {
	// 定义两个切片，分别存储系统权限和组织权限的数据
	var authData []system.SysAuthority
	var auth []organization.DataAuthority

	// 在数据库事务中执行操作
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 用于存储系统权限ID及其状态的映射
		var idMap = make(map[uint]*bool)

		// 查询系统权限数据并处理结果
		err := tx.Find(&authData).Error
		if err != nil {
			return err
		}
		for _, datum := range authData {
			// 将系统权限ID添加到映射中，并设置对应的状态为true
			idMap[datum.AuthorityId] = utils.GetBoolPointer(true)
		}

		// 查询组织权限数据并处理结果
		err = tx.Find(&auth).Error
		if err != nil {
			return err
		}
		for _, datum := range auth {
			// 根据组织权限ID更新映射中对应的状态
			if idMap[datum.AuthorityID] != nil {
				idMap[datum.AuthorityID] = utils.GetBoolPointer(false)
			} else {
				idMap[datum.AuthorityID] = nil
			}
		}

		// 用于存储要异步添加的组织权限数据和要删除的组织权限数据
		var ayncData []organization.DataAuthority
		var deleteAuth []organization.DataAuthority

		// 根据映射中的状态，生成异步添加和删除的组织权限数据
		for k, _ := range idMap {
			if idMap[k] != nil && *idMap[k] {
				ayncData = append(ayncData, organization.DataAuthority{
					AuthorityID:   k,
					AuthorityType: 0,
				})
			}
			if idMap[k] == nil {
				deleteAuth = append(deleteAuth, organization.DataAuthority{
					AuthorityID:   k,
					AuthorityType: 0,
				})
			}
		}

		// 如果有需要同步的数据，则执行同步操作
		if len(ayncData) > 0 || len(deleteAuth) > 0 {
			// 异步添加组织权限数据
			if len(ayncData) > 0 {
				err := tx.Create(&ayncData).Error
				if err != nil {
					return err
				}
			}

			// 删除组织权限数据
			if len(deleteAuth) > 0 {
				var deleteAuthIds []uint
				for i := range deleteAuth {
					deleteAuthIds = append(deleteAuthIds, deleteAuth[i].AuthorityID)
				}
				err = tx.Delete(&deleteAuth, "authority_id in (?)", deleteAuthIds).Error
				if err != nil {
					return err
				}
			}
			return nil
		} else {
			// 如果无需同步则返回错误信息
			return errors.New("当前无需同步")
		}
	})
}

func (orgService *OrganizationService) DeleteOrgUser(ids []uint, orgID uint) (err error) {
	return global.GVA_DB.Where("sys_user_id in (?) and organization_id = ?", ids, orgID).Delete(&[]organization.OrgUser{}).Error
}

func (orgService *OrganizationService) TransferOrgUser(ids []uint, orgID, toOrgID uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var CUsers []organization.OrgUser
		err := global.GVA_DB.Where("sys_user_id in (?) and organization_id in (?)", ids, []uint{orgID, toOrgID}).Delete(&[]organization.OrgUser{}).Error
		if err != nil {
			return err
		}
		for i := range ids {
			CUsers = append(CUsers, organization.OrgUser{SysUserID: ids[i], OrganizationID: toOrgID})
		}
		err = global.GVA_DB.Create(&CUsers).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (orgService *OrganizationService) SyncWecomOrganization(orgs []organization.OrganizationCus) (err error) {
	var dborgs []organization.OrganizationCus
	err = global.GVA_DB.Find(&dborgs).Error
	if err != nil {
		return err
	}
	var updateOrgs []organization.OrganizationCus
	//var orgMap = make(map[uint]organization.Organization)
	//对比orgs和dborgs 找出dborgs中没有的且不同的部分
	diff := utils2.GenericSliceDiff(orgs, dborgs)
	//获取dborgs中没有部分 并添加到数据库
	//完全一致不需要更新和新增
	if diff == nil {
		//s := err.Error()
		return
	}
	addOrgs := utils2.GenericSliceDiff(dborgs, diff)
	//获取需要更新的部分
	for _, org := range orgs {
		if org.ParentID == 0 {
			continue
		}
		for _, dborg := range dborgs {
			if org.ID == dborg.ID && (dborg.Name != org.Name || dborg.ParentID != org.ParentID) {
				updateOrgs = append(updateOrgs, org)
			}
		}
	}
	//添加addOrgs到数据库
	err = global.GVA_DB.Create(&addOrgs).Error
	//更新updateOrgs到数据库
	for _, org := range updateOrgs {
		err = global.GVA_DB.Where("id = ?", org.ID).Updates(org).Error
		if err != nil {
			return err
		}
	}
	//同步人员和部门关系

	//var sysUsers []model.SysUser
	//err = global.GVA_DB.Find(&sysUsers).Error
	//if err != nil {
	//	return err
	//}
	//var orgUsers []organization.OrgUser
	//err = global.GVA_DB.Find(&orgUsers).Error
	//if err != nil {
	//	return err
	//}
	//sysUserMap := make(map[uint]bool)
	//for i := range orgUsers {
	//	sysUserMap[orgUsers[i].SysUserID] = true
	//}
	//var addOrgUsers []organization.OrgUser
	//for _, user := range sysUsers {
	//	if sysUserMap[user.ID] {
	//		continue
	//	}
	//	addOrgUsers = append(addOrgUsers, organization.OrgUser{SysUserID: user.ID, OrganizationID: user.DeptId, IsAdmin: false})
	//}
	//err = global.GVA_DB.Create(&addOrgUsers).Error
	////更新部门下的人员
	//var updateOrgUsers []organization.OrgUser
	//for _, user := range sysUsers {
	//	for _, orgUser := range orgUsers {
	//		if user.ID == orgUser.SysUserID && user.DeptId != orgUser.OrganizationID {
	//			orgUser.OrganizationID = user.DeptId
	//			updateOrgUsers = append(updateOrgUsers, orgUser)
	//		}
	//	}
	//}
	//for _, orgUser := range updateOrgUsers {
	//	err = global.GVA_DB.Where("sys_user_id = ?", orgUser.SysUserID).Updates(orgUser).Error
	//	if err != nil {
	//		return err
	//	}
	//}

	//diff := utils2.GenericSliceDiff(orgs, dborgs)
	//utils2.
	return err
}

func (orgService *OrganizationService) SyncWecomUser() (err error) {
	var sysUsers []model.SysUser
	err = global.GVA_DB.Find(&sysUsers).Error
	if err != nil {
		return err
	}
	var orgUsers []organization.OrgUser
	err = global.GVA_DB.Find(&orgUsers).Error
	if err != nil {
		return err
	}
	sysUserMap := make(map[uint]bool)
	for i := range orgUsers {
		sysUserMap[orgUsers[i].SysUserID] = true
	}
	var addOrgUsers []organization.OrgUser
	for _, user := range sysUsers {
		if sysUserMap[user.ID] {
			continue
		}
		addOrgUsers = append(addOrgUsers, organization.OrgUser{SysUserID: user.ID, OrganizationID: user.DeptId, IsAdmin: false})
	}
	err = global.GVA_DB.Create(&addOrgUsers).Error
	//更新部门下的人员
	var updateOrgUsers []organization.OrgUser
	for _, user := range sysUsers {
		for _, orgUser := range orgUsers {
			if user.ID == orgUser.SysUserID && user.DeptId != orgUser.OrganizationID {
				orgUser.OrganizationID = user.DeptId
				updateOrgUsers = append(updateOrgUsers, orgUser)
			}
		}
	}
	for _, orgUser := range updateOrgUsers {
		err = global.GVA_DB.Where("sys_user_id = ?", orgUser.SysUserID).Updates(orgUser).Error
		if err != nil {
			return err
		}
	}
	return err
}
