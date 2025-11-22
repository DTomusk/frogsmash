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

func (r *UserRepo) GetUserByUserID(userID string, ctx context.Context, db shared.DBTX) (*models.User, error) {
	query := "SELECT id, email, password_hash, created_at, is_verified FROM users WHERE id = $1"
	row := db.QueryRowContext(ctx, query, userID)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.IsVerified); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByEmail(email string, ctx context.Context, db shared.DBTX) (*models.User, error) {
	query := "SELECT id, email, password_hash, created_at, is_verified FROM users WHERE email = $1"
	row := db.QueryRowContext(ctx, query, email)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.IsVerified); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User, ctx context.Context, db shared.DBTX) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO users (id, email, password_hash, created_at) VALUES ($1, $2, $3, NOW())",
		user.ID, user.Email, user.PasswordHash,
	)
	return err
}

func (r *UserRepo) SetUserIsVerified(userID string, isVerified bool, ctx context.Context, db shared.DBTX) error {
	_, err := db.ExecContext(ctx,
		"UPDATE users SET is_verified = $1 WHERE id = $2",
		isVerified, userID,
	)
	return err
}

func (r *UserRepo) GetUserEmail(userID string, ctx context.Context, db shared.DBTX) (string, error) {
	row := db.QueryRowContext(ctx,
		"SELECT email FROM users WHERE id = $1", userID)
	var email string
	err := row.Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}
