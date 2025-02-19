package curl

import (
	"context"
	"github.com/tianlin0/go-plat-utils/cache"
	"github.com/tianlin0/go-plat-utils/logs"
	"net/http"
	"time"
)

const (
	PrintOne = iota
	PrintAll
	PrintNone
)

// SetDefaultPrintType PrintOne只会默认打一条，PrintAll全打，PrintNone不打
func (g *genRequest) SetDefaultPrintType(b int) *genRequest {
	if b == PrintOne || b == PrintNone || b == PrintAll {
		g.defaultPrintLogInt = b
	}
	return g
}
func (g *genRequest) SetHttpTransport(t *http.Transport) *genRequest {
	g.transportConfig = t
	return g
}
func (g *genRequest) SetLogger(l logs.ILogger) *genRequest {
	g.logger = l
	return g
}
func (g *genRequest) SetHandler(b InjectHandler) *genRequest {
	g.handler = b
	return g
}

func (g *genRequest) SetUrl(s string) *genRequest {
	g.Url = s
	return g
}
func (g *genRequest) WithContext(ctx context.Context) *genRequest {
	g.ctx = ctx
	return g
}
func (g *genRequest) SetData(d interface{}) *genRequest {
	g.Data = d
	return g
}
func (g *genRequest) SetMethod(m string) *genRequest {
	g.Method = m
	return g
}
func (g *genRequest) SetHeader(h http.Header) *genRequest {
	g.Header = h
	return g
}
func (g *genRequest) SetTimeout(t time.Duration) *genRequest {
	g.Timeout = t
	return g
}
func (g *genRequest) SetCache(c cache.CommCache[string], t time.Duration) *genRequest {
	g.Cache = t
	g.cacheInstance = c
	return g
}
func (g *genRequest) SetRetryPolicy(p *RetryPolicy) *genRequest {
	if p == nil {
		g.retryPolicy = nil //去掉重试条件
		return g
	}

	if g.retryPolicy == nil {
		g.retryPolicy = p
	}
	if p.MaxAttempts > 0 {
		g.setRetryTimes(p.MaxAttempts)
	}
	if p.RetryConditionFunc != nil {
		g.retryPolicy.RetryConditionFunc = p.RetryConditionFunc
	}
	if p.RetryCondition != "" {
		g.setRetryCondition(p.RetryCondition)
	}
	g.retryPolicy.RespDateType = p.RespDateType
	return g
}
func (g *genRequest) setRetryCondition(c string) *genRequest {
	if g.retryPolicy == nil {
		g.retryPolicy = new(RetryPolicy)
	}
	g.retryPolicy.RetryCondition = c
	return g
}
func (g *genRequest) setRetryTimes(t int) *genRequest {
	if g.retryPolicy == nil {
		g.retryPolicy = new(RetryPolicy)
	}
	if t >= 0 {
		g.retryPolicy.MaxAttempts = t
		g.retryPolicy.leftAttempts = t
	}
	return g
}
