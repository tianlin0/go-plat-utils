package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/marspere/goencrypt"
	"log"
)

const (
	MD5 shaType = iota
	SHA1
	SHA256
	SHA512
)

type shaType int

// SHA 计算给定字符串的SHA值
func SHA(sharType shaType, s string, printType ...int) string {
	pType := goencrypt.PrintHex // 默认打印类型为十六进制
	if len(printType) > 0 {
		pType = printType[0]
	}

	// 使用辅助函数进行SHA计算
	return calculateSHA(sharType, s, pType)
}

// calculateSHA 辅助函数：根据sha类型计算SHA值
func calculateSHA(sharType shaType, s string, pType int) string {
	switch sharType {
	case MD5: // 计算出md5的值
		return hashMd5(s, pType) // 计算出md5的值并返回
	case SHA1: // 计算出sha1的值
		return calculateHash(goencrypt.SHA1, s, pType) // 调用计算哈希的函数
	case SHA256: // 计算出sha256的值
		return calculateHash(goencrypt.SHA256, s, pType) // 调用计算哈希的函数
	case SHA512: // 计算出sha512的值
		return calculateHash(goencrypt.SHA512, s, pType) // 调用计算哈希的函数
	default: // 处理不支持的sha类型
		log.Println("Unsupported SHA type") // 记录不支持的SHA类型日志
		return ""                           // 返回空字符串
	}
}

// calculateHash 辅助函数：实际进行哈希计算
func calculateHash(shaType int, s string, pType int) string {
	value, err := goencrypt.SHA(shaType, []byte(s), pType) // 计算哈希值
	if err != nil {                                        // 如果计算发生错误
		log.Println("Error calculating hash: ", err) // 记录错误信息
		return ""                                    // 返回空字符串
	}
	return value // 返回计算出的哈希值
}

func shaEncode(content []byte, decodeType int) string {
	if decodeType == goencrypt.PrintHex {
		return hex.EncodeToString(content)
	}
	if decodeType == goencrypt.PrintBase64 {
		return base64.StdEncoding.EncodeToString(content)
	}
	return string(content)
}

func hashMd5(s string, printType int) string {
	value, err := goencrypt.MD5(s)
	if err != nil {
		return ""
	}
	if printType == goencrypt.PrintHex {
		return string(value)
	}
	contentByte, err := hex.DecodeString(string(value))
	if err != nil {
		return string(value)
	}
	return shaEncode(contentByte, printType)
}

// Md5 计算出md5的值
func Md5(s string) string {
	return SHA(MD5, s)
}

// HmacSha256 转化为hmac
func HmacSha256(s, secret string) string {
	hashed := hmac.New(sha256.New, []byte(secret))
	hashed.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(hashed.Sum(nil))
}
