package utils

import (
	"fmt"
	"github.com/marspere/goencrypt"
	"github.com/samber/lo"
	"strings"
	"unicode/utf8"

	gguid "github.com/google/uuid"
	gouuid "github.com/nu7hatch/gouuid"
	"github.com/tianlin0/go-plat-utils/internal"
)

// RandomString 生成随机字符串
func RandomString(l int, sourceStr ...string) string {
	var str = append(lo.NumbersCharset, lo.LowerCaseLettersCharset...)
	if len(sourceStr) > 0 {
		sourceArr := make([]rune, 0, len(sourceStr))
		for _, one := range sourceStr {
			sourceArr = append(sourceArr, []rune(one)...)
		}
		if len(sourceArr) > 0 {
			str = sourceArr
		}
	}
	if l <= 0 {
		return ""
	}
	return lo.RandomString(l, str)

	//bytes := []byte(str)
	//result := make([]byte, 0)
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i := 0; i < l; i++ {
	//	result = append(result, bytes[r.Intn(len(bytes))])
	//}
	//return string(result)
}

// UnicodeDecodeString 解码unicode
func UnicodeDecodeString(s string) string {
	if s == "" {
		return s
	}
	newStr := make([]string, 0)
	for i := 0; i < len(s); {
		r, n := utf8.DecodeRuneInString(s[i:])
		newStr = append(newStr, fmt.Sprintf("%c", r))
		i += n
	}
	if len(newStr) == 0 {
		return s
	}
	return strings.Join(newStr, "")
}

// ChangeVariableName 将驼峰与小写互转
func ChangeVariableName(varName string, toType ...string) string {
	if varName == "" {
		return varName
	}

	if len(toType) == 1 {
		if toType[0] == "lower" {
			return internal.SnakeString(varName)
		} else if toType[0] == "upper" {
			return internal.PascalString(varName)
		}
	}

	// 检测是否都为小写
	isLower := true
	for i := 0; i < len(varName); i++ {
		c := varName[i]
		if isASCIIUpper(c) {
			isLower = false
			break
		}
	}

	if !isLower {
		return internal.SnakeString(varName)
	}
	return internal.PascalString(varName)
}

// NewUUID 新建uuid
func NewUUID() string {
	uuidStr := ""
	uuids, err := gguid.NewUUID()
	if err == nil {
		uuidStr = uuids.String()
		if uuidStr != "" {
			return uuidStr
		}
	}

	uuidStr = gguid.New().String()
	if uuidStr != "" {
		return uuidStr
	}

	uuidTemp, err := gouuid.NewV4()
	if err == nil {
		uuidStr = uuidTemp.String()
		if uuidStr != "" {
			return uuidStr
		}
	}

	return uuidStr
}

// GetUUID 获取uuid格式串
func GetUUID(s string) string {
	uuid, err := goencrypt.MD5(s)
	if len(uuid) != 32 || err != nil {
		return ""
	}
	return fmt.Sprintf("%s-%s-%s-%s-%s", uuid[0:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
