package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"lf_web_gin/server/core/internal"
	"lf_web_gin/server/global"
	"lf_web_gin/server/utils"
	"os"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.PRO_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.PRO_CONFIG.Zap.Director)
		_ = os.Mkdir(global.PRO_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.PRO_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
