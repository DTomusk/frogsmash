package http

import (
	"context"
	"frogsmash/internal/container"
	sharedDto "frogsmash/internal/delivery/shared/dto"
	"frogsmash/internal/delivery/shared/utils"
	"frogsmash/internal/delivery/upload/dto"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type UploadService interface {
	UploadImage(fileHeader *multipart.FileHeader, ctx context.Context) (string, error)
}

type UploadHandler struct {
	UploadService UploadService
}

func NewUploadHandler(c *container.Container) *UploadHandler {
	return &UploadHandler{
		UploadService: c.InfraServices.UploadService,
	}
}

// UploadImage godoc
// @Summary      Upload an image
// @Description  Uploads an image to the server
// @Router       /upload [post]
// @Accept       multipart/form-data
// @Produce      json
// @Param        image  formData  file  true  "Image file to upload"
func (h *UploadHandler) UploadImage(ctx *gin.Context) {
	isVerified, ok := utils.IsUserVerified(ctx)
	if !ok || !isVerified {
		ctx.JSON(403, gin.H{"error": "User is not verified"})
		return
	}
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(400, sharedDto.Response{
			Error: "Image file is required",
			Code:  sharedDto.InvalidRequestCode,
		})
		return
	}

	// TODO: don't call upload service directly
	// Or maybe do, I don't know
	fileUrl, err := h.UploadService.UploadImage(file, ctx)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: err.Error(),
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}

	ctx.JSON(200, dto.UploadImageResponse{
		URL: fileUrl,
	})
}
