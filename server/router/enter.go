package router

import (
	"lf_web_gin/server/router/example"
	"lf_web_gin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
