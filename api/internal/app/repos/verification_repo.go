package repos

import (
	"context"
	"frogsmash/internal/app/models"
)

type VerificationRepo struct{}

func NewVerificationRepo() *VerificationRepo {
	return &VerificationRepo{}
}

func (r *VerificationRepo) SaveVerificationCode(code *models.VerificationCode, ctx context.Context, db DBTX) error {
	// Implementation here
	return nil
}

func (r *VerificationRepo) DeleteVerificationCodesForUser(userID string, ctx context.Context, db DBTX) error {
	// Implementation here
	return nil
}
