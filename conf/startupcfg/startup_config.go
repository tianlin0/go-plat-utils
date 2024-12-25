package startupcfg

// StartupConfig 启动配置结构
type StartupConfig struct {
	ApiConfig    map[string]*ServiceApiConfig `json:"api" yaml:"api"`
	MySqlMap     map[string]*MysqlConfig      `json:"mysql" yaml:"mysql"`
	RedisMap     map[string]*RedisConfig      `json:"redis" yaml:"redis"`
	CustomConfig *CustomConfig                `json:"custom" yaml:"custom"`
}

// ServiceAPI 服务接口参数
func (c *StartupConfig) ServiceAPI(serviceName string) ServiceAPI {
	if c != nil && c.ApiConfig != nil {
		if api, ok := c.ApiConfig[serviceName]; ok {
			return api
		}
	}
	return nil
}

// MySQL 返回使用的MySQL连接参数
func (c *StartupConfig) MySQL(name string) Database {
	if c != nil && c.MySqlMap != nil {
		if mysql, ok := c.MySqlMap[name]; ok {
			return mysql
		}
	}
	return nil
}

// Redis 返回使用Redis连接参数
func (c *StartupConfig) Redis(name string) Database {
	if c != nil && c.RedisMap != nil {
		if redis, ok := c.RedisMap[name]; ok {
			return redis
		}
	}
	return nil
}

// Custom 返回自定义配置参数
func (c *StartupConfig) Custom() Custom {
	if c != nil {
		return c.CustomConfig
	}
	return nil
}
