package startupcfg

// Database 连接参数的抽象，包含使用sql.Open连接数据库时的参数
type Database interface {
	// DriverName 使用sql.Open连接数据库时的driverName参数
	DriverName() string
	// DatasourceName 使用sql.Open连接数据库时的datasourceName参数
	DatasourceName() string
	// ServerAddress 数据库服务器地址
	ServerAddress() string
	// Password 数据库用户密码
	Password() string
	// DatabaseName 数据库名称
	DatabaseName() interface{}
	// User 数据库用户
	User() string
	// Extend 扩展信息
	Extend(ExtendField) interface{}
}
