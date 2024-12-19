package conf

import (
	"context"
	gocache "github.com/patrickmn/go-cache"
	"github.com/timandy/routine"
	"strconv"
	"sync"
	"time"
)

var (
	expiration      = 20 * time.Minute
	cleanupInterval = 30 * time.Minute
	baseInt         = 10
	once            = sync.Once{}
)

var ctxCache *gocache.Cache
var localStorage routine.ThreadLocal[string]

// getCache 初始化并获得缓存
func getCache() *gocache.Cache {
	if ctxCache == nil {
		once.Do(func() {
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

// SetContext 设置上下文
func SetContext(ctx *context.Context, goIds ...string) {
	ctxFactory := getCache()
	goId := getCurrentGoId()
	ctxFactory.Set(goId, ctx, gocache.DefaultExpiration)
	localStorage.Set(goId)
}

// GetContext 获取上下文
func GetContext() (ctx *context.Context, goId string, b bool) {
	ctxFactory := getCache()

	goId = localStorage.Get()
	if goId == "" {
		return nil, "", false
	}

	val, ok := ctxFactory.Get(goId)
	if !ok {
		return nil, goId, false
	}
	ctx, ok = val.(*context.Context)
	if ok {
		return ctx, goId, true
	}
	return nil, goId, false
}

// DelContext 删除上下文
func DelContext() {
	ctxFactory := getCache()
	goId := localStorage.Get()
	if goId == "" {
		return
	}
	ctxFactory.Delete(goId)
}
