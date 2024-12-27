package conv_test

import (
	"github.com/tianlin0/go-plat-utils/conv"
	"testing"
)

func TestAnyToBool(t *testing.T) {
	iPtr := 90
	testCases := []*utils.TestStruct{
		{"bool true", []any{true}, []any{true, true}},
		{"bool false", []any{false}, []any{false}},
		{"int -1", []any{int(-1)}, []any{true}},
		{"int 1", []any{int(1)}, []any{true}},
		{"int 0", []any{int(0)}, []any{false}},
		{"int8 1", []any{int8(1)}, []any{true}},
		{"int8 0", []any{int8(0)}, []any{false}},
		{"int16 1", []any{int16(1)}, []any{true}},
		{"int16 0", []any{int16(0)}, []any{false}},
		{"int32 1", []any{int32(1)}, []any{true}},
		{"int32 0", []any{int32(0)}, []any{false}},
		{"int64 1", []any{int64(1)}, []any{true}},
		{"int64 100", []any{int64(100)}, []any{true}},
		{"int64 0", []any{int64(0)}, []any{false}},
		{"uint 1", []any{uint(1)}, []any{true}},
		{"uint 0", []any{uint(0)}, []any{false}},
		{"uint8 1", []any{uint8(1)}, []any{true}},
		{"uint8 0", []any{uint8(0)}, []any{false}},
		{"uint16 1", []any{uint16(1)}, []any{true}},
		{"uint16 0", []any{uint16(0)}, []any{false}},
		{"uint32 1", []any{uint32(1)}, []any{true}},
		{"uint32 0", []any{uint32(0)}, []any{false}},
		{"uint64 1", []any{uint64(1)}, []any{true}},
		{"uint64 0", []any{uint64(0)}, []any{false}},
		{"float32 1.0", []any{float32(1.0)}, []any{true}},
		{"float32 0.0", []any{float32(0.0)}, []any{false}},
		{"float64 1.0", []any{float64(1.0)}, []any{true}},
		{"float64 0.0", []any{float64(0.0)}, []any{false}},
		{"string abc", []any{"abc"}, []any{true}},
		{"string true", []any{"true"}, []any{true}},
		{"string false", []any{"false"}, []any{false}},
		{"empty string", []any{""}, []any{false}},
		{"nil value", []any{nil}, []any{false}},
		{"complex64 1+1i", []any{complex64(1 + 1i)}, []any{true}},
		{"complex64 0+0i", []any{complex64(0 + 0i)}, []any{false}},
		{"complex128 1+1i", []any{complex128(1 + 1i)}, []any{true}},
		{"complex128 0+0i", []any{complex128(0 + 0i)}, []any{false}},
		{"nil pointer", []any{(*int)(nil)}, []any{false}},
		{"non-nil pointer", []any{&iPtr}, []any{true}},
		{"empty slice", []any{[]int{}}, []any{false}},
	}
	utils.TestFunction(t, testCases, conv.Bool)
}
