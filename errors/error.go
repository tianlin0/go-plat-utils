package errors

import (
	"errors"
	"github.com/tianlin0/go-plat-utils/conf"
)

// CommError 通用错误信息
type CommError interface {
	Error() string
	Code() int64
}

type commErr struct {
	code    int64  `json:"code"`
	message string `json:"message"`
}

// Error 错误信息返回，实现error接口
func (err *commErr) Error() string {
	if err == nil {
		return conf.EmptyString
	}
	return err.message
}
func (err *commErr) Code() int64 {
	if err == nil {
		return conf.DefaultErrorCode
	}
	return err.code
}

// New 新建错误对象
func New(msg string, code ...int64) *commErr {
	err := &commErr{
		code:    conf.DefaultErrorCode,
		message: msg,
	}
	if len(code) > 0 {
		err.code = code[0]
	}
	return err
}

// Wrap 新增error Code
func Wrap(err error, code ...int64) error {
	if err == nil {
		return nil
	}
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	var tempCode int64 = conf.DefaultErrorCode
	var errTemp CommError
	if errors.As(err, &errTemp) {
		tempCode = errTemp.Code()
	} else {
		if len(code) > 0 {
			tempCode = code[0]
		}
	}
	return New(errStr, tempCode)
}
