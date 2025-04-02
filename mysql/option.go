package mysql

import (
	"fmt"
	"strings"
)

type Option func(*option)

// LogLevel log level
type LogLevel int

const (
	// Silent silent log level
	Silent LogLevel = iota + 1
	// Error error log level
	Error
	// Warn warn log level
	Warn
	// Info info log level
	Info
)

type option struct {
	host, user, password, database string
	port                           int
	maxOpenConns                   int
	maxIdleConns                   int
	logLevel                       LogLevel
	otherArgs                      []string
	dsn                            string //Data Source Name
}

/*
DSN（Data Source Name）是一种用于配置数据库连接的字符串，它包含连接到数据库所需的各种信息，比如主机地址、端口、
数据库名称、用户名、密码等。DSN 常用于像 GORM 这样的数据库库，以简化数据库连接设置。
postgers：host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
*/
func (o option) getDSN() string {
	if o.dsn != "" {
		return o.dsn
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		o.user, o.password, o.host, o.port, o.database, strings.Join(o.otherArgs, "&"))
}

func newOption() *option {
	return &option{
		host:     "localhost",
		port:     3306,
		logLevel: Error,
	}
}

func WithHost(host string) Option {
	return func(o *option) {
		o.host = host
	}
}

func WithPort(port int) Option {
	return func(o *option) {
		o.port = port
	}
}

func WithAuth(user, password string) Option {
	return func(o *option) {
		o.user = user
		o.password = password
	}
}

func WithDatabase(database string) Option {
	return func(o *option) {
		o.database = database
	}
}

func WithMaxOpenConns(num int) Option {
	return func(o *option) {
		o.maxOpenConns = num
	}
}

func WithLogLevel(lv LogLevel) Option {
	return func(o *option) {
		o.logLevel = lv
	}
}

// The default max idle connections is currently 2.
func WithMaxIdleConns(num int) Option {
	return func(o *option) {
		o.maxIdleConns = num
	}
}

func WithArgs(args ...string) Option {
	return func(o *option) {
		o.otherArgs = append(o.otherArgs, args...)
	}
}

func WithDSN(dsn string) Option {
	return func(o *option) {
		o.dsn = dsn
	}
}
