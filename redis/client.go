package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Cli struct {
	*redis.Client
	opt *redis.Options
}

func Connect(options ...Option) (*Cli, error) {
	opt := newOption()
	for _, f := range options {
		f(opt)
	}
	db := redis.NewClient(opt)

	err := db.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return &Cli{db, opt}, nil
}
