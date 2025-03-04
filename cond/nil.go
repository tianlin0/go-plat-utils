package cond

import (
	"reflect"
)

// IsNil 判断是否为空
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	vi := reflect.ValueOf(i)
	kind := vi.Kind()
	if kind == reflect.Ptr ||
		kind == reflect.Chan ||
		kind == reflect.Func ||
		kind == reflect.UnsafePointer ||
		kind == reflect.Map ||
		kind == reflect.Interface ||
		kind == reflect.Slice {
		return vi.IsNil()
	}
	return false
}

// IsZero 判断变量是否为零值
func IsZero(v interface{}) bool {
	return reflect.ValueOf(v).IsZero()
}
