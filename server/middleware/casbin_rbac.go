package middleware

import (
	"go.uber.org/zap"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		//获取请求的PATH
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, global.GVA_CONFIG.System.RouterPrefix)
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.AuthorityId))
		e := casbinService.Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			//打印对应的错误信息对应的参数
			global.GVA_LOG.Info("权限不足", zap.Any("role", sub), zap.Any("path", obj), zap.Any("act", act),
				zap.Any("waitUse", waitUse), zap.Any("path", path))

			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
