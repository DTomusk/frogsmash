package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface{}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(ctx context.Context, redisAddress string) (RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return &redisClient{client: rdb}, nil
}
