package curl

import (
	"context"
	"net/http"
)

type InjectHandler interface {
	// BeforeHandler 发送前的方法
	BeforeHandler(ctx context.Context, rs *Request, httpReq *http.Request) error
	// AfterHandler 发送后的方法
	AfterHandler(ctx context.Context, rp *Response) error
}

type client struct {
	handler InjectHandler
}

// NewClient 客户端
func NewClient() *client {
	return new(client)
}

// WithHandler 设置执行前方法
func (c *client) WithHandler(h InjectHandler) *client {
	c.handler = h
	return c
}

func (c *client) NewRequest(r *Request) *genRequest {
	gen := newRequest(r)
	gen.SetHandler(c.handler)
	return gen
}
