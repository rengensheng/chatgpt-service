package logger

import (
	"time"

	"github.com/goylold/lowcode/config"
	"github.com/hyahm/golog"
)

func init() {
	configBase := config.GetConConfig()
	golog.InitLogger(configBase.Log.LogFilePath+"/"+configBase.Log.LogFileName, 0, true, time.Hour*24*time.Duration(configBase.Log.Day))
}

func InitLogger() {
	golog.Info("加载日志配置....")
}
