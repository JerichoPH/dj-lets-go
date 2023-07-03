package webRoute

import (
	"dj-lets-go/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthorizationRouter 权鉴路由
type AuthorizationRouter struct{}

// Load 加载路由
func (AuthorizationRouter) Load(engine *gin.Engine) {
	r := engine.Group("authorization")
	{
		// 登陆
		r.GET("login", func(ctx *gin.Context) {
			engine.LoadHTMLFiles("templates/Authorization/login.html")
			ctx.HTML(http.StatusOK, "login.html", types.MapStringToAny{})
		})

		r.GET("register", func(ctx *gin.Context) {
			engine.LoadHTMLFiles("templates/Authorization/register.html")
			ctx.HTML(http.StatusOK, "register.html", types.MapStringToAny{})
		})
	}
}
