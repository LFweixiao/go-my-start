package initialize

import (
	"gorm.io/gorm"
	"lf_web_gin/server/global"
)

// Gorm 初始化数据链接
func Gorm() *gorm.DB {
	switch global.PRO_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()

	default:
		return GormMysql()
	}
}
