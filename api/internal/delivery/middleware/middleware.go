package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Limits the size of all incoming request bodies to maxBytes.
func MaxBodySize(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		c.Next()
	}
}
