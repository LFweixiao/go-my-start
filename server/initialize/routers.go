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

	//文件传入大小
	Router.MaxMultipartMemory = 8 << 20 // 8MiB

	// swagger 再此配置

	//前置路由
	PrivateGroup := Router.Group("/lf")
	{
		systemRouter.InitUserApiRouter(PrivateGroup) // 注册用户api
		systemRouter.InitFileApiRouter(PrivateGroup) // 注册文件api
	}

	return Router
}
