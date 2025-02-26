package id_test

import (
	"fmt"
	"github.com/tianlin0/go-plat-utils/utils/id"
	"testing"
)

func TestGeneratorBase32(t *testing.T) {
	aa, _ := id.GeneratorBase32()
	fmt.Println(aa)
	aa = id.GetXId()
	fmt.Println(aa)
}
