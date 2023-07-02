package webRoute

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// AuthorizationRouter 权鉴路由
type AuthorizationRouter struct{}

// Load 加载路由
func (AuthorizationRouter) Load(engine *gin.Engine) {
	r := engine.Group("authorization")
	{
		r.GET("login", func(ctx *gin.Context) {
			engine.LoadHTMLGlob("templates/Authorization/login.html")
			ctx.HTML(http.StatusOK, "login.html", gin.H{})
		})
	}
}
