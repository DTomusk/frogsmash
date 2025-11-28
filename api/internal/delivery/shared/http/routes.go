package http

import (
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/shared/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	authHttp "frogsmash/internal/delivery/auth/http"
	comparisonHttp "frogsmash/internal/delivery/comparison/http"
	verificationHttp "frogsmash/internal/delivery/verification/http"
)

func SetupRoutes(c *container.Container) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.MaxBodySize(c.Config.AppConfig.MaxFileSize + 1<<20))

	// TODO: review and inject values, right now use 100 requests per minute per IP
	rateLimiter := middleware.NewRedisFixedWindowRateLimiter(
		c.InfraServices.RedisClient,
		100,
		60,
		"rate_limiter")
	r.Use(rateLimiter.RateLimitMiddleware(func(ctx *gin.Context) string {
		// Use client IP as the key
		return ctx.ClientIP()
	}))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{c.Config.AppConfig.AllowedOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authHttp.RegisterAuthRoutes(r, c)
	comparisonHttp.RegisterComparisonRoutes(r, c)
	verificationHttp.RegisterVerificationRoutes(r, c)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
