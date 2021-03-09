package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/qm012/sessions"
	"sessionsdemo/api"
	"sessionsdemo/middleware"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	var age = 60

	sessions.SetCookie(sessions.ValueMap, "session_id", age, "/", "127.0.0.1", false, false)
	sessions.SetCookie(sessions.ValueString, "password", age, "/", "127.0.0.1", false, false)
	//内存版
	r.Use(sessions.Sessions(sessions.ChooseSessionStore(sessions.Memory)))

	//redis 版本
	//r.Use(sessions.Sessions(sessions.ChooseSessionStore(sessions.Memory,rdbClient)))

	publicGroup := r.Group("")
	{
		publicGroup.GET("/login", api.LoginHandler)
		publicGroup.POST("/login", api.LoginHandler)
		publicGroup.GET("/index", api.IndexHandler)
	}

	privateGroup := r.Group("").Use(middleware.ValidLogin())
	{
		privateGroup.GET("/home", api.HomeHandler)
		privateGroup.GET("/vip", api.VipHandler)
		privateGroup.GET("/logout", api.Logout)
	}
	return r
}
