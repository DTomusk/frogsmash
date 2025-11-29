package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	IncrementAndGet(ctx context.Context, key string, expirationSeconds int) (int64, error)
}

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

// TODO: review how this works when I'm more awake
func (r *redisClient) IncrementAndGet(ctx context.Context, key string, expirationSeconds int) (int64, error) {
	// If key doesn't exist, create and set to 1
	val, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	// Set expiration only when key is newly created
	// Expiration is expirationSeconds after it gets created
	if val == 1 {
		err = r.client.Expire(ctx, key, time.Duration(expirationSeconds)*time.Second).Err()
		if err != nil {
			return 0, err
		}
	}
	return val, nil
}
