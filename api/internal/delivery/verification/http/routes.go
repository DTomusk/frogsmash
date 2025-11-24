package http

import (
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterVerificationRoutes(r *gin.Engine, c *container.Container) {
	verificationHandler := NewVerificationHandler(c)

	verification := r.Group("/verify")

	optional := verification.Group("/")
	optional.Use(middleware.OptionalAuthMiddleware(c.Auth.JwtService))
	{
		optional.POST("", verificationHandler.VerifyUser)
	}

	protected := verification.Group("/")
	protected.Use(middleware.AuthMiddleware(c.Auth.JwtService))
	{
		protected.POST("/resend-email", verificationHandler.ResendVerificationEmail)
	}
}
