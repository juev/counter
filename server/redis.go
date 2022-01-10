package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// IRedis ...
type IRedis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
}

// Redis ...
type Redis struct {
	// Rdb ...
	Rdb *redis.Client
}

// NewRedisServer ...
func NewRedisServer() IRedis {
	r := Redis{
		Rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}

	return &r
}

// Get ...
func (r *Redis) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.Rdb.Get(ctx, key)
}

// Set ...
func (r *Redis) Set(ctx context.Context, key string, value interface{}) *redis.StatusCmd {
	return r.Rdb.Set(ctx, key, value, 0)
}

// Del ...
func (r *Redis) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.Rdb.Del(ctx, keys...)
}

// Incr ...
func (r *Redis) Incr(ctx context.Context, key string) *redis.IntCmd {
	return r.Rdb.Incr(ctx, key)
}
