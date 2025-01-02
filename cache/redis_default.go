package cache

import (
	startupCfg "github.com/tianlin0/go-plat-utils/conf/startupcfg"
)

var (
	defaultRedisCfg *startupCfg.RedisConfig
)

// SetDefaultRedisConfig 切换默认的redis连接
func SetDefaultRedisConfig(con *startupCfg.RedisConfig) {
	if con != nil {
		defaultRedisCfg = con
	}
}
