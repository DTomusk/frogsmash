package repos

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/app/user/models"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetUserByUserID(userID, tenantID string, ctx context.Context, db shared.DBTX) (*models.User, error) {
	query := "SELECT id, email, password_hash, created_at, is_verified FROM users WHERE id = $1 AND tenant_key = $2"
	row := db.QueryRowContext(ctx, query, userID, tenantID)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.IsVerified); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByEmail(email, tenantID string, ctx context.Context, db shared.DBTX) (*models.User, error) {
	query := "SELECT id, email, password_hash, created_at, is_verified FROM users WHERE email = $1 AND tenant_key = $2"
	row := db.QueryRowContext(ctx, query, email, tenantID)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.IsVerified); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User, tenantID string, ctx context.Context, db shared.DBTX) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO users (id, email, password_hash, created_at, tenant_key) VALUES ($1, $2, $3, NOW(), $4)",
		user.ID, user.Email, user.PasswordHash, tenantID,
	)
	return err
}

func (r *UserRepo) SetUserIsVerified(userID, tenantID string, isVerified bool, ctx context.Context, db shared.DBTX) error {
	_, err := db.ExecContext(ctx,
		"UPDATE users SET is_verified = $1 WHERE id = $2 AND tenant_key = $3",
		isVerified, userID, tenantID,
	)
	return err
}

func (r *UserRepo) GetUserEmail(userID, tenantID string, ctx context.Context, db shared.DBTX) (string, error) {
	row := db.QueryRowContext(ctx,
		"SELECT email FROM users WHERE id = $1 AND tenant_key = $2", userID, tenantID)
	var email string
	err := row.Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}
