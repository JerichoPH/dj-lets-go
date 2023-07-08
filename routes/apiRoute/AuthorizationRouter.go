package apiRoute

import (
	"dj-lets-go/controllers"
	"dj-lets-go/middlewares"

	"github.com/gin-gonic/gin"
)

// AuthorizationRouter 权鉴路由
type AuthorizationRouter struct{}

// NewAuthorizationRouter 构造函数
func NewAuthorizationRouter() AuthorizationRouter {
	return AuthorizationRouter{}
}

// Load 加载路由
func (AuthorizationRouter) Load(engine *gin.Engine) {
	r := engine.Group("api/authorization")
	{
		// 登陆
		r.POST("login", controllers.NewAuthorizationController().Login)

		// 注册
		r.POST("register", controllers.NewAuthorizationController().Register)
	}

	r2 := engine.Group("api/authorization", middlewares.CheckAuthorization())
	{
		// 检查当前用户是否登录
		r2.Any("checkIsLogin", controllers.NewAuthorizationController().AnyCheckIsLogin)

		// 获取当前用户角色
		r2.GET("roles", controllers.NewAuthorizationController().GetRoles)
		// 获取当前用户权限
		r2.GET("permissions", controllers.NewAuthorizationController().GetPermissions)
		// 获取当前用户菜单
		r2.GET("menus", controllers.NewAuthorizationController().GetMenus)
	}
}
