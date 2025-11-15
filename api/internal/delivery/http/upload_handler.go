package http

import (
	"context"
	"frogsmash/internal/container"
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
		UploadService: c.UploadService,
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
		ctx.JSON(400, gin.H{"error": "No image is received"})
		return
	}

	fileUrl, err := h.UploadService.UploadImage(file, ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Image uploaded successfully",
		"url":     fileUrl,
	})
}
