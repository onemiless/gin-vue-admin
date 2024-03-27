package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	wework "github.com/go-laoji/wecom-go-sdk/v2"
)

var wecomSdk wework.IWeWork

func Wecom() (wecomSdk wework.IWeWork) {
	wecomCfg := global.GVA_CONFIG.Wecom
	wecomSdk = wework.NewWeWork(wework.WeWorkConfig{CorpId: wecomCfg.CorpId})
	wecomSdk.SetAppSecretFunc(func(corpId uint) (corpid string, secret string, customizedApp bool) {
		return wecomCfg.CorpId, wecomCfg.AppSecret, true
	})
	return wecomSdk
	//router := gin.Default()
	//
	//router.GET("/user/list", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"data": wecomSdk.UserList(1, 1, 1)})
	//})
	//fmt.Print(wecomSdk.UserList(1, 1, 1))

	//redisCfg := global.GVA_CONFIG.Redis
	//client := redis.NewClient(&redis.Options{
	//	Addr:     redisCfg.Addr,
	//	Password: redisCfg.Password, // no password set
	//	DB:       redisCfg.DB,       // use default DB
	//})
	//pong, err := client.Ping(context.Background()).Result()
	//if err != nil {
	//	global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	//	panic(err)
	//} else {
	//	global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
	//	global.GVA_REDIS = client
	//}
}
