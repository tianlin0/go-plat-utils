package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	if length-unPadding < 0 { // 不能<0
		return []byte{}
	}
	return origData[:(length - unPadding)]
}

func aesEncrypt(origData, key []byte, iv string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)

	var ivBytes []byte
	if iv == "" {
		ivBytes = key[:blockSize]
	} else {
		ivBytes = []byte(iv)
	}

	blockMode := cipher.NewCBCEncrypter(block, ivBytes)
	crypt := make([]byte, len(origData))
	blockMode.CryptBlocks(crypt, origData)
	return crypt, nil
}

func aesDecrypt(crypt, key []byte, iv string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var ivBytes []byte
	if iv == "" {
		blockSize := block.BlockSize()
		ivBytes = key[:blockSize]
	} else {
		ivBytes = []byte(iv)
	}
	blockMode := cipher.NewCBCDecrypter(block, ivBytes)
	origData := make([]byte, len(crypt))

	blockMode.CryptBlocks(origData, crypt)
	origData = pKCS5UnPadding(origData)
	return origData, nil
}

func getAesKeyFromBase64(keyBase64 string) []byte {
	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		//表示传的不是base64格式的，则内部兼容处理一下
		key = []byte(keyBase64)
	}

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		//太短或太长，则进行转化
		newKey := Md5(string(key))
		key = []byte(newKey)
	}
	return key
}
