package http

import (
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterComparisonRoutes(r *gin.Engine, c *container.Container) {
	comparisonHandler := NewComparisonHandler(c)
	comparison := r.Group("/comparison")

	comparison.GET("/leaderboard", comparisonHandler.GetLeaderboard)

	protected := comparison.Group("/")
	protected.Use(middleware.AuthMiddleware(c.Auth.JwtService))
	{
		protected.GET("/items", comparisonHandler.GetItems)
		protected.POST("/compare", comparisonHandler.CompareItems)
		protected.GET("/latest-submission", comparisonHandler.GetTimeOfLatestSubmission)
		protected.POST("/submit-contender", comparisonHandler.SubmitContender)
	}
}
