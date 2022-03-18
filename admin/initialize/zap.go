package initialize

import (
	"catering/global"
	"catering/pkg"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap() {
	cfg := global.Config.Zap
	// 判断是否有Director文件夹
	if ok, _ := pkg.PathExists(cfg.Director); !ok {
		fmt.Printf("create %v directory\n", cfg.Director)
		_ = os.Mkdir(cfg.Director, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", cfg.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", cfg.Director), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", cfg.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", cfg.Director), errorPriority),
	}

	logger := zap.New(zapcore.NewTee(cores[:]...))

	//用文件名、行号和zap调用者的函数名注释每条消息
	if cfg.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	global.Log = logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	cfg := global.Config.Zap
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  cfg.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case cfg.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case cfg.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case cfg.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case cfg.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	cfg := global.Config.Zap
	if cfg.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := pkg.GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	cfg := global.Config.Zap
	enc.AppendString(t.Format(cfg.Prefix + "2006-01-02 15:04:05.000"))
}
