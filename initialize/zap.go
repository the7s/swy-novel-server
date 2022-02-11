package initialize

import (
	"fmt"
	"github.com/the7s/swy-novel-server/global"
	"github.com/the7s/swy-novel-server/library/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func Zap() (logger *zap.Logger) {
	// 判断文件是否存在，若不存在，则创建
	if ok, _ := utils.PathExist(global.SWY_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory \n", global.SWY_CONFIG.Zap.Director)
		_ = os.Mkdir(global.SWY_CONFIG.Zap.Director, os.ModePerm)
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
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", global.SWY_CONFIG.Zap.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", global.SWY_CONFIG.Zap.Director), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", global.SWY_CONFIG.Zap.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", global.SWY_CONFIG.Zap.Director), errorPriority),
	}

	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if global.SWY_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger

}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := getWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.SWY_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}

	if global.SWY_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.SWY_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {

	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.SWY_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	switch {
	case global.SWY_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.SWY_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.SWY_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.SWY_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}
