package services

import (
	"context"
	"fmt"
	"frogsmash/internal/app/auth/factories"
	"frogsmash/internal/app/auth/models"
	user "frogsmash/internal/app/user/models"

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

type UserService interface {
	CreateNewUser(email, password string, ctx context.Context, db shared.DBWithTxStarter) (string, error)
	GetUserByEmail(email string, ctx context.Context, db shared.DBTX) (*user.User, error)
	GetUserByUserID(userID string, ctx context.Context, db shared.DBTX) (*user.User, error)
}

type MessageClient interface {
	EnqueueMessage(ctx context.Context, message map[string]interface{}) error
}

type AuthService interface {
	Login(email, password string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error)
	Logout(refreshToken string, ctx context.Context, db shared.DBWithTxStarter) error
	Register(email, password string, ctx context.Context, db shared.DBWithTxStarter) error
	RefreshToken(refreshToken string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error)
}

type authService struct {
	refreshTokenRepo         RefreshTokenRepo
	hasher                   Hasher
	tokenService             TokenService
	userService              UserService
	messageClient            MessageClient
	refreshTokenLifetimeDays int
}

func NewAuthService(
	refreshTokenRepo RefreshTokenRepo,
	hasher Hasher,
	tokenService TokenService,
	userService UserService,
	messageClient MessageClient,
	refreshTokenLifetimeDays int) AuthService {
	return &authService{
		refreshTokenRepo:         refreshTokenRepo,
		hasher:                   hasher,
		tokenService:             tokenService,
		userService:              userService,
		messageClient:            messageClient,
		refreshTokenLifetimeDays: refreshTokenLifetimeDays,
	}
}

func (s *authService) Login(email, password string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error) {
	// Get the user by email
	user, err := s.userService.GetUserByEmail(email, ctx, db)
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

func (s *authService) Logout(refreshToken string, ctx context.Context, db shared.DBWithTxStarter) error {
	token, err := s.refreshTokenRepo.GetRefreshToken(refreshToken, ctx, db)
	if err != nil {
		return err
	}
	if token == nil {
		return fmt.Errorf("refresh token not found")
	}
	if err := s.refreshTokenRepo.RevokeTokens(token.UserID, ctx, db); err != nil {
		return err
	}
	return nil
}

func (s *authService) Register(email, password string, ctx context.Context, db shared.DBWithTxStarter) error {
	hashedPassword, err := s.hasher.HashPassword(password)
	if err != nil {
		return err
	}
	id, err := s.userService.CreateNewUser(email, hashedPassword, ctx, db)
	if err != nil {
		return err
	}

	if err := s.messageClient.EnqueueMessage(ctx, map[string]interface{}{
		"type":    "send_verification_email",
		"user_id": id,
		"email":   email,
	}); err != nil {
		return err
	}

	return nil
}

func (s *authService) RefreshToken(refreshToken string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error) {
	// Get existing token (the one matching the provided refresh token)
	token, err := s.refreshTokenRepo.GetRefreshToken(refreshToken, ctx, db)
	if err != nil {
		return "", nil, nil, err
	}
	// Validate the token (check not revoked, check not expired)
	if token == nil || token.Revoked || token.ExpiresAt.Before(time.Now()) {
		return "", nil, nil, fmt.Errorf("invalid refresh token")
	}
	user, err := s.userService.GetUserByUserID(token.UserID, ctx, db)
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

func (s *authService) rotateRefreshTokens(db shared.TxStarter, ctx context.Context, token *models.RefreshToken) error {
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
