package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
