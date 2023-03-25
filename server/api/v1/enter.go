package v1

import "lf_web_gin/server/api/v1/system"

type ApiGroup struct {
	System system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
