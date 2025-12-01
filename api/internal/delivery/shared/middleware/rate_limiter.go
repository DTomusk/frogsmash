package middleware

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type redisClient interface {
	IncrementAndGet(ctx context.Context, key string, expirationSeconds int) (int64, error)
}

type RedisFixedWindowRateLimiter struct {
	client        redisClient
	limit         int
	windowSeconds int
	keyPrefix     string
}

func NewRedisFixedWindowRateLimiter(client redisClient, limit int, windowSeconds int, keyPrefix string) *RedisFixedWindowRateLimiter {
	return &RedisFixedWindowRateLimiter{
		client:        client,
		limit:         limit,
		windowSeconds: windowSeconds,
		keyPrefix:     keyPrefix,
	}
}

func (r *RedisFixedWindowRateLimiter) RateLimitMiddleware(keyFn func(*gin.Context) string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientId := keyFn(ctx)
		key := r.keyPrefix + ":" + clientId

		val, err := r.client.IncrementAndGet(ctx.Request.Context(), key, r.windowSeconds)

		// Fail if redis error
		if err != nil {
			log.Printf("Rate limiter error: %v", err)
			ctx.AbortWithStatus(429)
			return
		}

		if val > int64(r.limit) {
			log.Printf("Rate limiter: client %s exceeded limit with value %d", clientId, val)
			ctx.AbortWithStatus(429)
			return
		}

		ctx.Next()
	}
}
