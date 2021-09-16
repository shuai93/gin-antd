package user

import (
	"backend/user/models"
	"backend/utils/common"
	_ "backend/utils/common"
	_ "backend/utils/logging"
	"backend/utils/response"
	"backend/utils/setting"
	_ "github.com/astaxie/beego/core/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

var key = "username"


func GetCurrentUser(c *gin.Context) {
	username, _ := c.Get(key)
	user, _ := models.GetUserByUserName(username.(string))
	userResponse := DefaultUserResponse(user)
	c.JSON(http.StatusOK, userResponse)
}

func GenerateUser(c *gin.Context) {
	json := make(map[string]string) //注意该结构接受的内容
	_ = c.BindJSON(&json)
	models.InsertUser(json["name"])
	c.JSON(http.StatusOK, response.Response{}.SuccessResponse())
}

func GenerateNotice(c *gin.Context) {
	models.InsertNotice()
	c.JSON(http.StatusOK, response.Response{}.SuccessResponse())
}

func GetCurrentUserNotice(c *gin.Context) {
	result := make(map[string]interface{})
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	maps["user_id"], _ = c.Get("userId")

	data["results"] = models.GetNoticeByUser(common.GetPage(c), setting.AppSetting.PageSize, maps)
	data["count"] = models.GetNoticeTotal(maps)
	result["data"] = data

	c.JSON(http.StatusOK, result)
}

