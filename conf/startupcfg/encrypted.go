package startupcfg

import (
	"fmt"
)

type (
	Encrypted string // Encrypted 加密串
)

var (
	encryptFunc = func(e string) (Encrypted, error) {
		return Encrypted(e), nil
	} // 默认将字符串转化为Encrypted类型
	decryptFunc = func(e Encrypted) (string, error) {
		return string(e), nil
	} // 默认的密码直接返回
)

// setEncryptHandler 设置加密函数,一般不需要设置
func setEncryptHandler(encryptF func(m string) (Encrypted, error)) {
	if encryptF != nil {
		encryptFunc = encryptF
	}
}

// SetDecryptHandler 设置解密函数
func SetDecryptHandler(decryptF func(m Encrypted) (string, error)) {
	if decryptF != nil {
		decryptFunc = decryptF
	}
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

//// MarshalJSON 实现json Marshaler接口 自定义json 编码
//func (e *Encrypted) MarshalJSON() ([]byte, error) {
//	decrypted, err := e.Get()
//	if err != nil {
//		return nil, err
//	}
//	return ([]byte)(fmt.Sprintf("\"%s\"", decrypted)), nil
//}
//
//// UnmarshalJSON 实现Unmarshaler接口 自定义json解码
//func (e *Encrypted) UnmarshalJSON(data []byte) error {
//	var raw json.RawMessage = data
//	var kindType string
//	if err := json.Unmarshal(raw, &kindType); err != nil {
//		return err
//	}
//
//	var str string
//	if err := json.Unmarshal(raw, &str); err != nil {
//		return fmt.Errorf("UnmarshalJSON Encrypted(%s) to string failed: %w ", data, err)
//	}
//	if str == "" {
//		return nil
//	}
//	//str 有可能是已加密过的字符串，或者是解密后的字符串，这里需要区分开
//
//	dd, err := encryptFunc(str)
//	if err != nil {
//		return fmt.Errorf("UnmarshalJSON Encrypted(%s) failed: %w ", str, err)
//	}
//	*e = dd
//	return nil
//}
//
//// MarshalYAML 实现yaml Marshaler接口
//func (e *Encrypted) MarshalYAML() (interface{}, error) {
//	decrypted, err := e.Get()
//	if err != nil {
//		return nil, err
//	}
//	return decrypted, nil
//}
//
//// UnmarshalYAML 实现Unmarshaler接口 自定义yaml解码
//func (e *Encrypted) UnmarshalYAML(value *yaml.Node) error {
//	if value.Value == "" {
//		return nil
//	}
//	dd, err := encryptFunc(value.Value)
//	if err != nil {
//		return fmt.Errorf("UnmarshalYAML Encrypted(%s) failed: %w ", value.Value, err)
//	}
//	*e = dd
//	return nil
//}
