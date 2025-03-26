package id_test

import (
	"fmt"
	"github.com/tianlin0/go-plat-utils/id-generator/id"
	"testing"
)

func TestAnyToBool(t *testing.T) {
	for i := 0; i < 20; i++ {
		ss := id.Generator(700)
		fmt.Println(ss)
	}
	fmt.Println("ok")
}
