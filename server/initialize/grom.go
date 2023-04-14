package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"lf_web_gin/server/global"
	"lf_web_gin/server/model/system"
	"os"
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

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {
	db := global.PRO_DB
	err := db.AutoMigrate(
		// 系统模块表
		system.SysUser{},
	)
	if err != nil {
		global.PRO_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	//global.PRO_LOG.Info("register table success")
}
