package crypto_test

import (
	"fmt"
	"github.com/tianlin0/go-plat-utils/crypto"
)

func ExampleMd5() {
	str := "hello"
	res := crypto.Md5(str)

	fmt.Println(res)

	// Output:
	// 5d41402abc4b2a76b9719d911017c592
}
