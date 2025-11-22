package repos

import (
	"context"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/app/verification/models"
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

func (r *VerificationRepo) GetVerificationCode(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error) {
	row := db.QueryRowContext(ctx,
		"SELECT user_id, code, expires_at FROM verification_codes WHERE code = $1",
		code,
	)
	var vc models.VerificationCode
	err := row.Scan(&vc.UserID, &vc.Code, &vc.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &vc, nil
}
