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
func (r *genRequest) SetDefaultPrintType(b int) *genRequest {
	if b == PrintOne || b == PrintNone || b == PrintAll {
		r.defaultPrintLogInt = b
	}
	return r
}
func (r *genRequest) SetHttpTransport(t *http.Transport) *genRequest {
	r.transportConfig = t
	return r
}
func (r *genRequest) SetLogger(l logs.ILogger) *genRequest {
	r.logger = l
	return r
}
func (r *genRequest) SetHandler(b InjectHandler) *genRequest {
	r.handler = b
	return r
}

func (r *genRequest) SetUrl(s string) *genRequest {
	r.Url = s
	return r
}
func (r *genRequest) WithContext(ctx context.Context) *genRequest {
	r.ctx = ctx
	return r
}
func (r *genRequest) SetData(d interface{}) *genRequest {
	r.Data = d
	return r
}
func (r *genRequest) SetMethod(m string) *genRequest {
	r.Method = m
	return r
}
func (r *genRequest) SetHeader(h http.Header) *genRequest {
	r.Header = h
	return r
}
func (r *genRequest) SetTimeout(t time.Duration) *genRequest {
	r.Timeout = t
	return r
}
func (r *genRequest) SetCache(c cache.CommCache, t time.Duration) *genRequest {
	r.Cache = t
	r.cacheInstance = c
	return r
}
func (r *genRequest) SetRetryPolicy(p *RetryPolicy) *genRequest {
	if p == nil {
		r.retryPolicy = nil //去掉重试条件
		return r
	}

	if r.retryPolicy == nil {
		r.retryPolicy = p
	}
	if p.MaxAttempts > 0 {
		r.setRetryTimes(p.MaxAttempts)
	}
	if p.RetryConditionFunc != nil {
		r.retryPolicy.RetryConditionFunc = p.RetryConditionFunc
	}
	if p.RetryCondition != "" {
		r.setRetryCondition(p.RetryCondition)
	}
	r.retryPolicy.RespDateType = p.RespDateType
	return r
}
func (r *genRequest) setRetryCondition(c string) *genRequest {
	if r.retryPolicy == nil {
		r.retryPolicy = new(RetryPolicy)
	}
	r.retryPolicy.RetryCondition = c
	return r
}
func (r *genRequest) setRetryTimes(t int) *genRequest {
	if r.retryPolicy == nil {
		r.retryPolicy = new(RetryPolicy)
	}
	if t >= 0 {
		r.retryPolicy.MaxAttempts = t
		r.retryPolicy.leftAttempts = t
	}
	return r
}
