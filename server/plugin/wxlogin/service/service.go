package service

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	global2 "github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/utils"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service"
	utils2 "github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
	"strings"
	"time"
)

type WxLoginService struct{}

func (e *WxLoginService) Signin(signature, timestamp, nonce, Token string) (err error) {
	ok := utils.CheckSignature(signature, timestamp, nonce, Token)
	if ok {
		return nil
	} else {
		return errors.New("校验失败")
	}
	// 写你的业务逻辑
}

func (e *WxLoginService) PingPang(textMsg model.WXTextMsg, openid string) (msg []byte, err error) {
	switch textMsg.Event {
	case global2.SCAN:
		fmt.Println("用户扫码", textMsg.EventKey)
		global2.LoginCodeMap.Store(textMsg.EventKey, global2.CodeInfo{
			CanLogin:   true,
			OpenId:     openid,
			CreateTime: time.Now().Unix(),
		})
	case global2.UNSUBSCRIBE:
		fmt.Println("用户取消关注")
	case global2.SUBSCRIBE:
		if strings.Index(textMsg.EventKey, "_") > -1 {
			arg := strings.Split(textMsg.EventKey, "_")
			if arg[0] == "qrscene" {
				global2.LoginCodeMap.Store(arg[1], global2.CodeInfo{
					CanLogin:   true,
					OpenId:     openid,
					CreateTime: time.Now().Unix(),
				})
			}
		}
	}
	return msg, err
}

func (e *WxLoginService) LoginOrCreate(user model.WXUserInfo, id uint) (*system.SysUser, error) {
	if id != 0 {
		//绑定逻辑
		err := global.GVA_DB.First(&model.WXUserInfo{}, "openid = ?", user.Openid).Error
		if err == nil {
			return nil, errors.New("此微信已绑定其他账号")
		}
		user.GvaUserId = id
		err = global.GVA_DB.Create(user).Error
		if err != nil {
			return nil, errors.New("此账号已绑定其他微信")
		}
		return nil, nil
	} else {
		//	扫码登录
		var loginUser model.WXUserInfo
		err := global.GVA_DB.First(&loginUser, "openid = ?", user.Openid).Error
		if err != nil {
			return nil, err
		}
		var resUser system.SysUser
		err = global.GVA_DB.First(&resUser, "id = ?", loginUser.GvaUserId).Error
		if err != nil {
			return nil, err
		}
		return &resUser, nil
	}
}

func (e *WxLoginService) ClearWx(ID uint) error {
	return global.GVA_DB.Delete(&model.WXUserInfo{}, "gva_user_id = ?", ID).Error
}

var userService = systemService.ServiceGroupApp.SystemServiceGroup.UserService

func (e *WxLoginService) Register(user model.Register, wxUser model.WXUserInfo) (*system.SysUser, error) {
	var regUser system.SysUser
	err := global.GVA_DB.First(&regUser, "username = ?", user.Username).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//	注册
			regUser = system.SysUser{
				Username:    user.Username,
				Password:    user.Password,
				NickName:    "微信注册用户",
				AuthorityId: global2.WxGlobal.AuthorityID,
				Authorities: []system.SysAuthority{
					{
						AuthorityId: global2.WxGlobal.AuthorityID,
					},
				},
			}
			gvaInfo, err := userService.Register(regUser)
			if err != nil {
				return nil, err
			}
			wxUser.GvaUserId = gvaInfo.ID
			err = global.GVA_DB.Create(wxUser).Error
			if err != nil {
				return nil, err
			}
			return &gvaInfo, nil
		}
		return nil, err
	} else {
		if ok := utils2.BcryptCheck(user.Password, regUser.Password); !ok {
			return nil, errors.New("密码错误")
		}
		wxUser.GvaUserId = regUser.ID
		err = global.GVA_DB.Create(wxUser).Error
		if err != nil {
			return nil, err
		}
	}
	return &regUser, nil
}
