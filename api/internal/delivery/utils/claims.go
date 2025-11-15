package utils

import (
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
