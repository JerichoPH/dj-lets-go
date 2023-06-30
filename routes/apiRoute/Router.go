package apiRoute

import (
	"github.com/gin-gonic/gin"
)

// Router 分组路由
type Router struct{}

// Register 组册路由
func (Router) Register(engine *gin.Engine) {
	NewTestRouter().Load(engine)          // 测试
	NewAuthorizationRouter().Load(engine) // 权鉴
	NewAccountRouter().Load(engine)       // 用户
}
