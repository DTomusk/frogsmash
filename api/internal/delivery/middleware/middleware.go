package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Limits the size of all incoming request bodies to maxBytes.
func MaxBodySize(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		c.Next()
	}
}

type TokenService interface {
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// AuthMiddleware verifies the JWT token from the Authorization header. Used for endpoints that require a logged in user.
func AuthMiddleware(s TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := parseTokenFromHeader(c, s)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token subject"})
			return
		}
		c.Set("sub", sub)

		isVerified, ok := claims["is_verified"].(bool)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token verification status"})
			return
		}
		c.Set("is_verified", isVerified)

		c.Next()
	}
}

// OptionalAuthMiddleware attempts to verify the JWT token from the Authorization header.
// If valid, it sets the user info in the context. If missing or invalid, it allows the request to proceed without user info.
func OptionalAuthMiddleware(s TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := parseTokenFromHeader(c, s)
		if ok {
			if sub, ok := claims["sub"].(string); ok {
				c.Set("sub", sub)
			}
			if isVerified, ok := claims["is_verified"].(bool); ok {
				c.Set("is_verified", isVerified)
			}
		}

		c.Next()
	}
}

func parseTokenFromHeader(c *gin.Context, s TokenService) (jwt.MapClaims, bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, false
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := s.ValidateToken(tokenString)
	if err != nil || !token.Valid {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}

	return claims, true
}
