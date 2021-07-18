package v1

import (
	"backend/utils/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentUser(c *gin.Context) {
	code := common.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":        code,
		"msg":         common.GetMsg(code),
		"name":        "admin",
		"userid":      1,
		"avatar":      "https://pic4.zhimg.com/80/v2-867a95c44703177811f2590b09396113_1440w.jpg?source=1940ef5c",
		"unreadCount": 11,
	})
}
