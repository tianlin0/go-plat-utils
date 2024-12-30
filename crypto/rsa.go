package crypto

import "encoding/base64"

// EncryptRSA RSA加密数据，key必须是成对出现
func EncryptRSA(oneKeyStr string, message string) (string, error) {
	rsa := new(RSASecurity)
	err := rsa.SetPublicAndPrivateKey(oneKeyStr, "")
	if err != nil {
		err = rsa.SetPublicAndPrivateKey("", oneKeyStr)
		if err != nil {
			return "", err
		}
		return rsa.PriKeyEncrypt(message, base64.StdEncoding.EncodeToString)
	}
	return rsa.PubKeyEncrypt(message, base64.StdEncoding.EncodeToString)
}

// DecryptRSA RSA解密数据，key必须是成对出现
func DecryptRSA(otherKeyStr string, cipherText string) (string, error) {
	rsa := new(RSASecurity)
	err := rsa.SetPublicAndPrivateKey("", otherKeyStr)
	if err != nil {
		err = rsa.SetPublicAndPrivateKey(otherKeyStr, "")
		if err != nil {
			return "", err
		}
		return rsa.PubKeyDecrypt(cipherText, base64.StdEncoding.DecodeString)
	}
	return rsa.PriKeyDecrypt(cipherText, base64.StdEncoding.DecodeString)
}
