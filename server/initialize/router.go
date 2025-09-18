package initialize

import (
	"server/global"
	"server/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()

	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	Router.Use(sessions.Sessions("session", store))

	// 路由分组
	routerGroup := router.RouterGroupApp

	publicGroup := Router.Group(global.Config.System.RouterPrefix)
	{
		routerGroup.InitBaseRouter(publicGroup)
	}

	return Router
}
