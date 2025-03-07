package uuid

import "github.com/rs/xid"
import "github.com/google/uuid"

// GetXId id 生成器
func GetXId() string {
	guid := xid.New()
	return guid.String()
}

func GetUUIDv7() string {
	id, _ := uuid.NewV7()
	return id.String()
}
