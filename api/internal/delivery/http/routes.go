package http

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/dto"
	"frogsmash/internal/delivery/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ItemsService interface {
	GetComparisonItems(ctx context.Context, dbtx repos.DBTX) (*models.Item, *models.Item, error)
	CompareItems(winnerId, loserId string, ctx context.Context, dbtx repos.DBTX) error
	GetLeaderboardPage(limit int, offset int, ctx context.Context, dbtx repos.DBTX) ([]*models.LeaderboardItem, int, error)
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
	r.GET("/leaderboard", itemsHandler.GetLeaderboard)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
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
		"items": dto.GetComparisonItemsResponse{
			LeftItem: dto.ItemDTO{
				ID:       item1.ID,
				Name:     item1.Name,
				ImageURL: item1.ImageURL,
			},
			RightItem: dto.ItemDTO{
				ID:       item2.ID,
				Name:     item2.Name,
				ImageURL: item2.ImageURL,
			},
		},
	})
}

// CompareItems godoc
// @Summary      Compare two items
// @Description  Records the result of a comparison between two items
// @Router       /compare [post]
// @Accept       json
// @Produce      json
// @Param        compareRequest  body      dto.CompareRequest  true  "Comparison Request"
func (h *ItemsHandler) CompareItems(ctx *gin.Context) {
	var request dto.CompareRequest
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

// GetLeaderboard godoc
// @Summary      Get leaderboard
// @Description  Retrieves a paginated leaderboard of items
// @Router       /leaderboard [get]
// @Produce      json
// @Param        page   query     int  false  "Page number"  default(1)
// @Param        limit  query     int  false  "Items per page"  default(10)
func (h *ItemsHandler) GetLeaderboard(ctx *gin.Context) {
	p := utils.NewPagination(ctx)

	items, total, err := h.ItemsService.GetLeaderboardPage(p.Limit, p.Offset, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get leaderboard: " + err.Error()})
		return
	}

	res := dto.NewPagedResponse(items, total, p.Page, p.Limit)

	ctx.JSON(200, res)
}
