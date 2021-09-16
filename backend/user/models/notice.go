package models

import (
	"time"
)

type Notice struct {
	ID           int `gorm:"primary_key"`
	UserId       int
	Avatar       string
	Title 		 string
	Datetime     time.Time `gorm:"column:date_time;default:null"`
	Type         string
	Status       string
	Read         bool
	Model
}


func (_ *Notice) TableName() string {
	return "user_notice"
}

// 用ID获取用户
func GetNotice(ID interface{}) (Notice, error) {
	var notice Notice
	result := Db.First(&notice, ID)
	return notice, result.Error
}

func GetNotices(userId int) (Notice, error) {
	var notice Notice
	result := Db.Where("user_id = ?", userId).First(&notice)
	return notice, result.Error
}



func InsertNotice(name string) bool {
	var notice = Notice{}
	notice.UserId = 1
	notice.Avatar = "https://pic4.zhimg.com/80/v2-867a95c44703177811f2590b09396113_1440w.jpg?source=1940ef5c"
	notice.Title = "测试一下"
	notice.Type = "message"
	notice.Status = "processing"
	notice.Read = true
	Db.Create(&notice)
	return true
}


