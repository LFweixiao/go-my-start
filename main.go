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
	global.PRO_DB = initialize.Gorm()

	//初始化 log 日志

	if global.PRO_DB != nil {
		initialize.RegisterTables() //初始化表
		//程序结束钱关闭数据库联机
		db, _ := global.PRO_DB.DB()
		defer db.Close()
	}

	//注册路由
	core.RunWindowsServer()
}
