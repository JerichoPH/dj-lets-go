package webRoute

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (Router) Register(engine *gin.Engine) {
	engine.StaticFS("/public", http.Dir("public"))
	
	HomeRouter{}.Load(engine)          // 欢迎页
	AuthorizationRouter{}.Load(engine) // 权鉴
	AccountRouter{}.Load(engine)       // 用户管理
	WsTestRouter{}.Load(engine)        // web-socket-test
}
