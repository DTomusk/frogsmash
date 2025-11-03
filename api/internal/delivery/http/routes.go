package http

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/container"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ItemsService interface {
	GetComparisonItems(ctx context.Context, dbtx repos.DBTX) (*models.Item, *models.Item, error)
	CompareItems(winnerId, loserId string, ctx context.Context, dbtx repos.DBTX) error
}

type ItemsHandler struct {
	ItemsService ItemsService
	db           *sql.DB
}

func NewItemsHandler(c *container.Container) *ItemsHandler {
	return &ItemsHandler{
		ItemsService: c.ItemsService,
		db:           c.DB,
	}
}

func SetupRoutes(c *container.Container) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{c.AllowedOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	itemsHandler := NewItemsHandler(c)

	r.GET("/items", itemsHandler.GetItems)
	r.POST("/compare", itemsHandler.CompareItems)

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

// GetItems godoc
// @Summary      Get two items for comparison
// @Description  Retrieves two distinct items for comparison
// @Router       /items [get]
// @Produce      json
func (h *ItemsHandler) GetItems(ctx *gin.Context) {
	item1, item2, err := h.ItemsService.GetComparisonItems(ctx.Request.Context(), h.db)
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
func (h *ItemsHandler) CompareItems(ctx *gin.Context) {
	var request CompareRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := h.ItemsService.CompareItems(
		request.WinnerId,
		request.LoserId,
		ctx.Request.Context(),
		h.db,
	)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "comparison recorded",
	})
}
