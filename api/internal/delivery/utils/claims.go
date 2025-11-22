package utils

import (
	"frogsmash/internal/app/auth/models"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (string, bool) {
	userID, ok := c.Get("sub")
	if !ok {
		return "", false
	}

	idStr, ok := userID.(string)
	return idStr, ok
}

func IsUserVerified(c *gin.Context) (bool, bool) {
	isVerified, ok := c.Get("is_verified")
	if !ok {
		return false, false
	}

	verifiedBool, ok := isVerified.(bool)
	return verifiedBool, ok
}

func GetClaims(c *gin.Context) (*models.Claims, bool) {
	v, ok := c.Get("claims")
	if !ok {
		return &models.Claims{}, false
	}

	claims, ok := v.(*models.Claims)
	return claims, ok
}
