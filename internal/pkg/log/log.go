package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-start/config"
	"os"
	"time"
)

var Logger *logrus.Logger

func NewLogger() {
	Logger = logrus.New()
	logCfg := config.Cfg.Log
	filePath := logCfg.DirPath
	fileName := logCfg.FileName
	file := filePath + fileName

	Logger.Out = os.Stdout
	Logger.SetLevel(logrus.DebugLevel)

	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logWriter, _ := rotatelogs.New(
		file+".%Y%m%d.log",
		//生成软链 指向最新的日志文件
		rotatelogs.WithLinkName(file),
		//文件最大保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志切割时间间隔(1天)(隔多久分割一次)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Logger.AddHook(lfHook)
}
