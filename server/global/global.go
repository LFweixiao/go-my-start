package global

import (
	"github.com/spf13/viper"
	"lf_web_gin/server/config"
)

//全局变量

var (
	PRO_VIPER  *viper.Viper
	PRO_CONFIG config.Server
)
