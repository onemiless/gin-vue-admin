package api

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	global2 "github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type WxLoginApi struct{}

// @Tags WxLogin
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxLogin/pingpang [get]
func (p *WxLoginApi) Signin(c *gin.Context) {
	// 与填写的服务器配置中的Token一致
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	if err := service.ServiceGroupApp.Signin(signature, timestamp, nonce, global2.WxGlobal.Token); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		c.String(200, echostr)
	}
}

// @Tags WxLogin
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxLogin/pingpang [post]
func (p *WxLoginApi) PingPang(c *gin.Context) {
	// 与填写的服务器配置中的Token一致
	var textMsg model.WXTextMsg
	openid := c.Query("openid")
	err := c.ShouldBindXML(&textMsg)
	fmt.Println(err, textMsg)
	//
	if _, err := service.ServiceGroupApp.PingPang(textMsg, openid); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		//c.XML(200, msg)
	}
}

// @Tags WxLogin
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxLogin/getLoginPic [get]
func (p *WxLoginApi) GetLoginPic(c *gin.Context) {
	loginFlag := c.Query("loginFlag")
	pic, err := wxLoginPassport.GetLoginPic(loginFlag)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithData(pic, c)
	}
}

// @Tags WxLogin
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxLogin/loginOrCreate [get]
func (p *WxLoginApi) LoginOrCreate(c *gin.Context) {
	loginFlag := c.Query("loginFlag")
	var ID uint
	mapInfo, ok := global2.LoginCodeMap.Load(loginFlag)
	if !ok {
		response.FailWithMessage("请重新获取二维码", c)
		return
	}
	if mapInfo.(global2.CodeInfo).CanLogin {
		user, err := wxLoginPassport.GetUserInfo(mapInfo.(global2.CodeInfo).OpenId)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		token, _ := c.GetQuery("state")
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, _ := j.ParseToken(token)
		if claims == nil {
			ID = 0
		} else {
			ID = claims.BaseClaims.ID
		}
		if resUser, err := service.ServiceGroupApp.LoginOrCreate(*user, ID); err != nil {
			global.GVA_LOG.Error("失败!", zap.Error(err))
			if errors.Is(gorm.ErrRecordNotFound, err) {
				response.OkWithDetailed(gin.H{"scan": true, "needRegister": true}, "登录或绑定微信", c)
			} else {
				response.FailWithMessage(err.Error(), c)
			}
			return
		} else {
			if ID != 0 {
				response.OkWithDetailed(gin.H{"scan": true}, "绑定成功", c)
			} else {
				var baseApi = new(system.BaseApi)
				baseApi.TokenNext(c, *resUser)
			}
			return
		}
	} else {
		response.OkWithDetailed(gin.H{"scan": false}, "未扫码", c)
		return
	}
}

// @Tags WxLogin
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxLogin/clearWx [post]
func (p *WxLoginApi) ClearWx(c *gin.Context) {
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}
	if err := service.ServiceGroupApp.ClearWx(id); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithMessage("解绑成功", c)
	}
}

// @Tags WxLogin
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /wxLogin/register [post]
func (p *WxLoginApi) Register(c *gin.Context) {
	var registReq model.Register
	_ = c.ShouldBindJSON(&registReq)
	mapInfo, ok := global2.LoginCodeMap.Load(registReq.LoginFlag)
	if !ok {
		response.FailWithMessage("请重新获取二维码", c)
		return
	}
	user, err := wxLoginPassport.GetUserInfo(mapInfo.(global2.CodeInfo).OpenId)
	if err != nil {
		response.FailWithMessage("获取微信用户信息失败", c)
	}
	if loginUser, err := service.ServiceGroupApp.Register(registReq, *user); err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		var baseApi = new(system.BaseApi)
		baseApi.TokenNext(c, *loginUser)
	}
}
