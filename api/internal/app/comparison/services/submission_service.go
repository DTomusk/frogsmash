package services

import (
	"context"
	"errors"
	"frogsmash/internal/app/comparison/models"
	"frogsmash/internal/app/shared"
	"mime/multipart"
	"time"
)

type SubmissionService interface {
	SubmitContender(userID string, fileHeader *multipart.FileHeader, ctx context.Context, db shared.DBTX) error
	GetTimeOfLatestSubmission(userID string, ctx context.Context, db shared.DBTX) (string, error)
}

type submissionService struct {
	uploadService UploadService
	repo          SubmissionRepo
}

type UploadService interface {
	UploadImage(fileHeader *multipart.FileHeader, ctx context.Context) (string, error)
}

type SubmissionRepo interface {
	GetLatestSubmissionByUser(userID string, ctx context.Context, db shared.DBTX) (*models.ImageUpload, error)
	GetTotalDataUploaded(ctx context.Context, db shared.DBTX) (int64, error)
	GetTimeOfLatestSubmission(userID string, ctx context.Context, db shared.DBTX) (string, error)
}

func NewSubmissionService(uploadService UploadService, repo SubmissionRepo) SubmissionService {
	return &submissionService{
		uploadService: uploadService,
		repo:          repo,
	}
}

func (s *submissionService) SubmitContender(userID string, fileHeader *multipart.FileHeader, ctx context.Context, db shared.DBTX) error {
	latest, err := s.repo.GetLatestSubmissionByUser(userID, ctx, db)
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
	totalData, err := s.repo.GetTotalDataUploaded(ctx, db)
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

func (s *submissionService) GetTimeOfLatestSubmission(userID string, ctx context.Context, db shared.DBTX) (string, error) {
	return s.repo.GetTimeOfLatestSubmission(userID, ctx, db)
}
