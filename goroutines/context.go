package goroutines

import (
	"context"
	gocache "github.com/patrickmn/go-cache"
	"github.com/timandy/routine"
	"strconv"
	"sync"
	"time"
)

var (
	expiration      = 20 * time.Minute //一个ctx的最长时间，避免长时间占用内存
	cleanupInterval = 30 * time.Minute
	baseInt         = 10
	onceCache       = sync.Once{}
	ctxCache        *gocache.Cache
	localStorage    routine.ThreadLocal[string]
)

// getCache 初始化并获得缓存
func getCache() *gocache.Cache {
	if ctxCache == nil {
		onceCache.Do(func() {
			ctxCache = gocache.New(expiration, cleanupInterval)
			localStorage = routine.NewInheritableThreadLocal[string]()
		})
	}
	return ctxCache
}

// getCurrentGoId 取得当前的协程ID
func getCurrentGoId() string {
	return strconv.FormatInt(routine.Goid(), baseInt)
}

// SetContext 设置上下文，需要在入口协程上执行，会返回当前的IdKey
func SetContext(ctx *context.Context) {
	_, ctxKey := GetContext()
	if ctxKey == "" {
		ctxKey = getCurrentGoId()
	}
	ctxFactory := getCache()
	ctxFactory.Set(ctxKey, ctx, gocache.DefaultExpiration)
	localStorage.Set(ctxKey)
}

// GetContext 获取上下文
func GetContext() (ctx *context.Context, ctxKey string) {
	ctxFactory := getCache()

	ctxKey = localStorage.Get()
	if ctxKey == "" {
		return nil, ""
	}

	val, ok := ctxFactory.Get(ctxKey)
	if !ok {
		return nil, ctxKey
	}
	ctx, ok = val.(*context.Context)
	if ok {
		return ctx, ctxKey
	}
	return nil, ctxKey
}

// DelContext 删除上下文
func DelContext() {
	ctxFactory := getCache()
	ctxKey := localStorage.Get()
	if ctxKey == "" {
		return
	}
	ctxFactory.Delete(ctxKey)
}
