package user

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	{
		r.GET("/current-user", GetCurrentUser)
		r.POST("/generate-user", GenerateUser)
		r.GET("/notices", GetCurrentUserNotice)
		r.POST("/generate-notice", GenerateNotice)
	}
	//api := r.Group("/user")

}
