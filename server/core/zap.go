package core

import (
	"os"
	"server/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger 初始化日志
func InitLogger() *zap.Logger {
	zapConfig := global.Config.Zap

	// 获取日志写入器
	writeSyncer := getWriteSyncer(zapConfig.Filename, zapConfig.MaxSize, zapConfig.MaxBackups, zapConfig.MaxAge)

	// 是否在控制台打印
	if zapConfig.IsConsolePrint {
		writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
	}

	// 获取日志编码器
	encoder := getEncoderConfig()

	// 解析日志级别
	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(zapConfig.Level)); err != nil {
		logLevel = zapcore.InfoLevel
	}

	// 创建核心
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	// 创建日志实例
	logger := zap.New(core, zap.AddCaller())

	return logger
}

// getWriteSyncer 获取日志写入器
func getWriteSyncer(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,   // 日志文件路径
		MaxSize:    maxSize,    // 最大日志文件大小, 单位MB
		MaxBackups: maxBackups, // 最大保留日志文件数量W
		MaxAge:     maxAge,     // 最大保留天数
	}
	return zapcore.AddSync(lumberjackLogger)
}

// getEncoderConfig 获取日志编码器配置
func getEncoderConfig() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
