package postgres

import (
	"strconv"
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
	host, user, password, port, database, sslmode, timezone string
	maxOpenConns                                            int
	maxIdleConns                                            int
	logLevel                                                LogLevel
	dsn                                                     string //Data Source Name
}

/*
DSN（Data Source Name）是一种用于配置数据库连接的字符串，它包含连接到数据库所需的各种信息，比如主机地址、端口、
数据库名称、用户名、密码等。DSN 常用于像 GORM 这样的数据库库，以简化数据库连接设置。
postgers：host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
*/

func (o option) getDSN() string {
	if o.dsn != "" {
		return o.dsn
	}
	args := []string{
		"host=" + o.host,
		"user=" + o.user,
		"password=" + o.password,
		"dbname=" + o.database,
		"port=" + o.port,
		"sslmode=" + o.sslmode,
		"TimeZone=" + o.timezone,
	}
	return strings.Join(args, " ")
}

func newOption() *option {
	return &option{
		host:     "localhost",
		port:     "5432",
		sslmode:  "disable",
		timezone: "Asia/Shanghai",
		database: "postgers",
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
		o.port = strconv.Itoa(port)
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

func WithSSLMode(sslmode bool) Option {
	return func(o *option) {
		if sslmode {
			o.sslmode = "enable"
		} else {
			o.sslmode = "disable"
		}
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

func WithDSN(dsn string) Option {
	return func(o *option) {
		o.dsn = dsn
	}
}
