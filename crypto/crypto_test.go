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
		{"AESEncrypt", []any{"tianlin0"}, []any{true}, func(input string) bool {

			appSecret := "IgkibX71IEf382PT"

			enStr, err := crypto.AESEncrypt(input, []byte(appSecret), appSecret)
			if err != nil {
				return false
			}
			oldStr, err := crypto.AESDecrypt(enStr, []byte(appSecret), appSecret)
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

func TestCreateRSAKeys(t *testing.T) {
	rsa := new(crypto.RSASecurity)
	pub, pri, err := rsa.CreateRSAKeys(0, crypto.PKCS8Type)

	fmt.Println(pub)
	fmt.Println(pri)
	fmt.Println(err)

	rsa.SetPublicAndPrivateKey(
		`-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANQjZeh61+XSP/XgQLv0EeWJylEsTQiV
r/wUSB2pY7XZyEIJ7czGbl3xe5nXxEMHjlgkWH14VzS3+cPIejrraBECAwEAAQ==
-----END PUBLIC KEY-----`,
		`-----BEGIN PRIVATE KEY-----
MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEA1CNl6HrX5dI/9eBA
u/QR5YnKUSxNCJWv/BRIHaljtdnIQgntzMZuXfF7mdfEQweOWCRYfXhXNLf5w8h6
OutoEQIDAQABAkAhPxb2k2IIq6XIhAfBLSQs5CZoCFheUw9Mo2UV+Pkeg6UW2DHU
TB1N4DFKlCpGpsYHaPgmpRqEilBtdslcNzJNAiEA7ZSBDicBSwChAdCHBKF6mTSh
fWgCfcBmItqNLrNrJhMCIQDkleyOFQjeOVxcTm0RA6cC/f/wJ5lc4++c16eL9Y3N
ywIgUmVDkO30I9f2/xMcEH4Ub9fx/fU5j/VPNt1HQ6AUFCMCIQDXhR/Pare8xrp1
caBV3Wq3YILSjJOFyIdgCti3FmOH9wIgMT5Fzig0Qsp47jwJ0ICRZCaXFYA0XKPI
TjWDiQ4P6p8=
-----END PRIVATE KEY-----
	`)

	kk, err := rsa.PubKeyEncryptBase64("tiantian")
	fmt.Println(kk, err)

	mm, err := rsa.PriKeyDecryptBase64(kk)
	fmt.Println(mm, err)

	nn, err := rsa.SignMd5WithRsa("abc")
	fmt.Println(nn, err)
	err = rsa.VerifySignMd5WithRsa("abc", nn)
	fmt.Println(err)

}
