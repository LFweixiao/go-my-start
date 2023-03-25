package initialize

import (
	"github.com/gin-gonic/gin"
	"lf_web_gin/server/router"
)

//初始化路由

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	//exampleRouter := router.RouterGroupApp.Example

	// swagger 再此配置

	//前置路由
	PrivateGroup := Router.Group("/lf")
	{
		systemRouter.InitUserApiRouter(PrivateGroup) // 注册用户api
	}

	return Router
}
