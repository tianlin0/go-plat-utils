package crypto_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/marspere/goencrypt"
	"github.com/tianlin0/go-plat-utils/crypto"
)

func ExampleMd5() {
	str := "hello"
	res := crypto.Md5(str)

	fmt.Println(res)

	// Output:
	// 5d41402abc4b2a76b9719d911017c592
}

func ExampleMd51() {
	s := "hello world"

	value, err := goencrypt.MD5(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	uuid := func(s string) string {
		d := md5.Sum([]byte(s))
		return hex.EncodeToString(d[:])
	}(s)

	if uuid == string(value) {
		fmt.Println("ok")
	}

	fmt.Println(uuid)
	// Output:
	// 5eb63bbbe01eeed093cb22bb8f5acdc3
}
