package user

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	{
		r.GET("/currentUser", GetCurrentUser)
		r.POST("/generateUser", GenerateUser)
		r.GET("/notices", GetCurrentUserNotice)
	}
	//api := r.Group("/user")

}
