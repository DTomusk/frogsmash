package http

import (
	"context"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/dto"
	"frogsmash/internal/delivery/utils"
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
		ctx.JSON(400, dto.Response{
			Error: "Image file is required",
			Code:  dto.InvalidRequestCode,
		})
		return
	}

	fileUrl, err := h.UploadService.UploadImage(file, ctx)
	if err != nil {
		ctx.JSON(500, dto.Response{
			Error: err.Error(),
			Code:  dto.InternalServerErrorCode,
		})
		return
	}

	ctx.JSON(200, dto.UploadImageResponse{
		URL: fileUrl,
	})
}
