package redis

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type Option func(options *redis.Options)

func newOption() *redis.Options {
	return &redis.Options{
		Addr: "127.0.0.1:6379",
	}
}

func WithDB(db int) Option {
	return func(o *redis.Options) {
		o.DB = db
	}
}

func WithDialTimeout(timeout time.Duration) Option {
	return func(o *redis.Options) {
		o.DialTimeout = timeout
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(o *redis.Options) {
		o.ReadTimeout = timeout
	}
}
func WriteTimeout(timeout time.Duration) Option {
	return func(o *redis.Options) {
		o.WriteTimeout = timeout
	}
}

func WithAuth(user, password string) Option {
	return func(o *redis.Options) {
		o.Username = user
		o.Password = password
	}
}

// Maximum number of connections allocated by the pool at a given time.
// When zero, there is no limit on the number of connections in the pool.
func WithMaxOpenConns(num int) Option {
	return func(o *redis.Options) {
		o.MaxActiveConns = num
	}
}

// The default max idle connections is currently 2.
func WithMinIdleConns(num int) Option {
	return func(o *redis.Options) {
		o.MinIdleConns = num
	}
}
