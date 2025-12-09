package http

import (
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterComparisonRoutes(r *gin.Engine, c *container.APIContainer) {
	comparisonHandler := NewComparisonHandler(c)
	comparison := r.Group("/comparison")

	tenant := comparison.Group("/")
	tenant.Use(middleware.TenantMiddleware())
	{
		tenant.GET("/leaderboard", comparisonHandler.GetLeaderboard)
	}

	protected := comparison.Group("/")
	protected.Use(middleware.AuthMiddleware(c.Auth.JwtService))
	{
		protected.POST("/compare", comparisonHandler.CompareItems)
		protected.GET("/latest-submission", comparisonHandler.GetTimeOfLatestSubmission)
		protected.POST("/submit-contender", comparisonHandler.SubmitContender)
		tenantProtected := protected.Group("/")
		tenantProtected.Use(middleware.TenantMiddleware())
		{
			tenantProtected.GET("/items", comparisonHandler.GetItems)
		}
	}
}
