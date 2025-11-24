package http

import (
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUploadRoutes(r *gin.Engine, c *container.Container) {
	uploadHandler := NewUploadHandler(c)

	upload := r.Group("/upload")
	upload.Use(middleware.AuthMiddleware(c.Auth.JwtService))
	{
		upload.POST("", uploadHandler.UploadImage)
	}
}
