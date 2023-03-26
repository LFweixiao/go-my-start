package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"lf_web_gin/server/global"
)

// Viper 解析配置文件
// 命令行 > 配置 > 默认
func Viper(path ...string) *viper.Viper {
	//TODO 命令行

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	// ./是当前project 所在的目录
	v.AddConfigPath("./")
	//查找并读取配置文
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//开启事件监听 配置热更新
	v.WatchConfig()
	//监听到热更新时 执行下面内容
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
		if err = v.Unmarshal(&global.PRO_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	//配置放入全局变量
	if err = v.Unmarshal(&global.PRO_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
