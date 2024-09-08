package logger

import (
	"github.com/open-auth/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Zap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSettings) *Zap {
	logLevel := config.Level
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncodeLog()
	sync := getWriteSync(config)
	core := zapcore.NewCore(encoder, sync, level)

	return &Zap{
		zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	}
}

func getEncodeLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// format timestamp
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> Time
	encoderConfig.TimeKey = "time"

	// info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller": "cli/main.log.go:3"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSync(config setting.LoggerSettings) zapcore.WriteSyncer {
	hook := lumberjack.Logger{
		Filename:   config.FileName,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	syncConsole := zapcore.AddSync(os.Stdout)
	return zapcore.NewMultiWriteSyncer(syncConsole, zapcore.AddSync(&hook))
}
