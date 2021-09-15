package casbin

import (
	"backend/models"
	"backend/utils/logging"
	"backend/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		username, _ := c.Get("username")
		sub := username.(string)
		e := models.Casbin()
		logging.Info(obj, act, sub)
		success := e.Enforce(sub, obj, act)
		if !success {
			logging.Info("e.Enforce err: %s", "很遗憾,权限验证没有通过")
			c.JSON(http.StatusForbidden, gin.H{
				"code": response.ErrorPermission,
				"msg":  response.Msg[response.ErrorPermission],
				"data": "鉴权失败",
			})
			c.Abort()
			return
		}
		logging.Info("恭喜您,权限验证通过")
		c.Next()

	}
}
