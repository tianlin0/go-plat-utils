package crypto

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

// GobEncode 序列化一个对象
func GobEncode(s interface{}, en Encoder) (string, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	err := enc.Encode(s)
	if err != nil {
		return "", err
	}

	if en == nil {
		en = base64.StdEncoding.EncodeToString
	}

	return en(buf.Bytes()), nil
}

// GobDecode 反序列化一个对象
func GobDecode(s string, data interface{}, de Decoder) error {
	var b bytes.Buffer
	if de == nil {
		de = base64.StdEncoding.DecodeString
	}
	old, err := de(s)
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
