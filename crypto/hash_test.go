package crypto_test

import (
	"github.com/tianlin0/go-plat-utils/crypto"
	"github.com/tianlin0/go-plat-utils/utils"
	"testing"
)

func TestSHAForHmac(t *testing.T) {
	key := "tianlin020250214"
	testCases := []*utils.TestStruct{
		{"SHAWithHmac", []any{crypto.SHA256, "hello world", key}, []any{"b390cd4fcd9864133e838efa76ee3e0b0e0b4774dc04a646edc956ba34b8072c"}, crypto.SHAWithHmac},
	}
	utils.TestFunction(t, testCases, nil)
}
