package curl

import (
	"net/http"
	"time"
)

// Request 请求变量
type Request struct {
	Url     string        `json:"url"`
	Data    interface{}   `json:"data"`
	Method  string        `json:"method"`
	Header  http.Header   `json:"header"`
	Timeout time.Duration `json:"timeout,omitempty"`
	Cache   time.Duration `json:"cache,omitempty"`
}

func newRequest(r *Request) *genRequest {
	nr := new(genRequest)
	nr.Url = r.Url
	nr.Data = r.Data
	nr.Method = r.Method
	nr.Header = r.Header
	nr.Timeout = r.Timeout
	nr.Cache = r.Cache
	return nr
}
