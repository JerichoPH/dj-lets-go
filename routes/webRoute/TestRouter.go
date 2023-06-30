package webRoute

import (
	"dj-lets-go/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WsTestRouter struct{}

func (WsTestRouter) Load(engine *gin.Engine) {

	r := engine.Group("test")
	{
		r.GET("ws", func(ctx *gin.Context) {
			engine.LoadHTMLFiles("templates/Test/ws.html")
			ctx.HTML(http.StatusOK, "ws.html", types.MapStringToAny{})
		})

		r.GET("sse", func(ctx *gin.Context) {
			engine.LoadHTMLFiles("templates/Test/sse.html")
			ctx.HTML(http.StatusOK, "sse.html", types.MapStringToAny{})
		})
	}

}
