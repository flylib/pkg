package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Cli struct {
	*redis.Client
}

func Connect(options ...Option) (*Cli, error) {
	o := newOption()
	for _, f := range options {
		f(o)
	}
	db := redis.NewClient(o)

	err := db.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return &Cli{db}, nil
}
