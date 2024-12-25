package startupcfg

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

type (
	Encrypted string // Encrypted 加密串
)

var (
	decryptFunc = func(e Encrypted) (string, error) {
		return string(e), nil
	} // 默认的密码直接返回
	encryptFunc = func(e string) (Encrypted, error) {
		return Encrypted(e), nil
	} // 默认将字符串转化为Encrypted类型
)

// SetDecryptHandler 设置解密函数
func SetDecryptHandler(f func(m Encrypted) (string, error)) {
	if f == nil {
		return
	}
	decryptFunc = f
}

// SetEncryptHandler 设置加密函数
func SetEncryptHandler(f func(m string) (Encrypted, error)) {
	if f == nil {
		return
	}
	encryptFunc = f
}

// Get 获取解密串儿
func (e Encrypted) Get() (string, error) {
	if string(e) == "" {
		return "", nil
	}
	if decryptFunc != nil {
		return decryptFunc(e)
	}
	return "", fmt.Errorf("no set defaultDecrypt")
}

// MarshalJSON 实现json Marshaler接口 自定义json 编码
func (e *Encrypted) MarshalJSON() ([]byte, error) {
	decrypted, err := e.Get()
	if err != nil {
		return nil, err
	}
	return ([]byte)(fmt.Sprintf("\"%s\"", decrypted)), nil
}

// UnmarshalJSON 实现Unmarshaler接口 自定义json解码
func (e *Encrypted) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("UnmarshalJSON Encrypted(%s) to string failed: %w ", data, err)
	}
	if str == "" {
		return nil
	}
	dd, err := encryptFunc(str)
	if err != nil {
		return fmt.Errorf("UnmarshalJSON Encrypted(%s) failed: %w ", str, err)
	}
	*e = dd
	return nil
}

// MarshalYAML 实现yaml Marshaler接口
func (e *Encrypted) MarshalYAML() (interface{}, error) {
	decrypted, err := e.Get()
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

// UnmarshalYAML 实现Unmarshaler接口 自定义yaml解码
func (e *Encrypted) UnmarshalYAML(value *yaml.Node) error {
	if value.Value == "" {
		return nil
	}
	dd, err := encryptFunc(value.Value)
	if err != nil {
		return fmt.Errorf("UnmarshalYAML Encrypted(%s) failed: %w ", value.Value, err)
	}
	*e = dd
	return nil
}
