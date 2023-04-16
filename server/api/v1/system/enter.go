package system

import (
	"lf_web_gin/server/server"
)

type ApiGroup struct {
	SystemUserApi
}

var (
	userServer = server.ServerGroupApp.SystemServiceGroup.UserServer
)
