package mocks

import (
	"context"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/app/verification/models"
)

type MockVerificationRepo struct {
	GetVerificationCodeFunc            func(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error)
	SaveVerificationCodeFunc           func(code *models.VerificationCode, ctx context.Context, db shared.DBTX) error
	DeleteVerificationCodesForUserFunc func(userID string, ctx context.Context, db shared.DBTX) error
	IsUserVerifiedFunc                 func(userID string, ctx context.Context, db shared.DBTX) (bool, error)
}

func (r *MockVerificationRepo) GetVerificationCode(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error) {
	return r.GetVerificationCodeFunc(code, ctx, db)
}

func (r *MockVerificationRepo) SaveVerificationCode(code *models.VerificationCode, ctx context.Context, db shared.DBTX) error {
	return r.SaveVerificationCodeFunc(code, ctx, db)
}

func (r *MockVerificationRepo) DeleteVerificationCodesForUser(userID string, ctx context.Context, db shared.DBTX) error {
	return r.DeleteVerificationCodesForUserFunc(userID, ctx, db)
}

func (r *MockVerificationRepo) IsUserVerified(userID string, ctx context.Context, db shared.DBTX) (bool, error) {
	return r.IsUserVerifiedFunc(userID, ctx, db)
}
