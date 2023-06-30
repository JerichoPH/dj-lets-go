package webRoute

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeRouter struct{}

func (HomeRouter) Load(engine *gin.Engine) {
	engine.StaticFS("/public", http.Dir("public"))
	r := engine.Group("")
	{
		r.GET("", func(ctx *gin.Context) {
			engine.LoadHTMLGlob("templates/Home/*")
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}
}
