package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewRedisCache(host, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
		PoolSize: 10000,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
