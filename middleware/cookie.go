package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qm012/sessions"
	"net/http"
	"sessionsdemo/api"
)

//检查是否登录才能访问登录后才能访问的页面
func ValidLogin() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if login := sessions.GetSession(ctx, "session_id").Get(api.IsLoginStr); login != nil {
			switch login.(type) {
			case bool:

			case string:
				if login == "" {
					goto Label
				}
			}
			ctx.Next()
			return
		}

	Label:
		path := fmt.Sprintf("/login?next=%s", ctx.Request.URL.Path)
		//ctx.HTML(http.StatusInternalServerError, "login.html?", gin.H{"err": "亲，需要登录哦！"})
		ctx.Redirect(http.StatusFound, path)
		ctx.Abort()
		return
	}
}
