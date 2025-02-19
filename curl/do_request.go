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

const (
	//默认超时
	defaultConnectTimeoutSecond  = 20 * time.Second
	defaultTimeoutSecond         = 30 * time.Second
	defaultMaxCoons              = 100
	defaultIdleConnTimeoutSecond = 600 * time.Second
)

// createDefaultTransport 创建默认的 http.Transport 配置
func createDefaultTransport() *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   defaultConnectTimeoutSecond, // 执行连接超时，表示建立连接的最大超时时间。如果在 defaultConnectTimeoutSecond 时间内未能成功建立与目标服务器的连接，连接操作将失败并返回错误
			KeepAlive: defaultTimeoutSecond,        // 长连接保持时间，指定了 TCP 连接的长连接保持时间。在连接建立后，如果在 defaultTimeoutSecond 时间内没有数据传输，TCP 协议会自动发送一个保活探测包来检测连接是否仍然有效。如果超过这个时间没有收到响应，连接可能会被关闭
		}).DialContext,
		DisableKeepAlives:   true,                         // 关闭长连接，当设置为 true 时，会禁用 HTTP 长连接。长连接允许在同一个 TCP 连接上进行多次 HTTP 请求和响应，避免了每次请求都重新建立连接的开销。如果将 DisableKeepAlives 设置为 true，那么每次 HTTP 请求都会建立一个新的 TCP 连接，请求完成后立即关闭该连接。
		MaxIdleConnsPerHost: -1,                           // 该属性指定了每个目标主机（IP 地址和端口的组合）允许保持的最大空闲连接数。当设置为 -1 时，实际上是禁用了空闲连接的复用，结合 DisableKeepAlives: true，进一步确保不会使用长连接。如果设置为一个正整数 n，那么对于每个目标主机，最多会保留 n 个空闲连接以备后续使用。
		MaxIdleConns:        defaultMaxCoons,              // 表示整个 http.Transport 实例允许保持的最大空闲连接总数。无论目标主机有多少个，所有空闲连接的总数不会超过 defaultMaxCoons。例如，如果 defaultMaxCoons 被设置为 100，那么当空闲连接数达到 100 时，新的空闲连接将不再保留，会被关闭。
		IdleConnTimeout:     defaultIdleConnTimeoutSecond, // 指定了空闲连接的最大保持时间。如果一个连接在 defaultIdleConnTimeoutSecond 时间内一直处于空闲状态（没有数据传输），那么该连接将被关闭。这有助于释放不再使用的连接资源，避免资源浪费。
		DisableCompression:  true,                         //当设置为 true 时，会禁用 HTTP 响应的压缩功能。
	}
}

// 默认使用短连接，长连接没有弄好的前提下
func createHTTPClient(timeout time.Duration, trans *http.Transport) *http.Client {
	tr := createDefaultTransport()

	// 如果传入的传输配置不为空，则合并配置
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
	return &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
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

	defer func(body io.ReadCloser) {
		_ = body.Close()
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
