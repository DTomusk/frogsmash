package services

import (
	"context"
	"crypto/rand"
	"fmt"
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

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type TokenService interface {
	GenerateToken(userID string) (string, error)
}

type AuthService struct {
	UserRepo                 UserRepo
	RefreshTokenRepo         RefreshTokenRepo
	Hasher                   Hasher
	TokenService             TokenService
	RefreshTokenLifetimeDays int
}

func NewAuthService(userRepo UserRepo, refreshTokenRepo RefreshTokenRepo, hasher Hasher, tokenService TokenService, refreshTokenLifetimeDays int) *AuthService {
	return &AuthService{
		UserRepo:                 userRepo,
		RefreshTokenRepo:         refreshTokenRepo,
		Hasher:                   hasher,
		TokenService:             tokenService,
		RefreshTokenLifetimeDays: refreshTokenLifetimeDays,
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
	return s.UserRepo.CreateUser(newUser, ctx, db)
}

func (s *AuthService) Login(email, password string, ctx context.Context, db repos.DBWithTxStarter) (string, string, error) {
	// Get the user by email
	user, err := s.UserRepo.GetUserByEmail(email, ctx, db)
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", fmt.Errorf("user not found")
	}
	// Verify password
	if !s.Hasher.CheckPasswordHash(password, user.PasswordHash) {
		return "", "", fmt.Errorf("invalid password")
	}
	// Generate JWT token
	jwt, err := s.TokenService.GenerateToken(user.ID)
	if err != nil {
		return "", "", err
	}
	// Generate refresh token
	refreshToken, err := generateRefreshToken(user.ID, s.RefreshTokenLifetimeDays)
	if err != nil {
		return "", "", err
	}
	// Save refresh token to database
	if err := s.rotateRefreshTokens(db, ctx, refreshToken); err != nil {
		return "", "", err
	}

	return jwt, refreshToken.Token, nil
}

func (s *AuthService) RefreshToken(refreshToken string, ctx context.Context, db repos.DBWithTxStarter) (string, string, error) {
	// Get existing token (the one matching the provided refresh token)
	token, err := s.RefreshTokenRepo.GetRefreshToken(refreshToken, ctx, db)
	if err != nil {
		return "", "", err
	}
	// Validate the token (check not revoked, check not expired)
	if token == nil || token.Revoked || token.ExpiresAt.Before(time.Now()) {
		return "", "", fmt.Errorf("invalid refresh token")
	}
	user, err := s.UserRepo.GetUserByUserID(token.UserID, ctx, db)
	if err != nil {
		return "", "", err
	}
	// Generate JWT token
	jwt, err := s.TokenService.GenerateToken(user.ID)
	if err != nil {
		return "", "", err
	}
	// Generate refresh token
	newRefreshToken, err := generateRefreshToken(user.ID, s.RefreshTokenLifetimeDays)
	if err != nil {
		return "", "", err
	}
	// Save refresh token to database
	if err := s.rotateRefreshTokens(db, ctx, newRefreshToken); err != nil {
		return "", "", err
	}
	return jwt, newRefreshToken.Token, nil
}

// TODO: consider where this should live
// IT's a small utility function that doesn't get used elsewhere and has no state, so it's fine here for now as a private function
// We probably don't need a service for it
func generateRefreshToken(userID string, tokenLifeTimeDays int) (*models.RefreshToken, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return &models.RefreshToken{
		Token:     fmt.Sprintf("%x", b),
		UserID:    userID,
		ExpiresAt: time.Now().Add(time.Duration(tokenLifeTimeDays) * 24 * time.Hour),
		Revoked:   false,
	}, nil
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
