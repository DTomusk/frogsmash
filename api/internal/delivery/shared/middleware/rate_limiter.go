package middleware

import (
	"frogsmash/internal/infrastructure/redis"

	"github.com/gin-gonic/gin"
)

type RedisFixedWindowRateLimiter struct {
	client        redis.RedisClient
	limit         int
	windowSeconds int
	keyPrefix     string
}

func NewRedisFixedWindowRateLimiter(client redis.RedisClient, limit int, windowSeconds int, keyPrefix string) *RedisFixedWindowRateLimiter {
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
			ctx.AbortWithStatus(429)
			return
		}

		if val > int64(r.limit) {
			ctx.AbortWithStatus(429)
			return
		}

		ctx.Next()
	}
}
