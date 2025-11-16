package services

import (
	"context"
	"fmt"
	"frogsmash/internal/app/factories"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
	"time"
)

type UserRepo interface {
	GetUserByUserID(userID string, ctx context.Context, db repos.DBTX) (*models.User, error)
	GetUserByEmail(email string, ctx context.Context, db repos.DBTX) (*models.User, error)
	CreateUser(user *models.User, ctx context.Context, db repos.DBTX) error
}

type RefreshTokenRepo interface {
	SaveRefreshToken(token string, userID string, expiresAt int64, ctx context.Context, db repos.DBTX) error
	RevokeTokens(userID string, ctx context.Context, db repos.DBTX) error
	GetRefreshToken(token string, ctx context.Context, db repos.DBTX) (*models.RefreshToken, error)
}

type VerificationRepo interface {
	SaveVerificationCode(code *models.VerificationCode, ctx context.Context, db repos.DBTX) error
}

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type TokenService interface {
	GenerateToken(userID string, isVerified bool) (string, error)
}

type AuthService struct {
	UserRepo                        UserRepo
	RefreshTokenRepo                RefreshTokenRepo
	Hasher                          Hasher
	TokenService                    TokenService
	VerificationRepo                VerificationRepo
	RefreshTokenLifetimeDays        int
	VerificationCodeLength          int
	VerificationCodeLifetimeMinutes int
}

func NewAuthService(
	userRepo UserRepo,
	refreshTokenRepo RefreshTokenRepo,
	hasher Hasher,
	tokenService TokenService,
	verificationRepo VerificationRepo,
	refreshTokenLifetimeDays int,
	verificationCodeLength int,
	verificationCodeLifetimeMinutes int) *AuthService {
	return &AuthService{
		UserRepo:                        userRepo,
		RefreshTokenRepo:                refreshTokenRepo,
		Hasher:                          hasher,
		TokenService:                    tokenService,
		VerificationRepo:                verificationRepo,
		VerificationCodeLength:          verificationCodeLength,
		VerificationCodeLifetimeMinutes: verificationCodeLifetimeMinutes,
		RefreshTokenLifetimeDays:        refreshTokenLifetimeDays,
	}
}

func (s *AuthService) RegisterUser(email, password string, ctx context.Context, db repos.DBTX) error {
	// Check email not in use
	existingUser, err := s.UserRepo.GetUserByEmail(email, ctx, db)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("email %s is already taken", email)
	}

	// Hash password
	hashedPassword, err := s.Hasher.HashPassword(password)
	if err != nil {
		return err
	}

	// Store user in database
	newUser := &models.User{
		Email:        email,
		PasswordHash: hashedPassword,
	}

	// TODO: wrap in transaction
	err = s.UserRepo.CreateUser(newUser, ctx, db)
	if err != nil {
		return err
	}

	// Create verification code and send verification email
	verificationCode, err := factories.GenerateVerificationCode(newUser.ID, s.VerificationCodeLength, s.VerificationCodeLifetimeMinutes)
	if err != nil {
		return err
	}

	// TODO: when we save a verification code, we should delete any existing unexpired codes for that user first
	err = s.VerificationRepo.SaveVerificationCode(verificationCode, ctx, db)
	if err != nil {
		return err
	}
	// TODO: Save verification code to database and send email
	// For now, send email directly
	// Later, queue email to be sent by background worker
	// So we need an email service

	return nil
}

func (s *AuthService) Login(email, password string, ctx context.Context, db repos.DBWithTxStarter) (string, *models.RefreshToken, *models.User, error) {
	// Get the user by email
	user, err := s.UserRepo.GetUserByEmail(email, ctx, db)
	if err != nil {
		return "", nil, nil, err
	}
	if user == nil {
		return "", nil, nil, fmt.Errorf("user not found")
	}
	// Verify password
	if !s.Hasher.CheckPasswordHash(password, user.PasswordHash) {
		return "", nil, nil, fmt.Errorf("invalid password")
	}
	// Generate JWT token
	jwt, err := s.TokenService.GenerateToken(user.ID, user.IsVerified)
	if err != nil {
		return "", nil, nil, err
	}
	// Generate refresh token
	refreshToken, err := factories.GenerateRefreshToken(user.ID, s.RefreshTokenLifetimeDays)
	if err != nil {
		return "", nil, nil, err
	}
	// Save refresh token to database
	if err := s.rotateRefreshTokens(db, ctx, refreshToken); err != nil {
		return "", nil, nil, err
	}

	return jwt, refreshToken, user, nil
}

func (s *AuthService) RefreshToken(refreshToken string, ctx context.Context, db repos.DBWithTxStarter) (string, *models.RefreshToken, *models.User, error) {
	// Get existing token (the one matching the provided refresh token)
	token, err := s.RefreshTokenRepo.GetRefreshToken(refreshToken, ctx, db)
	if err != nil {
		return "", nil, nil, err
	}
	// Validate the token (check not revoked, check not expired)
	if token == nil || token.Revoked || token.ExpiresAt.Before(time.Now()) {
		return "", nil, nil, fmt.Errorf("invalid refresh token")
	}
	user, err := s.UserRepo.GetUserByUserID(token.UserID, ctx, db)
	if err != nil {
		return "", nil, nil, err
	}
	// Generate JWT token
	jwt, err := s.TokenService.GenerateToken(user.ID, user.IsVerified)
	if err != nil {
		return "", nil, nil, err
	}
	// Generate refresh token
	newRefreshToken, err := factories.GenerateRefreshToken(user.ID, s.RefreshTokenLifetimeDays)
	if err != nil {
		return "", nil, nil, err
	}
	// Save refresh token to database
	if err := s.rotateRefreshTokens(db, ctx, newRefreshToken); err != nil {
		return "", nil, nil, err
	}
	return jwt, newRefreshToken, user, nil
}

func (s *AuthService) rotateRefreshTokens(db repos.TxStarter, ctx context.Context, token *models.RefreshToken) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = s.RefreshTokenRepo.RevokeTokens(token.UserID, ctx, tx)
	if err != nil {
		return err
	}
	err = s.RefreshTokenRepo.SaveRefreshToken(token.Token, token.UserID, token.ExpiresAt.Unix(), ctx, tx)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
