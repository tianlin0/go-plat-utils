package crypto_test

import (
	"github.com/tianlin0/go-plat-utils/crypto"
	"testing"
)

func TestMd5(t *testing.T) {
	testCases := []*utils.TestStruct{
		{"空字符串", []any{""}, []any{"d41d8cd98f00b204e9800998ecf8427e"}},
		{"字符串", []any{"test"}, []any{"098f6bcd4621d373cade4e832627b4f6"}},
		{"带空格的字符串", []any{"hello world"}, []any{"5eb63bbbe01eeed093cb22bb8f5acdc3"}},
	}
	utils.TestFunction(t, testCases, crypto.Md5)
}
