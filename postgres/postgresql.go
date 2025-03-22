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
	o := newOption()
	for _, f := range options {
		f(o)
	}

	db, err := gorm.Open(postgres.Open(o.getDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(o.logLevel)),
	})
	if err != nil {
		return nil, err
	}

	//DB is a database handle representing a pool of zero or more underlying connections.
	connPool, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = connPool.Ping()
	if err != nil {
		return nil, err
	}

	if o.maxOpenConns >= o.maxIdleConns {
		connPool.SetMaxOpenConns(o.maxOpenConns)
	}

	if o.maxIdleConns > 0 {
		connPool.SetMaxIdleConns(o.maxIdleConns)
	}

	return &Cli{DB: db}, err
}
