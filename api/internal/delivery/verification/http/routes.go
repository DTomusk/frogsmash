package http

import (
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterVerificationRoutes(r *gin.Engine, c *container.Container) {
	verificationHandler := NewVerificationHandler(c)

	verification := r.Group("/verify")
	verification.POST("/resend-email-anonymous",
		verificationHandler.ResendVerificationEmailAnonymous,
	)

	verification.POST("",
		middleware.OptionalAuthMiddleware(c.Auth.JwtService),
		verificationHandler.VerifyUser,
	)

	verification.POST("/resend-email",
		middleware.AuthMiddleware(c.Auth.JwtService),
		verificationHandler.ResendVerificationEmail,
	)
}
