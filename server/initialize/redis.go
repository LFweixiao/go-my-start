package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"lf_web_gin/server/global"
)

func Redis() {
	redisCfg := global.PRO_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.PRO_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.PRO_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.PRO_REDIS = client
		fmt.Printf(global.PRO_REDIS.String())
	}
}
