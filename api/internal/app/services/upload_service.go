package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
)

type StorageClient interface {
	UploadFile(fileName string, fileHeader *multipart.FileHeader, ctx context.Context) (string, error)
}

type UploadService struct {
	StorageClient StorageClient
	MaxFileSize   int64
}

func NewUploadService(storageClient StorageClient, maxFileSize int64) *UploadService {
	return &UploadService{
		StorageClient: storageClient,
		MaxFileSize:   maxFileSize,
	}
}

func (s *UploadService) UploadImage(fileHeader *multipart.FileHeader, ctx context.Context) (string, error) {
	// TODO: pull this out
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	ext := filepath.Ext(fileHeader.Filename)
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
