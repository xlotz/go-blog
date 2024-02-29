package routers

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
)

type RouterGroup struct {
	//*gin.Engine
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	//routerGroupApp := RouterGroup{router}
	//routerGroupApp.SettingsRouter()
	// 对路由进行分组
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingsRouter()
	return router
}
