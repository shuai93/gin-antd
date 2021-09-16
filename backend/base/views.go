package base

import (
	"backend/base/ws"
	"backend/user/models"
	"backend/utils/common"
	"backend/utils/gredis"
	"backend/utils/logging"
	"backend/utils/response"
	"encoding/json"
	"github.com/astaxie/beego/core/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	type User struct {
		Username string `json:"username";valid:"Required; MaxSize(50)"`
		Password string `json:"password";valid:"Required; MaxSize(50)"`
	}
	var user User

	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    4001,
			"message": "parameter error",
		})
	}
	username := user.Username
	password := user.Password

	valid := validation.Validation{}
	ok, _ := valid.Valid(&user)

	data := make(map[string]interface{})
	code := response.SUCCESS

	if ok {
		success, user := models.CheckPassword(username, password)
		if success {
			token, err := common.GenerateToken(user.Username, user.PasswordDigest, user.ID)
			if err != nil {
				code = response.ErrorAuthToken
			} else {
				data["token"] = token
				data["id"] = user.ID
				data["username"] = user.Username
				data["role"] = user.Role
				code = response.SUCCESS
			}

		} else {
			for _, err := range valid.Errors {
				logging.Info(err.Key, err.Message)
			}
			code = response.ErrorAuthToken
		}
	} else {
		for _, err := range valid.Errors {
			logging.Warn(err.Key, err.Message)
		}
	}

	result := response.Response{}
	result.Code = code
	result.Msg = response.Msg[code]
	result.Data = data

	c.JSON(http.StatusOK, result)
}

func Redis(context *gin.Context) {

	b, _ := gredis.Get("hub")
	//s, _ := gredis.GetString("hub")
	//atoi, _ := strconv.Atoi(s)
	//i := 1 / atoi
	data := json.Unmarshal(b, ws.Hub{})
	context.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
