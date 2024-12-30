package crypto

// XorEncode 异或加密
func XorEncode(msg string, key int, en EnDeCoder) string {
	byteList := []byte(msg)
	pwd := make([]byte, len(byteList))
	for i := 0; i < len(byteList); i++ {
		pwd[i] = byteList[i] ^ byte(key)
	}

	if en == nil {
		en = new(Base64Coder)
	}

	return en.Encode(pwd)
}

// XorDecode 异或解密
func XorDecode(msg string, key int, de EnDeCoder) (string, error) {
	if de == nil {
		de = new(Base64Coder)
	}
	pwdList, err := de.Decode(msg)
	if err != nil {
		return "", err
	}
	old := make([]byte, len(pwdList))
	for i := 0; i < len(pwdList); i++ {
		old[i] = pwdList[i] ^ byte(key)
	}
	return string(old), nil
}
