package core

import "lf_web_gin/server/initialize"

type server interface {
	ListenAndServer() error
}

func RunWindowsServer() {
	//redis

	//db 数据库

	//读取配置文件

	initialize.Routers().Run()
}
