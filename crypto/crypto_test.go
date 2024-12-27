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
		{"BCryptPasswordEncoder", []any{"tianlin0"}, []any{true}, func(input string) bool {
			//每次都会不一样
			enStr, err := crypto.BCryptPasswordEncoder(input)
			if err != nil {
				return false
			}
			fmt.Println("BCryptPasswordEncoder:", input, enStr)

			has, err := crypto.BCryptCompareHashAndPassword(enStr, input)
			if err != nil {
				return false
			}

			return has
		}},
		{"GobEncode", []any{"tianlin0"}, []any{"tianlin0"}, func(input string) string {
			//每次都会不一样
			enStr, err := crypto.GobEncode(input)
			if err != nil {
				return ""
			}
			fmt.Println("GobEncode:", input, enStr)

			aa := ""
			err = crypto.GobDecode(enStr, &aa)
			if err != nil {
				return ""
			}

			return aa
		}},
		{"XorEncode", []any{"tianlin0"}, []any{"tianlin0"}, func(input string) string {
			keyInt := 5
			enStr := crypto.XorEncode(input, keyInt)

			fmt.Println("XorEncode:", input, enStr)

			return crypto.XorDecode(enStr, keyInt)
		}},
	}
	utils.TestFunction(t, testCases, nil)
}
