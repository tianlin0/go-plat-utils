package curl

import (
	"fmt"
	"github.com/tianlin0/go-plat-utils/logs"
	"github.com/tianlin0/go-plat-utils/utils"
	"io"
	"net"
	"net/http"
	"time"
)

var (
	//默认超时
	defaultConnectTimeoutSecond  = 20 * time.Second
	defaultTimeoutSecond         = 30 * time.Second
	defaultMaxCoons              = 100
	defaultIdleConnTimeoutSecond = 600 * time.Second
)

// 默认使用短连接，长连接没有弄好的前提下
func createHTTPClient(timeout time.Duration, trans *http.Transport) *http.Client {
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   defaultConnectTimeoutSecond, // 执行连接超时
			KeepAlive: defaultTimeoutSecond,        // 长连接保持时间
		}).DialContext,
		DisableKeepAlives:   true,                         //关闭长连接
		MaxIdleConnsPerHost: -1,                           //关闭长连接
		MaxIdleConns:        defaultMaxCoons,              // 长连接个数
		IdleConnTimeout:     defaultIdleConnTimeoutSecond, // 长连接保持时间
		DisableCompression:  true,
	}
	if trans != nil {
		if trans.MaxIdleConns == 0 {
			trans.MaxIdleConns = tr.MaxIdleConns
		}
		if trans.IdleConnTimeout == 0 {
			trans.IdleConnTimeout = tr.IdleConnTimeout
		}
		if trans.DialContext == nil {
			trans.DialContext = tr.DialContext
		}
		tr = trans
	}
	client := &http.Client{Transport: tr, Timeout: timeout}
	return client
}

// doRequest 发起实际的HTTP请求
func doRequest(req *http.Request, reqTemp *Request, cLogger logs.ILogger, timeout time.Duration, trans *http.Transport) (status int,
	resData string, resHeader http.Header, err error) {
	timeStart := time.Now()

	httpClient := createHTTPClient(timeout, trans)
	req.Close = true

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}

	status = resp.StatusCode
	if len(resp.Header) > 0 {
		resHeader = resp.Header
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resData = string(body)

	if status != http.StatusOK {
		curlTime := utils.GetSinceMilliTime(timeStart)
		msg := fmt.Sprintf("[comm-request do-request][%dms] url=%s, StatusCode=%d, resData=%s, resHeader=%v",
			curlTime, reqTemp.Url, status, resData, resHeader)
		if cLogger != nil {
			cLogger.Error(msg) //业务逻辑的错，不能算错误
		}
	}
	return status, resData, resHeader, err
}
