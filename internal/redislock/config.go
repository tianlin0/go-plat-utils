package lredis

import (
	"fmt"
	"time"
)

const (
	DefaultExpireTime = 10 * time.Second // 默认过期时间10s
	DefaultKeyFront   = "{redis-lock}"
)

func getLockerKeyName(key string) string {
	return fmt.Sprintf("%s%s", DefaultKeyFront, key)
}
