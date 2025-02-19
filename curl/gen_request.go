package curl

import (
	"bytes"
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tianlin0/go-plat-utils/cache"
	"github.com/tianlin0/go-plat-utils/conv"
	"github.com/tianlin0/go-plat-utils/logs"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type genRequest struct {
	Request
	retryPolicy *RetryPolicy

	defaultPrintLogInt int //0 表示默认，只打印一条，1表示完全打开所有信息，-1 表示完全关闭
	transportConfig    *http.Transport
	logger             logs.ILogger
	handler            InjectHandler
	ctx                context.Context
	cacheInstance      cache.CommCache[string]
}

func (g *genRequest) buildGenRequest() {
	if g.Data == nil {
		g.Data = ""
	}
	g.Method = getMethod(g.Method)

	if g.Timeout <= 0 {
		g.Timeout = defaultTimeoutSecond
	}

	if g.Cache > 0 {
		if g.Cache > defaultMaxCacheTime {
			g.Cache = defaultMaxCacheTime
		}
	}

	g.Url = strings.TrimSpace(g.Url)
	g.Header = getHeaders(g.Header, g.Method, g.Data)

	if g.handler == nil {
		g.handler = defaultHandler
	}

	if g.retryPolicy == nil {
		g.retryPolicy = new(RetryPolicy)
	}
	g.retryPolicy.buildRetryable()

}

func (g *genRequest) getRequest() *Request {
	req := new(Request)
	req.Url = g.Url
	req.Data = g.Data
	req.Method = g.Method
	req.Header = g.Header
	req.Timeout = g.Timeout
	req.Cache = g.Cache
	return req
}

func (g *genRequest) Submit(ctx context.Context) *Response {
	g.buildGenRequest()

	resp := NewResponse(g.getRequest())

	dataString, err := getDataString(g.Data)
	if err != nil {
		resp.Error = err
		return resp
	}

	if g.Url == "" {
		resp.Error = fmt.Errorf("url请求地址为空")
		return resp
	}

	_, err = url.Parse(g.Url)
	if err != nil {
		resp.Error = fmt.Errorf("url格式错误：%s, %v", g.Url, err)
		return resp
	}

	if ctx == nil {
		ctx = g.ctx
	}

	if ctx == nil {
		ctx = context.Background()
	}

	startTime := time.Now()
	if g.Cache > 0 {
		respTxt := getDataFromCache(ctx, g)
		if respTxt != "" {
			resp.Response = respTxt
			resp.fromCache = true
			resp.SetDuration(startTime)

			simpleCurlStrData := resp.getLoggerResponse(startTime)
			logStr := fmt.Sprintf("[comm-request cache return]id:%s, data:%s",
				resp.Id, conv.String(simpleCurlStrData))
			printLog(ctx, g.logger, g.defaultPrintLogInt, logStr)

			return resp
		}
	}

	postUrl := getNewPostUrl(g.Url, g.Method, dataString)

	logStr := fmt.Sprintf("[comm-request request] url:%s", postUrl)
	printLog(ctx, g.logger, g.defaultPrintLogInt, logStr)

	startTime = time.Now()
	allResp := g.httpRequest(ctx, dataString, resp)
	allResp.SetDuration(startTime)

	//对返回值进行检查
	if g.retryPolicy != nil {
		isRetry, err := g.retryPolicy.onlyCheckCondition(allResp.Response)
		if err != nil {
			allResp.Error = err
		} else {
			if isRetry {
				allResp.Error = fmt.Errorf(allResp.Response)
			}
		}
	}

	simpleCurlStrData := resp.getLoggerResponse(startTime)
	if g.defaultPrintLogInt == PrintOne || g.defaultPrintLogInt == PrintAll {
		simpleCurlStrData.printLoggerResponse(ctx, g.logger)
	}

	if allResp.HttpStatus == http.StatusOK &&
		allResp.Error == nil &&
		allResp.Response != "" &&
		g.Cache > 0 {
		setDataToCache(ctx, g, allResp, g.Cache)
	}

	return allResp
}

func (g *genRequest) getHttpRequest(ctx context.Context, dataString string) (*http.Request, error) {
	postUrl := getNewPostUrl(g.Url, g.Method, dataString)

	httpReq, err := http.NewRequest(g.Method, postUrl, bytes.NewBufferString(dataString))
	if err != nil {
		logStr := fmt.Sprintf("[comm-request request] url:%s, error: %s", postUrl, err.Error())
		printLog(ctx, g.logger, g.defaultPrintLogInt, logStr)
		return nil, err
	}

	if len(g.Header) > 0 {
		for k, v := range g.Header {
			httpReq.Header = setHeaderValues(httpReq.Header, k, v...)
		}
	}

	return httpReq, nil
}

// 递归使用
func (g *genRequest) httpRequest(ctx context.Context, dataString string, resp *Response) *Response {
	httpReq, err := g.getHttpRequest(ctx, dataString)
	if err != nil {
		resp.Error = err
		return resp
	}
	resp.Error = nil

	newRequest := g.getRequest()
	if g.handler != nil {
		err = g.handler.BeforeHandler(ctx, newRequest, httpReq)
		if err != nil {
			resp.Error = err
			return resp
		}
	}

	bodyIsJson := true
	if g.retryPolicy != nil {
		if g.retryPolicy.RespDateType == "string" {
			bodyIsJson = false
		}
	}

	startTime := time.Now()
	{ //直接请求
		status, resData, resHeader, err := doRequest(httpReq, newRequest, g.logger, g.Timeout, g.transportConfig)
		resp.Response = resData
		resp.Request = newRequest
		resp.HttpStatus = status
		resp.Header = resHeader
		resp.Error = err
		resp.SetDuration(startTime)
	}

	if bodyIsJson {
		if resp.Error == nil {
			var obj interface{}
			err = jsoniter.Unmarshal([]byte(resp.Response), &obj)
			if err != nil {
				//返回的不是json格式
				resp.Error = fmt.Errorf("url: %s, response not json: %s，Request=>SetRetryPolicy=>RespDateType=string", newRequest.Url, resp.Response)
			}
		}
	}

	simpleCurlStrData := resp.getLoggerResponse(startTime)
	logStr := fmt.Sprintf("[comm-request http-request return]id:%s, data:%s, error:%v", simpleCurlStrData.Id,
		conv.String(simpleCurlStrData), resp.Error)
	printLog(ctx, g.logger, g.defaultPrintLogInt, logStr)

	if g.handler != nil {
		err = g.handler.AfterHandler(ctx, resp)
		if err != nil {
			resp.Error = err
			return resp
		}
	}

	//如果有返回，则判断是否成功。
	isRetry := canRetry(g.retryPolicy, resp.Response)
	if isRetry {
		return g.httpRequest(ctx, dataString, resp)
	}
	return resp
}
