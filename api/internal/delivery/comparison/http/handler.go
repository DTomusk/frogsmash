package http

import (
	"context"
	"frogsmash/internal/app/comparison/models"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/comparison/dto"
	sharedDto "frogsmash/internal/delivery/shared/dto"
	"frogsmash/internal/delivery/shared/utils"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type ComparisonService interface {
	GetComparisonItems(ctx context.Context, db shared.DBTX) (*models.Item, *models.Item, error)
	CompareItems(winnerId, loserId, userId string, ctx context.Context, db shared.DBTX) error
	GetLeaderboardPage(limit int, offset int, ctx context.Context, db shared.DBTX) ([]*models.LeaderboardItem, int, error)
}

type SubmissionService interface {
	SubmitContender(userID string, fileHeader *multipart.FileHeader, ctx context.Context, db shared.DBTX) error
	GetTimeOfLatestSubmission(userID string, ctx context.Context, db shared.DBTX) (string, error)
}

type ComparisonHandler struct {
	ComparisonService ComparisonService
	SubmissionService SubmissionService
	db                shared.DBTX
}

func NewComparisonHandler(c *container.Container) *ComparisonHandler {
	return &ComparisonHandler{
		ComparisonService: c.Comparison.ComparisonService,
		SubmissionService: c.Comparison.SubmissionService,
		db:                c.InfraServices.DB,
	}
}

// GetItems godoc
// @Summary      Get two items for comparison
// @Description  Retrieves two distinct items for comparison
// @Router       /comparison/items [get]
// @Produce      json
func (h *ComparisonHandler) GetItems(ctx *gin.Context) {
	item1, item2, err := h.ComparisonService.GetComparisonItems(ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: "Failed to get items: " + err.Error(),
			Code:  sharedDto.InternalServerErrorCode,
		})
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
// @Router       /comparison/compare [post]
// @Accept       json
// @Produce      json
// @Param        compareRequest  body      dto.CompareRequest  true  "Comparison Request"
func (h *ComparisonHandler) CompareItems(ctx *gin.Context) {
	user_id, ok := utils.GetUserID(ctx)
	if !ok {
		ctx.JSON(401, sharedDto.Response{
			Error: "Unauthorized",
			Code:  sharedDto.UnauthorizedCode,
		})
		return
	}
	var request dto.CompareRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, sharedDto.Response{
			Error: "Invalid request",
			Code:  sharedDto.InvalidRequestCode,
		})
		return
	}

	err := h.ComparisonService.CompareItems(
		request.WinnerId,
		request.LoserId,
		user_id,
		ctx.Request.Context(),
		h.db,
	)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, sharedDto.Response{
		Message: "Comparison recorded successfully",
	})
}

// GetLeaderboard godoc
// @Summary      Get leaderboard
// @Description  Retrieves a paginated leaderboard of items
// @Router       /comparison/leaderboard [get]
// @Produce      json
// @Param        page   query     int  false  "Page number"  default(1)
// @Param        limit  query     int  false  "Items per page"  default(10)
func (h *ComparisonHandler) GetLeaderboard(ctx *gin.Context) {
	p := utils.NewPagination(ctx)

	items, total, err := h.ComparisonService.GetLeaderboardPage(p.Limit, p.Offset, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: "Failed to get leaderboard: " + err.Error(),
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}

	res := sharedDto.NewPagedResponse(items, total, p.Page, p.Limit)

	ctx.JSON(200, res)
}

// UploadImage godoc
// @Summary      Upload an image
// @Description  Uploads an image to the server
// @Router       /comparison/submit-contender [post]
// @Accept       multipart/form-data
// @Produce      json
// @Param        image  formData  file  true  "Image file to upload"
func (h *ComparisonHandler) SubmitContender(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(400, sharedDto.Response{
			Error: "Image file is required",
			Code:  sharedDto.InvalidRequestCode,
		})
		return
	}

	userId, ok := utils.GetUserID(ctx)
	if !ok {
		ctx.JSON(401, sharedDto.Response{
			Error: "Unauthorized",
			Code:  sharedDto.UnauthorizedCode,
		})
		return
	}

	err = h.SubmissionService.SubmitContender(userId, file, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: err.Error(),
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}

	// TODO: add response dto
	ctx.JSON(200, sharedDto.Response{
		Message: "Image submitted successfully",
	})
}

// GetTimeOfLatestSubmission godoc
// @Summary      Get time of latest submission
// @Description  Retrieves the time of the latest submission by the user
// @Router       /comparison/latest-submission [get]
// @Produce      json
func (h *ComparisonHandler) GetTimeOfLatestSubmission(ctx *gin.Context) {
	userId, ok := utils.GetUserID(ctx)
	if !ok {
		ctx.JSON(401, sharedDto.Response{
			Error: "Unauthorized",
			Code:  sharedDto.UnauthorizedCode,
		})
		return
	}

	uploadedAt, err := h.SubmissionService.GetTimeOfLatestSubmission(userId, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: "Failed to get latest submission time: " + err.Error(),
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}

	// TODO: add response dto
	ctx.JSON(200, dto.GetLatestSubmissionResponse{
		UploadedAt: uploadedAt,
	})
}
