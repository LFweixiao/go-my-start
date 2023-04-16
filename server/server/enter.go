package server

import "lf_web_gin/server/server/system"

type ServerGroup struct {
	SystemServiceGroup system.ServerGroup
}

var ServerGroupApp = new(ServerGroup)
