package http

import (
	"frogsmash/internal/app"
	"frogsmash/internal/container"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type EventsService interface {
	LogEvent(winnerId, loserId string) error
}

type EventsHandler struct {
	EventsService EventsService
}

func NewEventsHandler(c *container.Container) *EventsHandler {
	return &EventsHandler{
		EventsService: c.EventsService,
	}
}

func SetupRoutes(c *container.Container) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	eventsHandler := NewEventsHandler(c)
	r.GET("/items", GetItems)

	r.POST("/compare", eventsHandler.CompareItems)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	WinnerId string `json:"winner_id"`
	LoserId  string `json:"loser_id"`
}

func (h *EventsHandler) CompareItems(ctx *gin.Context) {
	var request CompareRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Call service function
	// Service checks that IDs have associated items
	// Then it queues a message to update scores asynchronously
	err := h.EventsService.LogEvent(
		request.WinnerId,
		request.LoserId,
	)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to log event"})
		return
	}

	// Placeholder implementation
	ctx.JSON(200, gin.H{
		"status": "comparison recorded",
	})
}
