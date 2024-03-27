package passport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wxlogin/model"
	"io"
	"net/http"
	"time"
)

type WxLoginPassport struct {
}

func (w *WxLoginPassport) ReflashAccessToken() error {
	var res *http.Response
	var req *http.Request
	var err error
	for {
		wUrl := fmt.Sprintf(
			"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
			global.WxGlobal.AppID,
			global.WxGlobal.Appsecret,
		)
		client := http.Client{}
		req, err = http.NewRequest("GET", wUrl, nil)
		if err != nil {
			break
		}
		res, err = client.Do(req)
		if err != nil {
			break
		}
		reqBody, err := io.ReadAll(res.Body)
		if err != nil {
			break
		}
		var accessTokenRes model.AccessTokenRes
		json.Unmarshal(reqBody, &accessTokenRes)
		global.WxGlobal.AccessToken = accessTokenRes.AccessToken
		fmt.Println(global.WxGlobal.AccessToken)
		time.Sleep(time.Second * 7200)
	}
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}
	return err
}

func (w *WxLoginPassport) GetLoginPic(loginFlag string) (loginPicRes *model.LoginPicRes, err error) {
	if global.WxGlobal.AccessToken == "" {
		go func() {
			err := w.ReflashAccessToken()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wUrl := fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s",
		global.WxGlobal.AccessToken,
	)

	var reqBody = new(model.PicCreate)
	reqBody.ActionInfo.Scene.SceneStr = loginFlag
	reqBody.ActionName = "QR_STR_SCENE"
	reqBody.ExpireSeconds = 60 * 5
	reqBodyByte, err := json.Marshal(reqBody)
	reader := bytes.NewReader(reqBodyByte)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", wUrl, reader)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var loginPic model.LoginPicRes
	json.Unmarshal(resBody, &loginPic)
	global.LoginCodeMap.Store(loginFlag, global.CodeInfo{
		CanLogin:   false,
		OpenId:     "",
		CreateTime: time.Now().Unix(),
	})
	return &loginPic, nil
}

func (w *WxLoginPassport) GetUserInfo(openid string) (wxUserInfo *model.WXUserInfo, err error) {
	wUrl := fmt.Sprintf(
		"https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN",
		global.WxGlobal.AccessToken,
		openid,
	)
	client := http.Client{}
	req, err := http.NewRequest("GET", wUrl, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var userInfo model.WXUserInfo
	bodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bodyByte, &userInfo)
	return &userInfo, nil
}
