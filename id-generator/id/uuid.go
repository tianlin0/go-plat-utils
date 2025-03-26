package id

import "github.com/rs/xid"
import "github.com/google/uuid"

// GetXId 20字符 id 生成器,如：cvhmhh6s295a4l56g4a0
func GetXId() string {
	guid := xid.New()
	return guid.String()
}

// GetUUIDv7 36字符 id 生成器,如：0195d052-4c80-7217-ad19-1acb84b04d4f
func GetUUIDv7() string {
	id, _ := uuid.NewV7()
	return id.String()
}
