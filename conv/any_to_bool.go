package conv

import (
	"reflect"
	"strings"
)

// Bool 将给定的值转换为bool
func Bool(i any) (bool, bool) {
	if i == nil {
		return false, false
	}
	if b, ok := i.(bool); ok {
		return b, true
	}
	if b, ok := Int64(i); ok {
		if b == 0 {
			return false, true
		}
		return true, true
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return false, false
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() != 0, true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() != 0, true
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0, true
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() != 0, true
	case reflect.String:
		val := strings.ToLower(v.String())
		if val == "true" {
			return true, true
		} else if val == "false" {
			return false, true
		}
		return v.String() != "", false
	default:
		val := strings.ToLower(String(i))
		if val == "true" {
			return true, true
		} else if val == "false" {
			return false, true
		}
		return false, false
	}
}
