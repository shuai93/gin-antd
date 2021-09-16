package models

import (
	"backend/models"
	_ "backend/models"
	_ "github.com/jinzhu/gorm"
	_ "time"
)

type Article struct {
	ID int `gorm:"primary_key" json:"id"`

	models.Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (c *Article) TableName() string {
	return "blog_article"
}

func ExistArticleByID(id int) bool {
	var article Article
	models.Db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int64) {
	models.Db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	models.Db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	models.Db.Where("id = ?", id).First(&article)
	//db.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) bool {
	models.Db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	models.Db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	models.Db.Where("id = ?", id).Delete(Article{})
	return true
}

func CleanAllArticle() bool {
	models.Db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})

	return true
}
