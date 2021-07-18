package jwt

import (
	"backend/utils/common"
	"backend/utils/logging"
	"backend/utils/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var token string

		code = common.SUCCESS

		token = c.Query("token")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
		}

		logging.Info("token is %s", token)
		if token == "" {
			code = common.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = common.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
			if claims != nil {
				c.Set("username", claims.Username)
			}
		}

		if code != common.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  common.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		logging.Info("token认证成功")
		c.Next()
	}
}
