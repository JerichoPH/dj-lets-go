package webRoute

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

// Load 加载路由
func (AccountRouter) Load(engine *gin.Engine) {
	r := engine.Group("account")
	{
		// 列表
		r.GET("", func(ctx *gin.Context) {
			engine.LoadHTMLFiles("templates/Account/index.html")
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}
}
