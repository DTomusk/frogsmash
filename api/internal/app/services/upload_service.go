package services

import "mime/multipart"

type StorageClient interface {
	UploadFile()
}

type UploadService struct {
	StorageClient StorageClient
}

func NewUploadService(storageClient StorageClient) *UploadService {
	return &UploadService{
		StorageClient: storageClient,
	}
}

func (s *UploadService) UploadImage(fileHeader *multipart.FileHeader) (string, error) {
	// Implement the logic to upload the image using the StorageClient
	// and return the URL of the uploaded image or an error if it fails.
	return "", nil
}
