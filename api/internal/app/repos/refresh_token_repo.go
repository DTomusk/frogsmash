package repos

import (
	"context"
	"frogsmash/internal/app/models"
)

type RefreshTokenRepo struct{}

func NewRefreshTokenRepo() *RefreshTokenRepo {
	return &RefreshTokenRepo{}
}

func (r *RefreshTokenRepo) SaveRefreshToken(token string, userID string, expiresAt int64, ctx context.Context, db DBTX) error {
	// TODO: do a transaction to revoke all existing tokens for the user
	_, err := db.ExecContext(ctx, "INSERT INTO refresh_tokens (token, user_id, expires_at) VALUES ($1, $2, to_timestamp($3))", token, userID, expiresAt)
	return err
}

func (r *RefreshTokenRepo) RevokeTokens(userID string, ctx context.Context, db DBTX) error {
	_, err := db.ExecContext(ctx, "UPDATE refresh_tokens SET revoked = TRUE WHERE user_id = $1", userID)
	return err
}

func (r *RefreshTokenRepo) GetRefreshToken(token string, ctx context.Context, db DBTX) (*models.RefreshToken, error) {
	row := db.QueryRowContext(ctx, "SELECT token, user_id, expires_at, revoked FROM refresh_tokens WHERE token = $1", token)
	var rt models.RefreshToken
	err := row.Scan(&rt.Token, &rt.UserID, &rt.ExpiresAt, &rt.Revoked)
	if err != nil {
		return nil, err
	}
	return &rt, nil
}
