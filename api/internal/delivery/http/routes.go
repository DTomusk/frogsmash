package http

import (
	"frogsmash/internal/app"
	"frogsmash/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(c *container.Container) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/items", GetItems)

	r.POST("/compare", CompareItems)

	return r
}

// Gets two random distinct items for comparison from storage
func GetItems(ctx *gin.Context) {
	item1, item2 := app.PickTwoItems()
	ctx.JSON(200, gin.H{
		"items": []int{item1.Id, item2.Id},
	})
}

type CompareRequest struct {
	WinnerId int `json:"winner_id"`
	LoserId  int `json:"loser_id"`
}

func CompareItems(ctx *gin.Context) {
	var request CompareRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Call service function
	// Service checks that IDs have associated items
	// Then it queues a message to update scores asynchronously

	// Placeholder implementation
	ctx.JSON(200, gin.H{
		"status": "comparison recorded",
	})
}
