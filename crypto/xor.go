package crypto

import (
	"encoding/base64"
)

// XorEncode 异或加密
func XorEncode(msg string, key int, en Encoder) string {
	byteList := []byte(msg)
	pwd := make([]byte, len(byteList))
	for i := 0; i < len(byteList); i++ {
		pwd[i] = byteList[i] ^ byte(key)
	}

	if en == nil {
		en = base64.StdEncoding.EncodeToString
	}

	return en(pwd)
}

// XorDecode 异或解密
func XorDecode(msg string, key int, de Decoder) (string, error) {
	if de == nil {
		de = base64.StdEncoding.DecodeString
	}
	pwdList, err := de(msg)
	if err != nil {
		return "", err
	}
	old := make([]byte, len(pwdList))
	for i := 0; i < len(pwdList); i++ {
		old[i] = pwdList[i] ^ byte(key)
	}
	return string(old), nil
}
