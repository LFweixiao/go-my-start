package main

import (
	"lf_web_gin/server/core"
	"lf_web_gin/server/global"
	"lf_web_gin/server/initialize"
)

func main() {
	//启动项目 读取配置文件
	global.PRO_VIPER = core.Viper()

	//链接数据库
	initialize.Gorm()

	//注册路由
	core.RunWindowsServer()
}
