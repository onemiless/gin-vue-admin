package wxlogin

import (
	"fmt"
	global2 "github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/passport"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/router"
	"github.com/gin-gonic/gin"
	"time"
)

var wxLoginPassport = new(passport.WxLoginPassport)

type WxLoginPlugin struct {
}

func CreateWxLoginPlug(AppID, Appsecret, Token string, AuthorityID uint) *WxLoginPlugin {
	global.WxGlobal.AppID = AppID
	global.WxGlobal.Appsecret = Appsecret
	global.WxGlobal.Token = Token
	global.WxGlobal.AuthorityID = AuthorityID

	global2.GVA_DB.AutoMigrate(model.WXUserInfo{})
	if global.WxGlobal.AccessToken == "" {
		go func() {
			err := wxLoginPassport.ReflashAccessToken()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	go func() {
		for {
			time.Sleep(300 * time.Second)
			global.LoginCodeMap.Range(func(key, value interface{}) bool {
				if time.Now().Unix()-value.(global.CodeInfo).CreateTime > 300 {
					global.LoginCodeMap.Delete(key)
				}
				return false
			})
		}
	}()
	return &WxLoginPlugin{}
}

func (*WxLoginPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitWxLoginRouter(group)
}

func (*WxLoginPlugin) RouterPath() string {
	return "wxLogin"
}
