package http

import (
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, c *container.Container) {
	authHandler := NewAuthHandler(c)

	auth := r.Group("/auth")

	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.POST("/refresh-token", authHandler.RefreshToken)

	protected := auth.Group("/")
	protected.Use(middleware.AuthMiddleware(c.Auth.JwtService))
	{
		protected.GET("/me", authHandler.GetMe)
	}
}
