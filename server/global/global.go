package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"lf_web_gin/server/config"

	"golang.org/x/sync/singleflight"
)

//全局变量

var (
	PRO_DB                  *gorm.DB
	PRO_VIPER               *viper.Viper
	PRO_CONFIG              config.Server
	PRO_LOG                 *zap.Logger
	PRO_REDIS               *redis.Client
	GVA_Concurrency_Control = &singleflight.Group{}
)

/*// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}*/
