package http

import (
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(c *container.Container) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.MaxBodySize(c.Config.AppConfig.MaxFileSize + 1<<20))

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

	itemsHandler := NewItemsHandler(c)
	uploadHandler := NewUploadHandler(c)
	authHandler := NewAuthHandler(c)

	r.GET("/leaderboard", itemsHandler.GetLeaderboard)
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/refresh-token", authHandler.RefreshToken)
	r.POST("/verify", authHandler.VerifyUser)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(c.Auth.JwtService))
	{
		protected.GET("/items", itemsHandler.GetItems)
		protected.POST("/compare", itemsHandler.CompareItems)
		protected.POST("/upload", uploadHandler.UploadImage)
		protected.POST("/resend-verification", authHandler.ResendVerificationEmail)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
