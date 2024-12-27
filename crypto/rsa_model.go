package crypto

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
)

type RSAPrivateKeyType int

const (
	PKCS1Type RSAPrivateKeyType = iota
	PKCS8Type

	rSAKeyBit512  = 512
	rSAKeyBit1024 = 1024
	rSAKeyBit2048 = 2048
	rSAKeyBit4096 = 4096
)

type RSASecurity struct {
	pubStr string          //公钥字符串
	priStr string          //私钥字符串
	pubKey *rsa.PublicKey  //公钥
	priKey *rsa.PrivateKey //私钥
}

func getBits(bits int) int {
	if bits < 0 {
		bits = rSAKeyBit512
	} else if bits < 1024 {
		bits = rSAKeyBit1024
	} else if bits < 2048 {
		bits = rSAKeyBit2048
	} else if bits < 4096 {
		bits = rSAKeyBit4096
	} else {
		bits = rSAKeyBit4096
	}
	return bits
}

// CreateRSAKeys 生成RSA私钥和公钥
func (r *RSASecurity) CreateRSAKeys(bits int, keyType RSAPrivateKeyType) (publicKeyString string, privateKeyString string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, getBits(bits))
	if err != nil {
		return "", "", err
	}

	// 将私钥转换为PEM格式
	var privateKeyBlock *pem.Block
	if keyType == PKCS1Type {
		privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privateKeyBlock = &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		}
	} else {
		privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return "", "", err
		}
		privateKeyBlock = &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: privateKeyBytes,
		}
	}

	// 将公钥转换为PEM格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	privateKeyString = string(pem.EncodeToMemory(privateKeyBlock))
	publicKeyString = string(pem.EncodeToMemory(publicKeyBlock))

	if r.pubStr == "" || r.priStr == "" {
		r.SetPublicAndPrivateKey(publicKeyString, privateKeyString)
	}

	return publicKeyString, privateKeyString, nil
}

// SetPublicAndPrivateKey 设置公钥私钥
func (r *RSASecurity) SetPublicAndPrivateKey(pubStr string, priStr string) (err error) {
	if pubStr != "" {
		r.pubStr = pubStr
		r.pubKey, err = r.getPublicKey()
		if err != nil {
			return err
		}
	}
	if priStr != "" {
		r.priStr = priStr
		r.priKey, err = r.getPrivateKey()
		if err != nil {
			return err
		}
	}
	return nil
}

// getPublicKey *rsa.PrivateKey
func (r *RSASecurity) getPublicKey() (*rsa.PublicKey, error) {
	return getPubKey([]byte(r.pubStr))
}

// getPrivateKey *rsa.PublicKey
func (r *RSASecurity) getPrivateKey() (*rsa.PrivateKey, error) {
	return getPriKey([]byte(r.priStr))
}

// PubKeyEncrypt 公钥加密
func (r *RSASecurity) PubKeyEncrypt(input []byte) ([]byte, error) {
	if r.pubKey == nil {
		return []byte(""), fmt.Errorf(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(r.pubKey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return io.ReadAll(output)
}

// PriKeyDecrypt 私钥解密
func (r *RSASecurity) PriKeyDecrypt(input []byte) ([]byte, error) {
	if r.priKey == nil {
		return []byte(""), fmt.Errorf(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(r.priKey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}

	return io.ReadAll(output)
}

// PubKeyEncryptBase64 公钥加密
func (r *RSASecurity) PubKeyEncryptBase64(input string) (string, error) {
	oldByte, err := r.PubKeyEncrypt([]byte(input))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(oldByte), nil
}

// PriKeyDecryptBase64 私钥解密
func (r *RSASecurity) PriKeyDecryptBase64(encodeBase64String string) (string, error) {
	newByte, err := base64.StdEncoding.DecodeString(encodeBase64String)
	if err != nil {
		return "", err
	}
	newStr, err := r.PriKeyDecrypt(newByte)
	if err != nil {
		return "", err
	}
	return string(newStr), nil
}

// PriKeyEncrypt 私钥加密，用于数字签名
func (r *RSASecurity) PriKeyEncrypt(input []byte) ([]byte, error) {
	if r.priKey == nil {
		return []byte(""), fmt.Errorf(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(r.priKey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return io.ReadAll(output)
}

// PubKeyDecrypt 公钥解密，用于数字签名
func (r *RSASecurity) PubKeyDecrypt(input []byte) ([]byte, error) {
	if r.pubKey == nil {
		return []byte(""), fmt.Errorf(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(r.pubKey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}
	return io.ReadAll(output)
}

// PriKeyEncryptBase64 私钥加密
func (r *RSASecurity) PriKeyEncryptBase64(input string) (string, error) {
	oldByte, err := r.PriKeyEncrypt([]byte(input))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(oldByte), nil
}

// PubKeyDecryptBase64 公钥解密
func (r *RSASecurity) PubKeyDecryptBase64(encodeBase64String string) (string, error) {
	newByte, err := base64.StdEncoding.DecodeString(encodeBase64String)
	if err != nil {
		return "", err
	}
	newStr, err := r.PubKeyDecrypt(newByte)
	if err != nil {
		return "", err
	}
	return string(newStr), nil
}

// SignMd5 使用RSAWithMD5算法签名
func (r *RSASecurity) SignMd5(data string) (string, error) {
	md5Hash := md5.New()
	sData := []byte(data)
	md5Hash.Write(sData)
	hashed := md5Hash.Sum(nil)

	signByte, err := rsa.SignPKCS1v15(rand.Reader, r.priKey, crypto.MD5, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return string(sign), err
}

// SignSha1 使用RSAWithSHA1算法签名
func (r *RSASecurity) SignSha1(data string) (string, error) {
	sha1Hash := sha1.New()
	sData := []byte(data)
	sha1Hash.Write(sData)
	hashed := sha1Hash.Sum(nil)

	signByte, err := rsa.SignPKCS1v15(rand.Reader, r.priKey, crypto.SHA1, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return string(sign), err
}

// SignSha256 使用RSAWithSHA256算法签名
func (r *RSASecurity) SignSha256(data string) (string, error) {
	sha256Hash := sha256.New()
	sData := []byte(data)
	sha256Hash.Write(sData)
	hashed := sha256Hash.Sum(nil)

	signByte, err := rsa.SignPKCS1v15(rand.Reader, r.priKey, crypto.SHA256, hashed)
	sign := base64.StdEncoding.EncodeToString(signByte)
	return string(sign), err
}

// VerifySignMd5 使用RSAWithMD5验证签名
func (r *RSASecurity) VerifySignMd5(data string, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := md5.New()
	hash.Write([]byte(data))
	return rsa.VerifyPKCS1v15(r.pubKey, crypto.MD5, hash.Sum(nil), sign)
}

// VerifySignSha1 使用RSAWithSHA1验证签名
func (r *RSASecurity) VerifySignSha1(data string, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(data))
	return rsa.VerifyPKCS1v15(r.pubKey, crypto.SHA1, hash.Sum(nil), sign)
}

// VerifySignSha256 使用RSAWithSHA256验证签名
func (r *RSASecurity) VerifySignSha256(data string, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := sha256.New()
	hash.Write([]byte(data))

	return rsa.VerifyPKCS1v15(r.pubKey, crypto.SHA256, hash.Sum(nil), sign)
}
