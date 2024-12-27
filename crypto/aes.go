package crypto

//// AESEncrypt aes
//func AESEncrypt(encryptStr string, key []byte, iv string) (string, error) {
//	encryptBytes := []byte(encryptStr)
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return "", err
//	}
//
//	blockSize := block.BlockSize()
//	encryptBytes = pKCS5Padding(encryptBytes, blockSize)
//
//	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
//	encrypted := make([]byte, len(encryptBytes))
//	blockMode.CryptBlocks(encrypted, encryptBytes)
//	return base64.URLEncoding.EncodeToString(encrypted), nil
//}
//
//// AESDecrypt aes
//func AESDecrypt(decryptStr string, key []byte, iv string) (string, error) {
//	decryptBytes, err := base64.URLEncoding.DecodeString(decryptStr)
//	if err != nil {
//		return "", err
//	}
//
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return "", err
//	}
//
//	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
//	decrypted := make([]byte, len(decryptBytes))
//
//	blockMode.CryptBlocks(decrypted, decryptBytes)
//	decrypted = pKCS5UnPadding(decrypted)
//	return string(decrypted), nil
//}
