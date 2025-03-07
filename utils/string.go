package utils

import (
	"fmt"
	"github.com/marspere/goencrypt"
	"github.com/samber/lo"
	"github.com/tianlin0/go-plat-utils/utils/id"
	"regexp"
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
	uuidGenerators := []func() (string, error){ // 定义一个切片，存储不同的UUID生成函数
		func() (string, error) {
			uuids, err := gguid.NewUUID()
			if err != nil {
				return "", err
			}
			return uuids.String(), nil
		}, // 使用gguid生成UUID
		func() (string, error) { return gguid.New().String(), nil }, // 使用gguid的另一个生成方法
		func() (string, error) {
			uuidTemp, err := gouuid.NewV4()
			if err != nil {
				return "", err
			}
			return uuidTemp.String(), nil
		}, // 使用gouuid生成UUID
		func() (string, error) {
			uuidTemp, err := id.GeneratorBase32()
			if err != nil {
				return "", err
			}
			return GetUUID(uuidTemp), nil
		}, // 使用sonyflake生成UUID
	}

	for _, generator := range uuidGenerators { // 遍历每个UUID生成函数
		uuidStr, err := generator()      // 尝试生成UUID
		if err == nil && uuidStr != "" { // 如果没有错误且UUID字符串不为空
			return uuidStr // 返回生成的UUID字符串
		}
	}

	return "" // 如果所有方法都失败，返回空字符串
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

// ReplaceDynamicVariables 将动态分隔符包裹的变量替换为 [变量名] 格式
// 参数：
// - input：原始字符串
// - startDelimiter：起始分隔符
// - endDelimiter：结束分隔符
func ReplaceDynamicVariables(input string, startDelimiter, endDelimiter string, replaceStartDelimiter, replaceEndDelimiter string) string {
	// 转义分隔符中的特殊字符，确保它们在正则表达式中被正确处理
	escapedStart := regexp.QuoteMeta(startDelimiter)
	escapedEnd := regexp.QuoteMeta(endDelimiter)

	// 构建正则表达式：匹配 startDelimiter + 变量名 + endDelimiter
	regexPattern := fmt.Sprintf(`%s(.*?)%s`, escapedStart, escapedEnd)
	re := regexp.MustCompile(regexPattern)

	// 替换匹配到的变量
	output := re.ReplaceAllStringFunc(input, func(match string) string {
		// 提取变量名（去掉分隔符）
		varName := strings.TrimPrefix(strings.TrimSuffix(match, endDelimiter), startDelimiter)
		return fmt.Sprintf("%s%s%s", replaceStartDelimiter, varName, replaceEndDelimiter)
	})

	return output
}
