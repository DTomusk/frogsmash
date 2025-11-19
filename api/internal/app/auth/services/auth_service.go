package services

import (
	"context"
	"fmt"
	"frogsmash/internal/app/auth/factories"
	"frogsmash/internal/app/auth/models"
	"frogsmash/internal/app/shared"
	"time"
)

type RefreshTokenRepo interface {
	SaveRefreshToken(token string, userID string, expiresAt int64, ctx context.Context, db shared.DBTX) error
	RevokeTokens(userID string, ctx context.Context, db shared.DBTX) error
	GetRefreshToken(token string, ctx context.Context, db shared.DBTX) (*models.RefreshToken, error)
}

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type TokenService interface {
	GenerateToken(userID string, isVerified bool) (string, error)
}

type AuthService struct {
	userRepo                 UserRepo
	refreshTokenRepo         RefreshTokenRepo
	hasher                   Hasher
	tokenService             TokenService
	refreshTokenLifetimeDays int
}

func NewAuthService(
	userRepo UserRepo,
	refreshTokenRepo RefreshTokenRepo,
	hasher Hasher,
	tokenService TokenService,
	refreshTokenLifetimeDays int) *AuthService {
	return &AuthService{
		userRepo:                 userRepo,
		refreshTokenRepo:         refreshTokenRepo,
		hasher:                   hasher,
		tokenService:             tokenService,
		refreshTokenLifetimeDays: refreshTokenLifetimeDays,
	}
}

func (s *AuthService) Login(email, password string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *models.User, error) {
	// Get the user by email
	user, err := s.userRepo.GetUserByEmail(email, ctx, db)
	if err != nil {
		return "", nil, nil, err
	}
	if user == nil {
		return "", nil, nil, fmt.Errorf("user not found")
	}
	// Verify password
	if !s.hasher.CheckPasswordHash(password, user.PasswordHash) {
		return "", nil, nil, fmt.Errorf("invalid password")
	}
	// Generate JWT token
	jwt, err := s.tokenService.GenerateToken(user.ID, user.IsVerified)
	if err != nil {
		return "", nil, nil, err
	}
	// Generate refresh token
	refreshToken, err := factories.GenerateRefreshToken(user.ID, s.refreshTokenLifetimeDays)
	if err != nil {
		return "", nil, nil, err
	}
	// Save refresh token to database
	if err := s.rotateRefreshTokens(db, ctx, refreshToken); err != nil {
		return "", nil, nil, err
	}

	return jwt, refreshToken, user, nil
}

func (s *AuthService) RefreshToken(refreshToken string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *models.User, error) {
	// Get existing token (the one matching the provided refresh token)
	token, err := s.refreshTokenRepo.GetRefreshToken(refreshToken, ctx, db)
	if err != nil {
		return "", nil, nil, err
	}
	// Validate the token (check not revoked, check not expired)
	if token == nil || token.Revoked || token.ExpiresAt.Before(time.Now()) {
		return "", nil, nil, fmt.Errorf("invalid refresh token")
	}
	user, err := s.userRepo.GetUserByUserID(token.UserID, ctx, db)
	if err != nil {
		return "", nil, nil, err
	}
	// Generate JWT token
	jwt, err := s.tokenService.GenerateToken(user.ID, user.IsVerified)
	if err != nil {
		return "", nil, nil, err
	}
	// Generate refresh token
	newRefreshToken, err := factories.GenerateRefreshToken(user.ID, s.refreshTokenLifetimeDays)
	if err != nil {
		return "", nil, nil, err
	}
	// Save refresh token to database
	if err := s.rotateRefreshTokens(db, ctx, newRefreshToken); err != nil {
		return "", nil, nil, err
	}
	return jwt, newRefreshToken, user, nil
}

func (s *AuthService) rotateRefreshTokens(db shared.TxStarter, ctx context.Context, token *models.RefreshToken) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = s.refreshTokenRepo.RevokeTokens(token.UserID, ctx, tx)
	if err != nil {
		return err
	}
	err = s.refreshTokenRepo.SaveRefreshToken(token.Token, token.UserID, token.ExpiresAt.Unix(), ctx, tx)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
