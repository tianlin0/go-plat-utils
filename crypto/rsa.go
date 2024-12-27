package crypto

// EncryptRSA RSA加密数据，key必须是成对出现
func EncryptRSA(oneKeyStr string, message string) (string, error) {
	rsa := new(RSASecurity)
	err := rsa.SetPublicAndPrivateKey(oneKeyStr, "")
	if err != nil {
		err = rsa.SetPublicAndPrivateKey("", oneKeyStr)
		if err != nil {
			return "", err
		}
		return rsa.PriKeyEncryptBase64(message)
	}
	return rsa.PubKeyEncryptBase64(message)
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
		return rsa.PubKeyDecryptBase64(cipherText)
	}
	return rsa.PriKeyDecryptBase64(cipherText)
}
