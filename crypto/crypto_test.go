package crypto_test

import (
	"fmt"
	"github.com/tianlin0/go-plat-utils/crypto"
	"github.com/tianlin0/go-plat-utils/utils"
	"testing"
)

func TestAesCbc(t *testing.T) {
	key := "jasonsjiang29121"
	testCases := []*utils.TestStruct{
		{"空字符串", []any{""}, []any{"d41d8cd98f00b204e9800998ecf8427e"}, crypto.Md5},
		{"字符串", []any{"test"}, []any{"098f6bcd4621d373cade4e832627b4f6"}, crypto.Md5},
		{"带空格的字符串", []any{"hello world"}, []any{"5eb63bbbe01eeed093cb22bb8f5acdc3"}, crypto.Md5},
		{"CBCEncrypt", []any{"tianlin0"}, []any{true}, func(input string) bool {
			//每次都会不一样
			enStr, err := crypto.CBCEncrypt(input, key)
			if err != nil {
				return false
			}
			oldStr, err := crypto.CBCDecrypt(enStr, key)
			if err != nil {
				return false
			}
			if oldStr == input {
				return true
			}
			return false
		}},
		{"AesEncryptBase64", []any{"tianlin0"}, []any{true}, func(input string) bool {
			//每次都会不一样
			enStr, err := crypto.AesEncryptBase64(input, key)
			if err != nil {
				return false
			}
			oldStr, err := crypto.AesDecryptBase64(enStr, key)
			if err != nil {
				return false
			}
			if oldStr == input {
				return true
			}
			return false
		}},
		{"Argon2PasswordEncoder", []any{"tianlin0"}, []any{true}, func(input string) bool {
			//每次都会不一样
			enStr, err := crypto.Argon2PasswordEncoder(input)
			if err != nil {
				return false
			}
			fmt.Println("Argon2PasswordEncoder:", input, enStr)

			has, err := crypto.Argon2CompareHashAndPassword(enStr, input)
			if err != nil {
				return false
			}

			return has
		}},
	}
	utils.TestFunction(t, testCases, nil)
}
