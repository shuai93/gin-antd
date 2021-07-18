package api

import (
	"backend/models"
	"backend/utils/common"
	"backend/utils/logging"
	"backend/utils/util"
	"github.com/astaxie/beego/adapter/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username";valid:"Required; MaxSize(50)"`
	Password string `json:"password";valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {

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
	data["username"] = username
	code := common.INVALID_PARAMS
	status := "error"

	if ok {
		isExist, id := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = common.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				data["id"] = id
				status = "ok"
				code = common.SUCCESS
			}

		} else {
			for _, err := range valid.Errors {
				logging.Info(err.Key, err.Message)
			}
			code = common.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"status": status,
		"msg":    common.GetMsg(code),
		"data":   data,
	})
}
