package jwt

import (
	"backend/utils/common"
	"backend/utils/logging"
	"backend/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var token string

		code = response.SUCCESS

		token = c.Query("token")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
		}

		logging.Info("token is %s", token)
		if token == "" {
			code = response.ErrorAuth
		} else {
			claims, err := common.ParseToken(token)
			if err != nil {
				code = response.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = response.ErrorAuthCheckTokenTimeout
			}
			if claims != nil {
				c.Set("username", claims.Username)
				c.Set("userId", claims.UserId)
			}
		}

		if code != response.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  response.Msg[code],
				"data": data,
			})

			c.Abort()
			return
		}
		logging.Info("token认证成功")
		c.Next()
	}
}
