package tools

import (
	"github.com/stretchr/testify/assert"
	"github.com/tianlin0/go-plat-utils/conv"
	"testing"
)

type TestStruct struct {
	Name     string
	Inputs   []interface{}
	Expected []interface{}
}

// TestFunction 统一检查函数
func TestFunction(t *testing.T, paramList []*TestStruct, checkFunc any) {
	for _, tc := range paramList {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := CallFunction(checkFunc, tc.Inputs...)
			funcName, _ := GetFunctionName(checkFunc)
			if err != nil {
				t.Errorf("[%s] funcName: %s, error: (%v), Input: (%v)", tc.Name, funcName, err, conv.String(tc.Inputs))
				return
			}
			//比较最小的内容，有可能有后面返回error，不写出来的情况
			//也有可能以前写的方法返回3个值，后面改为2个
			minNum := len(tc.Expected)
			if len(tc.Expected) > len(result) {
				minNum = len(result)
			}

			for i := 0; i < minNum; i++ {
				assert.Equal(t, tc.Expected[i], result[i])
			}
		})
	}
}
