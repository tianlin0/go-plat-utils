package startupcfg

// Startup 服务运行配置，有别于启动配置，运行配置是从配置文件或配置中心获取的配置，而非启动参数或环境变量获取的配置
type Startup interface {
	// MySQL 返回使用的MySQL连接参数
	MySQL(name string) Database
	// Redis 返回使用Redis连接参数
	Redis(name string) Database
	// ServiceAPI 服务接口参数
	ServiceAPI(serviceName string) ServiceAPI
	// Custom 自定义配置参数
	Custom() Custom
}
