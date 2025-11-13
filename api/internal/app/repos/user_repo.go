package repos

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/models"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetUserByEmail(email string, ctx context.Context, db DBTX) (*models.User, error) {
	query := "SELECT id, username, email, password_hash, created_at FROM users WHERE email = $1"
	row := db.QueryRowContext(ctx, query, email)
	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByUsername(username string, ctx context.Context, db DBTX) (*models.User, error) {
	query := "SELECT id, username, email, password_hash, created_at FROM users WHERE username = $1"
	row := db.QueryRowContext(ctx, query, username)
	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User, ctx context.Context, db DBTX) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO users (username, email, password_hash, created_at) VALUES ($1, $2, $3, NOW())",
		user.Username, user.Email, user.PasswordHash,
	)
	return err
}
