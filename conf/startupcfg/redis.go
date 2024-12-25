package startupcfg

import (
	"fmt"
	"net/url"
)

// RedisConfig redis配置
type RedisConfig struct {
	PasswordEncoded Encrypted `json:"pwEncoded" yaml:"pwEncoded"`
	Address         string    `json:"address" yaml:"address"`
	Database        int64     `json:"database" yaml:"database"`
	Username        string    `json:"username" yaml:"username"`
	UseTLS          bool      `json:"useTLS" yaml:"useTLS"`
}

// DriverName 驱动名称
func (c *RedisConfig) DriverName() string {
	return string(DriverRedis)
}

// DatasourceName 连接数据库时的datasourceName参数
func (c *RedisConfig) DatasourceName() string {
	return fmt.Sprintf("redis://%s:%s@%s/%d",
		c.Username,
		url.QueryEscape(c.Password()),
		c.Address,
		c.Database)
}

// ServerAddress redis服务器地址
func (c *RedisConfig) ServerAddress() string {
	return c.Address
}

// Password redis数据库用户密码
func (c *RedisConfig) Password() string {
	pass, err := c.PasswordEncoded.Get()
	if err != nil {
		return ""
	}
	return pass
}

// DatabaseName redis数据库名称
func (c *RedisConfig) DatabaseName() interface{} {
	return c.Database
}

// User redis数据库用户
func (c *RedisConfig) User() string {
	return c.Username
}

// Extend 扩展字段
func (c *RedisConfig) Extend(name ExtendField) interface{} {
	switch name {
	case extendRedisUseTLS:
		return c.UseTLS
	}
	return nil
}
