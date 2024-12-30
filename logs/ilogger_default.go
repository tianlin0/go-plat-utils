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
	defaultConfig = &Config{}
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
	if ctx != nil {
		goroutines.SetContext(&ctx)
	}
	return DefaultLogger()
}
