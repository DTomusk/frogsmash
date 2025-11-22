package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type StorageClient interface {
	UploadFile(fileName string, fileHeader *multipart.FileHeader, ctx context.Context) (string, error)
}

type UploadService interface {
	UploadImage(fileHeader *multipart.FileHeader, ctx context.Context) (string, error)
}

type uploadService struct {
	StorageClient StorageClient
	MaxFileSize   int64
}

func NewUploadService(storageClient StorageClient, maxFileSize int64) *uploadService {
	return &uploadService{
		StorageClient: storageClient,
		MaxFileSize:   maxFileSize,
	}
}

func (s *uploadService) UploadImage(fileHeader *multipart.FileHeader, ctx context.Context) (string, error) {
	// TODO: pull this out
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExtensions[ext] {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	// Validate file size
	if fileHeader.Size > s.MaxFileSize {
		return "", fmt.Errorf("file size exceeds the maximum limit of %d bytes", s.MaxFileSize)
	}
	// Record that a file is being uploaded in the db
	// Upload file
	filename := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(fileHeader.Filename))
	fileURL, err := s.StorageClient.UploadFile(filename, fileHeader, ctx)
	if err != nil {
		return "", err
	}
	// Update db record with file URL (or failure if upload failed)
	return fileURL, nil
}
