package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
)

var (
	pcdRegexp, rawRegexp *regexp.Regexp
)

func init() {
	pcdRegexp, _ = regexp.Compile(`@i_(\w+)`)
	rawRegexp, _ = regexp.Compile(`{(\w+)}`)
}

type Cli struct {
	*gorm.DB
}

func Connect(dsn string, options ...Option) (*Cli, error) {
	opt := option{}
	for _, f := range options {
		f(&opt)
	}
	if int(opt.logLevel) == 0 {
		opt.logLevel = Error
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(opt.logLevel)),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	if opt.maxOpenConns >= opt.maxIdleConns {
		sqlDB.SetMaxOpenConns(opt.maxOpenConns)
	}

	if opt.maxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(opt.maxIdleConns)
	}

	return &Cli{DB: db}, err
}

func (c *Cli) Begin() *Tx {
	return &Tx{c.DB.Begin()}
}

func (c *Cli) ExecProcedure(sql string, args map[string]interface{}) *gorm.DB {
	return execSql(pcdRegexp, c.DB, sql, args)
}

func (c *Cli) ExecRawSql(sql string, args map[string]interface{}) *gorm.DB {
	return execSql(rawRegexp, c.DB, sql, args)
}

func execSql(exp *regexp.Regexp, db *gorm.DB, sql string, args map[string]interface{}) *gorm.DB {
	inKeys := exp.FindAllString(sql, -1)
	values := make([]interface{}, len(inKeys))
	for i, key := range inKeys {
		values[i] = args[key]
	}
	sql = rawRegexp.ReplaceAllString(sql, "?")
	return db.Raw(sql, values...)
}
