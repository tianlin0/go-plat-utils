package startupcfg

// Custom 自定义配置抽象
type Custom interface {
	// GetSensitive 查询敏感配置（加密）对应key的value
	GetSensitive(key string) (string, error)
	// GetNormal 查询非敏感配置对应key的value
	GetNormal(key string) interface{}
}
