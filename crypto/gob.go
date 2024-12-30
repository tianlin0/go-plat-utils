package crypto

import (
	"bytes"
	"encoding/gob"
)

// GobEncode 序列化一个对象
func GobEncode(s interface{}, en EnDeCoder) (string, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	err := enc.Encode(s)
	if err != nil {
		return "", err
	}

	if en == nil {
		en = new(Base64Coder)
	}

	return en.Encode(buf.Bytes()), nil
}

// GobDecode 反序列化一个对象
func GobDecode(s string, data interface{}, de EnDeCoder) error {
	var b bytes.Buffer
	if de == nil {
		de = new(Base64Coder)
	}
	old, err := de.Decode(s)
	if err != nil {
		return err
	}
	_, err = b.Write(old)
	if err != nil {
		return err
	}
	enc := gob.NewDecoder(&b)
	err = enc.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
