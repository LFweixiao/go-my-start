package system

import (
	"lf_web_gin/server/server"
)

type ApiGroup struct {
	SystemUserApi
	SystemFileApi
}

var (
	userServer = server.ServerGroupApp.SystemServiceGroup.UserServer
	fileServer = server.ServerGroupApp.SystemServiceGroup.FileServer
)
