package models

import (
	"backend/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int `gorm:"primary_key"`
	Username       string
	PasswordDigest string
	Nickname       string
	Status         string
	Mobile         string
	Role           string
	Avatar         string `gorm:"size:1000"`
	SuperUser      bool
	models.Model
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

func (_ *User) TableName() string {
	return "user_user"
}

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := models.Db.First(&user, ID)
	return user, result.Error
}

func GetUserByUserName(username string) (User, error) {
	var user User
	result := models.Db.Where("username = ?", username).First(&user)
	return user, result.Error
}

// CheckPassword 校验密码
func CheckPassword(username string, password string) (bool, User) {
	info, _ := GetUserByUserName(username)
	err := bcrypt.CompareHashAndPassword([]byte(info.PasswordDigest), []byte(password))
	return err == nil, info
}

func InsertUser(name string) bool {
	var user = User{}
	user.Username = name
	_ = user.SetPassword(name)
	user.Nickname = name
	user.Status = "active"
	user.Role = "admin"
	user.Mobile = "13888888888"
	user.SuperUser = true
	user.Avatar = "https://pic4.zhimg.com/80/v2-867a95c44703177811f2590b09396113_1440w.jpg?source=1940ef5c"
	models.Db.Create(&user)
	return true
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}
