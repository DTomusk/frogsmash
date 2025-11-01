package http

import (
	"frogsmash/internal/app/models"
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

type ItemsService interface {
	GetComparisonItems() (*models.Item, *models.Item, error)
}

type ItemsHandler struct {
	ItemsService ItemsService
}

func NewItemsHandler(c *container.Container) *ItemsHandler {
	return &ItemsHandler{
		ItemsService: c.ItemsService,
	}
}

func SetupRoutes(c *container.Container) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	itemsHandler := NewItemsHandler(c)
	eventsHandler := NewEventsHandler(c)
	r.GET("/items", itemsHandler.GetItems)

	r.POST("/compare", eventsHandler.CompareItems)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

type GetComparisonItemsResponse struct {
	LeftItem  ItemDTO `json:"left_item"`
	RightItem ItemDTO `json:"right_item"`
}

type ItemDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

// Gets two random distinct items for comparison from storage
// TODO: define return type
// GetItems godoc
// @Summary      Get two items for comparison
// @Description  Retrieves two distinct items for comparison
// @Router       /items [get]
// @Produce      json
func (h *ItemsHandler) GetItems(ctx *gin.Context) {
	item1, item2, err := h.ItemsService.GetComparisonItems()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get items"})
		return
	}
	ctx.JSON(200, gin.H{
		"items": GetComparisonItemsResponse{
			LeftItem: ItemDTO{
				ID:       item1.ID,
				Name:     item1.Name,
				ImageURL: item1.ImageURL,
			},
			RightItem: ItemDTO{
				ID:       item2.ID,
				Name:     item2.Name,
				ImageURL: item2.ImageURL,
			},
		},
	})
}

// CompareRequest godoc
// @Description  Request payload for comparing two items
type CompareRequest struct {
	WinnerId string `json:"winner_id"`
	LoserId  string `json:"loser_id"`
}

// CompareItems godoc
// @Summary      Compare two items
// @Description  Records the result of a comparison between two items
// @Router       /compare [post]
// @Accept       json
// @Produce      json
// @Param        compareRequest  body      CompareRequest  true  "Comparison Request"
func (h *EventsHandler) CompareItems(ctx *gin.Context) {
	var request CompareRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

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
