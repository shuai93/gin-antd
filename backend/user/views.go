package user

import (
	"backend/models"
	_ "backend/utils/common"
	_ "backend/utils/logging"
	"backend/utils/response"
	_ "github.com/astaxie/beego/core/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentUser(c *gin.Context) {
	result := DefaultUserResponse(UserDemo())
	c.JSON(http.StatusOK, result)
}

func GenerateUser(c *gin.Context) {
	json := make(map[string]string) //注意该结构接受的内容
	c.BindJSON(&json)
	models.GenerateUser(json["name"])
	c.JSON(http.StatusOK, response.Response{}.SuccessResponse())
}

func GetCurrentUserNotice(c *gin.Context) {
	result := make(map[string]interface{})
	notice1 := map[string]interface{}{
		"id":       "000000001",
		"avatar":   "你收到了 14 份新周报",
		"title":    "你收到了 14 份新周报",
		"datetime": "2017-08-09",
		"type":     "notification",
	}

	notice2 := map[string]interface{}{
		"id":          "000000009",
		"title":       "任务名称",
		"description": "任务需要在 2017-01-12 20:00 前启动",
		"extra":       "未开始",
		"status":      "todo",

		"type": "event",
	}

	notice3 := map[string]interface{}{
		"id":          "000000007",
		"avatar":      "https://gw.alipayobjects.com/zos/rmsportal/fcHMVNCjPOsbUGdEduuv.jpeg",
		"title":       "朱偏右 回复了你",
		"description": "这种模板用于提醒谁与你发生了互动，左侧放『谁』的头像",
		"datetime":    "2017-08-07",
		"type":        "message",
		"clickClose":  true,
	}
	notice4 := map[string]interface{}{
		"id":       "000000002",
		"avatar":   "你收到了 14 份新周报",
		"title":    "你收到了 14 份新周报",
		"datetime": "2017-08-09",
		"type":     "notification",
		"status":   "processing",
		"read":     true,
	}

	data := [...]map[string]interface{}{notice1, notice2, notice3, notice4}

	result["data"] = data

	c.JSON(http.StatusOK, result)
}
