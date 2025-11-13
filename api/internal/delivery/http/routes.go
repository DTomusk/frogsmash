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
	r.Use(middleware.MaxBodySize(c.MaxRequestSize + 1<<20))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{c.AllowedOrigin},
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

	r.GET("/items", itemsHandler.GetItems)
	r.POST("/compare", itemsHandler.CompareItems)
	r.GET("/leaderboard", itemsHandler.GetLeaderboard)

	uploadHandler := NewUploadHandler(c)
	r.POST("/upload", uploadHandler.UploadImage)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
