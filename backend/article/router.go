package article

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	{
		//获取标签列表
		r.GET("/tags", GetTags)
		//新建标签
		r.POST("/tags", AddTag)
		//更新指定标签
		r.PUT("/tags/:id", EditTag)
		//删除指定标签
		r.DELETE("/tags/:id", DeleteTag)

		r.GET("/articles", GetArticles)
		//获取指定文章
		r.GET("/articles/:id", GetArticle)
		//新建文章
		r.POST("/articles", AddArticle)
		//更新指定文章
		r.PUT("/articles/:id", EditArticle)
		//删除指定文章
		r.DELETE("/articles/:id", DeleteArticle)
	}

}
