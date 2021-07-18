package routers

import (
	"backend/api"
	"backend/api/v1"
	"backend/api/ws"
	"backend/middleware/casbin"
	"backend/middleware/jwt"
	"backend/utils/gredis"
	"backend/utils/setting"
	"encoding/json"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func redis(context *gin.Context) {

	b, _ := gredis.Get("hub")

	data := json.Unmarshal(b, ws.Hub{})
	context.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("views/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", "flysnow_org")
	})

	r.GET("/redis", redis)

	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/login", api.GetAuth)

	hub := ws.NewHub()
	go hub.Run()

	r.GET("/ws", func(context *gin.Context) {
		ws.Handler(context, hub)
	})

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	apiv1.Use(casbin.Casbin())

	{
		apiv1.GET("/message/:id", func(context *gin.Context) {
			ws.GetMessage(context, hub)
		})
		apiv1.POST("/message", func(context *gin.Context) {
			ws.PushMessage(context, hub)
		})

	}

	{
		apiv1.GET("/currentUser", v1.GetCurrentUser)
	}

	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
