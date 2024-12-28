package logs_test

import (
	"context"
	"fmt"
	"github.com/tianlin0/go-plat-utils/goroutines"
	"github.com/tianlin0/go-plat-utils/logs"
	"testing"
	"time"
)

func TestLoggerData(t *testing.T) {
	logData := logs.NewLogData(&logs.LogCommData{
		CreateTime: time.Now(),
		LogId:      "logid12345",
		UserId:     "userid44444",
		Env:        "dev",
		Path:       "/name/pri",
		Method:     "get",
		Extend: map[string]interface{}{
			"service": "dnf",
		},
	})

	time.Sleep(5 * time.Millisecond)

	logData.AddMessage(logs.DEBUG, "fileName.go", 112, "有了错误", "有了第二个错误")

	str := logData.String()

	fmt.Println(str)

}
func TestPrintLogger(t *testing.T) {
	prtLogger := logs.NewPrintLogger(logs.DEBUG, &logs.LogCommData{
		CreateTime: time.Now(),
		LogId:      "",
		UserId:     "userid44444",
		Env:        "dev",
		Path:       "/name/pri",
		Method:     "get",
		Extend: map[string]interface{}{
			"service": "dnf",
		},
	})

	time.Sleep(5 * time.Millisecond)

	ctx := context.Background()
	sonCtx := context.WithValue(ctx, "logid", "1234567")
	goroutines.SetContext(&sonCtx)

	prtLogger.Debug("有了错误", "有了第二个错误")
	prtLogger.Error("有了第二个错误dddd")

}
func TestDefaultLogger(t *testing.T) {
	prtLogger := logs.DefaultLogger()

	prtLogger.Debug("有了错误", "有了第二个错误")
	prtLogger.Error("有了第二个错误dddd")

	ctx := context.Background()
	sonCtx := context.WithValue(ctx, "logid", "1234567")

	ctxLogger := logs.CtxLogger(sonCtx)
	ctxLogger.Info("aaaaa")

}
