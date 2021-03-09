package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qm012/sessions"
	"net/http"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

var (
	UsernameStr = "username"
	PasswordStr = "password"
	IsLoginStr  = "isLogin"
)

func Logout(ctx *gin.Context) {
	sessions.GetSession(ctx, "password").Delete()
	sessions.GetSession(ctx, "session_id").Delete()
	ctx.JSON(http.StatusOK, gin.H{"message": "成功啦"})
}

func LoginHandler(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodPost {
		var user User
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}
		sessions.GetSession(ctx, "password").Set(user.Password)
		sessions.GetSession(ctx, "session_id").Set(true, IsLoginStr)
		sessions.GetSession(ctx, "session_id").Set(user.Username, UsernameStr)
		sessions.GetSession(ctx, "session_id").Set(user.Password, PasswordStr)
		//ctx.HTML(http.StatusMovedPermanently, "home.html", gin.H{"username": user.Username})
		nextPath := ctx.DefaultQuery("next", "/index")
		fmt.Println("next =", ctx.Query("next"))
		ctx.Redirect(http.StatusMovedPermanently, nextPath)
		return
	}
	ctx.HTML(http.StatusOK, "login.html", nil)
}
func IndexHandler(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodPost {
		return
	}
	ctx.HTML(http.StatusOK, "index.html", nil)
}
func HomeHandler(ctx *gin.Context) {

	if ctx.Request.Method == http.MethodPost {

		return
	}
	username := sessions.GetSession(ctx, "session_id").Get(UsernameStr)
	ctx.HTML(http.StatusOK, "home.html", gin.H{"username": username})
}
func VipHandler(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodPost {

		return
	}
	username := sessions.GetSession(ctx, "session_id").Get(UsernameStr)
	password := sessions.GetSession(ctx, "password").Get()
	login := sessions.GetSession(ctx, "session_id").Get(IsLoginStr)

	ctx.HTML(http.StatusOK, "vip.html", gin.H{
		"username": username,
		"password": password,
		"login":    login,
	})
}
