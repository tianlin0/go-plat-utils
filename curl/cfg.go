package curl

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"time"
)

var (
	defaultHandler InjectHandler

	defaultMaxCacheTime       = 3600 * 24 * 2 * time.Second //最大用来存2天
	defaultMethod             = http.MethodPost
	defaultReturnKey          = "code"
	defaultReturnVal          = "0"
	defaultPrintLogDataLength = 200 //默认打印日志的时候，数据最长，避免显示太多了

	jsonApi = jsoniter.Config{
		SortMapKeys: true,
	}.Froze()
)

// SetDefaultHandler 设置全局通用的trace方法
func SetDefaultHandler(j InjectHandler) {
	defaultHandler = j
}
