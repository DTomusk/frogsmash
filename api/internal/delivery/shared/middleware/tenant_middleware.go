package middleware

import "github.com/gin-gonic/gin"

func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Don't use origin in real life, but fine for a hobby app
		// Check that the origin header gets sent
		origin := c.Request.Header.Get("Origin")

		var tenantID string
		// TODO: use real origins in production
		switch origin {
		case "https://frogsmash.co.uk":
			tenantID = "frog"
		case "https://spicklepickle.xyz":
			tenantID = "book"
		default:
			tenantID = "frog"
		}

		c.Set("tenant_id", tenantID)

		c.Next()
	}
}
