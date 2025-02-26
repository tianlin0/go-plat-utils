package id

import (
	"github.com/rs/xid"
)

// GetXId id 生成器
func GetXId() string {
	guid := xid.New()
	return guid.String()
}
