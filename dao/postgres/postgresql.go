package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Cli struct {
	*gorm.DB
}

func Connect(options ...Option) (*Cli, error) {
	o := option{}
	for _, f := range options {
		f(&o)
	}
	if int(o.logLevel) == 0 {
		o.logLevel = Error
	}

	db, err := gorm.Open(postgres.Open(o.getDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(o.logLevel)),
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

	if o.maxOpenConns >= o.maxIdleConns {
		sqlDB.SetMaxOpenConns(o.maxOpenConns)
	}

	if o.maxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(o.maxIdleConns)
	}

	return &Cli{DB: db}, err
}
