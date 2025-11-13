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
}

func NewUploadService(storageClient StorageClient) *UploadService {
	return &UploadService{
		StorageClient: storageClient,
	}
}

func (s *UploadService) UploadImage(fileHeader *multipart.FileHeader, ctx context.Context) (string, error) {
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
