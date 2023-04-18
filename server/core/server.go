package core

import (
	"lf_web_gin/server/global"
	"lf_web_gin/server/initialize"
)

type server interface {
	ListenAndServer() error
}

func RunWindowsServer() {
	//redis
	if global.PRO_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	//db 数据库

	//读取配置文件

	initialize.Routers().Run()
}
