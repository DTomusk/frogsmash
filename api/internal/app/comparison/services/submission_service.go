package services

import (
	"context"
	"errors"
	"frogsmash/internal/app/comparison/models"
	"mime/multipart"
	"time"
)

type SubmissionService interface {
	SubmitContender(userID string, fileHeader *multipart.FileHeader, ctx context.Context) error
	GetTimeOfLatestSubmission(userID string, ctx context.Context) (string, error)
}

type submissionService struct {
	uploadService UploadService
	repo          SubmissionRepo
}

type UploadService interface {
	UploadImage(fileHeader *multipart.FileHeader, ctx context.Context) (string, error)
}

type SubmissionRepo interface {
	GetLatestSubmissionByUser(userID string, ctx context.Context) (*models.ImageUpload, error)
	GetTotalDataUploaded(ctx context.Context) (int64, error)
	GetTimeOfLatestSubmission(userID string, ctx context.Context) (string, error)
}

func NewSubmissionService(uploadService UploadService, repo SubmissionRepo) SubmissionService {
	return &submissionService{
		uploadService: uploadService,
		repo:          repo,
	}
}

func (s *submissionService) SubmitContender(userID string, fileHeader *multipart.FileHeader, ctx context.Context) error {
	latest, err := s.repo.GetLatestSubmissionByUser(userID, ctx)
	if err != nil {
		return err
	}
	if latest != nil && latest.UploadedAt != "" {
		t, err := time.Parse(time.RFC3339, latest.UploadedAt)
		if err != nil {
			return err
		}
		now := time.Now().UTC()
		midnightUTC := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		if t.After(midnightUTC) {
			return errors.New("user has already submitted an image today")
		}
	}
	totalData, err := s.repo.GetTotalDataUploaded(ctx)
	if err != nil {
		return err
	}
	// TODO: inject limit
	if totalData >= 5*1024*1024*1024 {
		return errors.New("total data upload limit reached")
	}
	_, err = s.uploadService.UploadImage(fileHeader, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *submissionService) GetTimeOfLatestSubmission(userID string, ctx context.Context) (string, error) {
	return s.repo.GetTimeOfLatestSubmission(userID, ctx)
}
