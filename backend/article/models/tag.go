package models

import (
	"backend/models"
	_ "gorm.io/gorm"
	_ "time"
)

type Tag struct {
	ID int `gorm:"primary_key" json:"id"`

	models.Model

	Name       string `json:"name"`
	CreateDby  string `json:"created_by"`
	ModifieDby string `json:"modified_by"`
	State      int    `json:"state"`
}

func (c *Tag) TableName() string {
	return "blog_tag"
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	models.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int64) {
	models.Db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	models.Db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createDby string) bool {
	models.Db.Create(&Tag{
		Name:      name,
		State:     state,
		CreateDby: createDby,
	})

	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	models.Db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	models.Db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	models.Db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

func CleanAllTag() bool {
	models.Db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}
