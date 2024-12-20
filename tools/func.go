package tools

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// GetFunctionName 获取方法名
func GetFunctionName(i any) (name string, isMethod bool) {
	if fullName, ok := i.(string); ok {
		return fullName, false
	}
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	isMethod = strings.ContainsAny(fullName, "*")
	elements := strings.Split(fullName, ".")
	shortName := elements[len(elements)-1]
	return strings.TrimSuffix(shortName, "-fm"), isMethod
}

// CallFunction 根据参数，可以调用任何方法
func CallFunction(function any, args ...any) (result []any, err error) {
	// 将空接口转换为reflect.Value
	fnValue := reflect.ValueOf(function)

	// 检查是否为可调用的函数
	if fnValue.Kind() != reflect.Func {
		return nil, fmt.Errorf("[CallFunction] not function")
	}

	// 将参数列表转换为reflect.Value切片
	var retValues []reflect.Value
	if len(args) > 0 {
		fnArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			fnArgs[i] = reflect.ValueOf(arg)
		}
		// 调用函数，捕获panic
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		// 调用函数并返回结果
		retValues = fnValue.Call(fnArgs)
	} else {
		retValues = fnValue.Call(nil)
	}

	fnType := reflect.TypeOf(function)
	numOut := fnType.NumOut()

	if numOut == 0 && len(retValues) == 0 {
		return []any{}, nil
	}

	if len(retValues) != numOut {
		return []any{}, fmt.Errorf("[CallFunction] ret error, "+
			"retValue number is %d, but Type number is %d", len(retValues), numOut)
	}

	//retTypes := make([]reflect.Type, numOut)
	//for i := 0; i < numOut; i++ {
	//	retTypes[i] = fnType.Out(i)
	//}

	result = make([]any, numOut)
	for i := 0; i < numOut; i++ {
		result[i] = retValues[i].Interface()
	}

	return result, nil
}
