package http

import (
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, c *container.APIContainer) {
	authHandler := NewAuthHandler(c)

	auth := r.Group("/auth")

	tenant := auth.Group("/")
	tenant.Use(middleware.TenantMiddleware())
	{
		tenant.POST("/register", authHandler.Register)
		tenant.POST("/login", authHandler.Login)
		tenant.POST("/logout", authHandler.Logout)
		tenant.POST("/refresh-token", authHandler.RefreshToken)
		tenant.POST("/google-login", authHandler.GoogleLogin)
		protected := tenant.Group("/")
		protected.Use(middleware.AuthMiddleware(c.Auth.JwtService))
		{
			protected.GET("/me", authHandler.GetMe)
		}
	}
}
