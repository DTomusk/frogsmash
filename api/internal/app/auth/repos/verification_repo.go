package repos

import (
	"context"
	"frogsmash/internal/app/auth/models"
	"frogsmash/internal/app/shared"
)

type VerificationRepo struct{}

func NewVerificationRepo() *VerificationRepo {
	return &VerificationRepo{}
}

func (r *VerificationRepo) SaveVerificationCode(code *models.VerificationCode, ctx context.Context, db shared.DBTX) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO verification_codes (user_id, code, expires_at) VALUES ($1, $2, $3)",
		code.UserID, code.Code, code.ExpiresAt,
	)
	return err
}

func (r *VerificationRepo) DeleteVerificationCodesForUser(userID string, ctx context.Context, db shared.DBTX) error {
	_, err := db.ExecContext(ctx,
		"DELETE FROM verification_codes WHERE user_id = $1",
		userID,
	)
	return err
}
