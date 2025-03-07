package ruleengine

import (
	"github.com/shopspring/decimal"
)

// customerFunc 自定义方法列表
type customerFunc struct {
}

func (r *customerFunc) getAllDecimalList(args ...interface{}) []decimal.Decimal {
	decimalList := make([]decimal.Decimal, 0)
	for _, arg := range args {
		var d decimal.Decimal
		switch v := arg.(type) {
		case float64:
			d = decimal.NewFromFloat(v)
		case float32:
			d = decimal.NewFromFloat32(v)
		case int:
			d = decimal.NewFromInt(int64(v))
		case int64:
			d = decimal.NewFromInt(v)
		case int32:
			d = decimal.NewFromInt32(v)
		case uint:
			d = decimal.NewFromUint64(uint64(v))
		case uint32:
			d = decimal.NewFromUint64(uint64(v))
		case uint64:
			d = decimal.NewFromUint64(v)
		}
		if !d.IsZero() {
			decimalList = append(decimalList, d)
		}
	}
	return decimalList
}

// relationByNumber 两数相加
func (r *customerFunc) relationByNumber(f func(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal, args ...interface{}) float64 {
	decimalList := r.getAllDecimalList(args...)
	if len(decimalList) == 0 {
		return 0
	}
	var total decimal.Decimal
	for i, d := range decimalList {
		if i == 0 {
			total = d
			continue
		}
		total = f(total, d)
	}
	ret, _ := total.Float64()
	return ret
}

// AddByNumber 两数相加
func (r *customerFunc) AddByNumber(args ...interface{}) (interface{}, error) {
	return r.relationByNumber(func(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
		return d1.Add(d2)
	}, args...), nil
}

// SubByNumber 两数相减
func (r *customerFunc) SubByNumber(args ...interface{}) (interface{}, error) {
	return r.relationByNumber(func(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
		return d1.Sub(d2)
	}, args...), nil
}

// MulByNumber 两数相乘
func (r *customerFunc) MulByNumber(args ...interface{}) (interface{}, error) {
	return r.relationByNumber(func(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
		return d1.Mul(d2)
	}, args...), nil
}

// DivByNumber 两数相除
func (r *customerFunc) DivByNumber(args ...interface{}) (interface{}, error) {
	return r.relationByNumber(func(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
		return d1.Div(d2)
	}, args...), nil
}
