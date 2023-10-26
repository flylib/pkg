package redis

import "github.com/redis/go-redis/v9"

type Option func(options *redis.Options)

func WithDB(db int) Option {
	return func(o *redis.Options) {
		o.DB = db
	}
}

func WithAuth(user, pwd string) Option {
	return func(o *redis.Options) {
		o.Username = user
		o.Password = pwd
	}
}

// Maximum number of connections allocated by the pool at a given time.
// When zero, there is no limit on the number of connections in the pool.
func WithMaxOpenConns(num int) Option {
	return func(o *redis.Options) {
		o.MaxActiveConns = num
		o.MaxIdleConns = num
	}
}

// The default max idle connections is currently 2.
func WithMinIdleConns(num int) Option {
	return func(o *redis.Options) {
		o.MinIdleConns = num
	}
}
