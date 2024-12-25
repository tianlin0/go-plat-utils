package startupcfg

type (
	DriverType  string // DriverType 连接类型
	ExtendField string // ExtendField 扩展字段名
)

const (
	_MYSQL_CHARSET = "utf8"
	_TDMQ_PROTOCOL = "pulsar"
)

var (
	DriverMysql DriverType = "mysql"
	DriverRedis DriverType = "redis"

	extendMysqlCharset ExtendField = "charset"
	extendRedisUseTLS  ExtendField = "TLS"

	_, _ Database = new(MysqlConfig), new(RedisConfig)
)
