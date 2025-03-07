package conv_test

import (
	"github.com/tianlin0/go-plat-utils/conv"
	"testing"
)

type AA struct {
}

func TestUnmarshal(t *testing.T) {
	var aa *AA
	conv.Unmarshal(aa, new(AA))
}
