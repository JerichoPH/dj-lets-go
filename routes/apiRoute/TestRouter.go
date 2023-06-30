package apiRoute

import (
	"dj-lets-go/controllers"

	"github.com/gin-gonic/gin"
)

// TestRouter 路由
type TestRouter struct{}

// NewTestRouter 构造函数
func NewTestRouter() TestRouter {
	return TestRouter{}
}

// Load 加载路由
func (TestRouter) Load(engine *gin.Engine) {
	r := engine.Group(
		"api/test",
		// middlewares.CheckAuthorization(),
		// middlewares.CheckPermission(),
	)
	{
		// 新建
		r.Any("sendToWebsocket", func(ctx *gin.Context) { controllers.NewTestController().AnySendToWebsocket(ctx) })
		r.Any("sendToTcpServer", func(ctx *gin.Context) { controllers.NewTestController().AnySendToTcpServer(ctx) })
		r.Any("sendToTcpClient", func(ctx *gin.Context) { controllers.NewTestController().AnySendToTcpClient(ctx) })
	}
}
