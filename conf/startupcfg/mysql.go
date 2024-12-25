package startupcfg

import (
	"fmt"
	"log"
	"net/url"
)

// MysqlConfig mysql配置
type MysqlConfig struct {
	UserName        string    `json:"username" yaml:"username"`
	PasswordEncoded Encrypted `json:"pwEncoded" yaml:"pwEncoded"`
	Address         string    `json:"address" yaml:"address"`
	Database        string    `json:"database" yaml:"database"`
	Charset         string    `json:"charset" yaml:"charset"`
}

// DriverName 使用sql.Open连接数据库时的driverName参数
func (c *MysqlConfig) DriverName() string {
	return string(DriverMysql)
}

// DatasourceName 使用sql.Open连接数据库时的datasourceName参数
func (c *MysqlConfig) DatasourceName() string {
	if c.Charset == "" {
		c.Charset = _MYSQL_CHARSET
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=Local",
		c.UserName,
		url.QueryEscape(c.Password()),
		c.Address,
		c.Database,
		c.Charset)
}

// ServerAddress mysql服务器地址
func (c *MysqlConfig) ServerAddress() string {
	return c.Address
}

// User mysql数据库用户
func (c *MysqlConfig) User() string {
	return c.UserName
}

// Password mysql数据库用户密码
func (c *MysqlConfig) Password() string {
	pass, err := c.PasswordEncoded.Get()
	if err != nil {
		log.Println("password decode error:", err)
		return ""
	}
	return pass
}

// DatabaseName mysql数据库名称
func (c *MysqlConfig) DatabaseName() interface{} {
	return c.Database
}

// Extend 扩展字段
func (c *MysqlConfig) Extend(name ExtendField) interface{} {
	switch name {
	case extendMysqlCharset:
		return c.Charset
	}
	return nil
}
