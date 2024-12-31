package curl

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/tianlin0/go-plat-utils/cache"
	"github.com/tianlin0/go-plat-utils/conv"
	"github.com/tianlin0/go-plat-utils/goroutines"
	"time"
)

const (
	ns = "comm-request"
)

// responseCacheStruct 返回的缓存结构
type responseCacheStruct struct {
	Time     time.Time `json:"time"`
	Response string    `json:"response"`
}

func getDataFromCache(p *genRequest) string {
	if p.Cache == 0 || p.cacheInstance == nil {
		return ""
	}
	cacheId := getRequestId(p.getRequest())

	retData, err := cache.NsGet(p.cacheInstance, ns, cacheId)
	if err != nil || retData == "" {
		return ""
	}

	cacheData := new(responseCacheStruct)
	err = jsoniter.Unmarshal([]byte(retData), cacheData)
	if err != nil {
		_, _ = cache.NsDel(p.cacheInstance, ns, cacheId)
		return ""
	}
	//超时
	if time.Now().Sub(cacheData.Time) > p.Cache {
		_, _ = cache.NsDel(p.cacheInstance, ns, cacheId)
		return ""
	}

	findData := false
	if p.retryPolicy != nil {
		retOk, err := p.retryPolicy.onlyCheckCondition(cacheData.Response)
		if err == nil && !retOk {
			findData = true
		}
	}

	if findData {
		return cacheData.Response
	}

	return ""
}

func setDataToCache(g *genRequest, p *Response, cacheTime time.Duration) {
	if g.Cache == 0 || g.cacheInstance == nil {
		return
	}
	cacheId := getRequestId(p.Request)

	goroutines.GoAsync(func(params ...interface{}) {
		cacheData := responseCacheStruct{
			Time:     time.Now(),
			Response: p.Response,
		}
		cacheStr := conv.String(cacheData)
		if cacheStr == "" {
			return
		}

		_, _ = cache.NsSet(g.cacheInstance, ns, cacheId, cacheStr, cacheTime)
	}, nil)
}
