package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Cli struct {
	*redis.Client
}

func Connect(host string, options ...Option) (*Cli, error) {
	o := redis.Options{
		Addr: host,
	}

	for _, f := range options {
		f(&o)
	}

	rdb := redis.NewClient(&o)

	return &Cli{
		rdb,
	}, rdb.Ping(context.Background()).Err()
}
