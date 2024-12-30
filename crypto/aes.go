package crypto

import "encoding/base64"

// AesEncrypt 对字符串对称加密，输入输出都为base64字符
func AesEncrypt(origString string, keyBase64 string, en Encoder) (string, error) {
	if origString == "" {
		return "", nil
	}
	origData := []byte(origString)

	key := getAesKeyFromBase64(keyBase64)
	iv := ""

	encrypted, err := aesEncrypt(origData, key, iv)
	if err != nil {
		return "", err
	}
	if en == nil {
		en = base64.StdEncoding.EncodeToString
	}
	return en(encrypted), nil
}

// AesDecrypt 对字符串对称解密，输入输出都为base64字符
func AesDecrypt(encodeString string, keyBase64 string, de Decoder) (string, error) {
	if encodeString == "" {
		return "", nil
	}
	key := getAesKeyFromBase64(keyBase64)
	iv := ""

	if de == nil {
		de = base64.StdEncoding.DecodeString
	}
	encodeByte, err := de(encodeString)
	if err != nil {
		return "", err
	}

	decodeByte, err := aesDecrypt(encodeByte, key, iv)
	if err != nil {
		return "", err
	}
	return string(decodeByte), nil
}
