package base

import (
	"backend/article"
	"backend/base/ws"
	"backend/middlewares/casbin"
	"backend/middlewares/jwt"
	"backend/middlewares/recover"
	"backend/user"
	"backend/utils/setting"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {

	var hub = ws.NewHub()
	go hub.Run()

	r := gin.Default()
	r.Use(recover.Recover)
	r.Use(gin.Logger())
	gin.SetMode(setting.ServerSetting.RunMode)

	// swagger 接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// redis 测试接口
	r.GET("/redis", Redis)

	// 登录接口
	r.POST("/api/login", UserLogin)

	// web socket 接口
	wsRouter(r, hub)

	api := r.Group("/api/v1")
	api.Use(jwt.JWT())
	api.Use(casbin.Casbin())

	messageRouter(api, hub)
	user.InitRouter(api)
	article.InitRouter(api)

	return r
}

func wsRouter(r *gin.Engine, hub *ws.Hub) {

	r.LoadHTMLFiles("main/ws/views/index.html")
	// ws html 测试
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", "full_stack")
	})

	r.GET("/ws", func(context *gin.Context) {
		ws.Handler(context, hub)
	})
}

func messageRouter(api *gin.RouterGroup, hub *ws.Hub) {
	{
		api.GET("/message/:id", func(context *gin.Context) {
			ws.GetMessage(context, hub)
		})
		api.POST("/message", func(context *gin.Context) {
			ws.PushMessage(context, hub)
		})

	}
}
