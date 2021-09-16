package models

import (
	"backend/models"
	"time"
)

type Notice struct {
	ID           int `gorm:"primary_key"`
	UserId       int `json:"user_id"`
	Avatar       string `json:"avatar"`
	Title 		 string `json:"title"`
	Description  string `json:"description"`
	Datetime     time.Time `gorm:"column:date_time;default:null"json:"datetime"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	Read         bool `json:"read"`
	models.Model
}


func (_ *Notice) TableName() string {
	return "user_notice"
}

// 用ID获取用户
func GetNotice(ID interface{}) (Notice, error) {
	var notice Notice
	result := models.Db.First(&notice, ID)
	return notice, result.Error
}

func GetNoticeTotal(maps interface{}) (count int64) {
	models.Db.Model(&Notice{}).Where(maps).Count(&count)
	return
}



func GetNoticeByUser(pageNum int, pageSize int, maps interface{}) (notices []Notice) {
	models.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&notices)

	return
}



func InsertNotice() bool {
	var notice = Notice{}
	notice.UserId = 1
	notice.Avatar = "https://pic4.zhimg.com/80/v2-867a95c44703177811f2590b09396113_1440w.jpg?source=1940ef5c"
	notice.Title = "测试一下"
	notice.Type = "message"
	notice.Status = "processing"
	notice.Read = true
	models.Db.Create(&notice)
	return true
}


