package logs

import (
	"context"
	"github.com/tianlin0/go-plat-utils/cond"
	"github.com/tianlin0/go-plat-utils/goroutines"
)

type Config struct {
	DefaultLogger ILogger
	LoggerCtxName string //logger在context里的名字
	LogLevel      LogLevel
}

var (
	defaultConfig = &Config{
		LoggerCtxName: "context_logger_name",
		LogLevel:      INFO,
	}
)

// SetConfig 设置默认日志，不能包含ctx，不然全局唯一会有问题
func SetConfig(cfg *Config) {
	if cfg == nil {
		return
	}
	if !cond.IsNil(cfg.DefaultLogger) {
		defaultConfig.DefaultLogger = cfg.DefaultLogger
	}
	if cfg.LogLevel > 0 {
		defaultConfig.LogLevel = cfg.LogLevel
	}

	// 初始化设置
	if defaultConfig.LogLevel > 0 {
		DefaultLogger().SetLevel(defaultConfig.LogLevel)
	}
}

// DefaultLogger 获取系统默认的日志
func DefaultLogger() ILogger {
	if cond.IsNil(defaultConfig.DefaultLogger) {
		logger := NewPrintLogger(DEBUG)
		logger.SetCallerSkip(3)
		defaultConfig.DefaultLogger = logger
	}
	return defaultConfig.DefaultLogger
}

// CtxLogger 获取系统默认的日志
func CtxLogger(ctx context.Context) ILogger {
	// 先从ctx里获取是否存在logger
	if ctx != nil {
		cLogger := getCtxLoggerFromContext(ctx)
		if cLogger != nil {
			goroutines.SetContext(&ctx)
			return cLogger
		}
	}

	// 再从全局保存的context里获取
	oneCtxPtr, _, _ := goroutines.GetContext()
	if oneCtxPtr != nil {
		cLogger := getCtxLoggerFromContext(*oneCtxPtr)
		if cLogger != nil {
			return cLogger
		}
	}
	// 再新建一个logger
	if ctx != nil {
		cLoggerNew, _ := NewCtxLogger(ctx, defaultConfig.LogLevel, nil)
		return cLoggerNew
	}
	return DefaultLogger()
}
