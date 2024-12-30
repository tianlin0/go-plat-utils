package logs

import (
	"context"
)

// ctxLogger 自定义日志的使用方式
type ctxLogger struct {
	ctx         context.Context //存储当前上下文
	logExecute  LogExecute
	logCommData *LogCommData
	logLevel    LogLevel
	callerSkip  int
}

//// NewCtxLogger 例子，一个完整的日志，需要实现如下方法
//func NewCtxLogger(ctx context.Context, level LogLevel, logFun LogExecute, logCommData ...*LogCommData) (ILogger, context.Context) {
//	logInstance := buildLogger(ctx, getCtxLoggerFromContext(ctx), level, logFun)
//	if logCommData != nil && len(logCommData) > 0 {
//		logInstance.logCommData = logCommData[0].Init()
//	} else {
//		if logInstance.logCommData == nil {
//			logInstance.logCommData = &LogCommData{}
//		}
//	}
//	return setLoggerToContext(ctx, logInstance)
//}
//
//// GetCtxLogger 只是取得ctx实例，为空则使用默认的
//func (cl ctxLogger) GetCtxLogger(ctx context.Context) ILogger {
//	ctxLoggerInstance := getCtxLoggerFromContext(ctx)
//	if ctxLoggerInstance != nil {
//		return ctxLoggerInstance
//	}
//	return DefaultLogger()
//}
//
//// 获取一个logger
//func getOneLogger(ctx context.Context) *ctxLogger {
//	logInstance := getCtxLoggerFromContext(ctx) //获取原始的
//	if logInstance != nil {
//		return logInstance
//	}
//	return buildLogger(ctx, logInstance, INFO, nil) //默认使用新的
//}
//
//// 从ctx获取logger
//func getCtxLoggerFromContext(ctx context.Context) *ctxLogger {
//	if ctx == nil {
//		return nil
//	}
//	if logInfoTemp := ctx.Value(defaultConfig.LoggerCtxName); logInfoTemp != nil {
//		if logInstance, ok := logInfoTemp.(*ctxLogger); ok {
//			return logInstance
//		}
//	}
//	return nil
//}
//
//func setLoggerToContext(ctx context.Context, logger *ctxLogger) (*ctxLogger, context.Context) {
//	if ctx == nil {
//		ctx = context.Background()
//	}
//	if logger == nil {
//		//直接拿原始的
//		logger = getOneLogger(ctx)
//	}
//	newCtx := context.WithValue(ctx, defaultConfig.LoggerCtxName, logger)
//	logger.ctx = newCtx
//	return logger, newCtx
//}
//
//func buildLogger(ctx context.Context, logInstance *ctxLogger, level LogLevel, logFun LogExecute) *ctxLogger {
//	if ctx == nil {
//		ctx = context.Background()
//	}
//	if logInstance == nil {
//		if logFun == nil {
//			logFun = func(ctx context.Context, logInfo *LogData) {
//				if logInfo == nil || len(logInfo.Message) == 0 {
//					return
//				}
//				//默认是控制台打印
//				logger := NewPrintLogger(logInfo.LogLevel, &logInfo.LogCommData)
//				logger.SetCallerSkip(6)
//				Logger(logger, logInfo.LogLevel, logInfo.Message...)
//			}
//		}
//		logInstance = &ctxLogger{
//			ctx:        ctx,
//			logLevel:   level,
//			logExecute: logFun,
//			callerSkip: 2,
//		}
//	}
//
//	if level >= DEBUG {
//		logInstance.SetLevel(level)
//	}
//	if logFun != nil {
//		logInstance.logExecute = logFun
//	}
//	if logInstance.logCommData == nil {
//		logInstance.logCommData = (&LogCommData{}).Init()
//	}
//	if logInstance.logCommData.LogId == "" {
//		logInstance.logCommData.LogId = httputil.GetLogId()
//	}
//
//	return logInstance
//}
