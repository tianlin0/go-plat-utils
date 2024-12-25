package startupcfg

// ServiceAPI 服务接口抽象
type ServiceAPI interface {
	// DomainName 接口域名
	DomainName() string
	// Url 接口Url
	Url(apiName string) string
	// AuthData 接口其他数据（鉴权数据等）
	AuthData(key string) (string, error)
}
