package services

import (
	"context"
	"fmt"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
)

type UserRepo interface {
	GetUserByUsername(username string, ctx context.Context, db repos.DBTX) (*models.User, error)
	GetUserByEmail(email string, ctx context.Context, db repos.DBTX) (*models.User, error)
	CreateUser(user *models.User, ctx context.Context, db repos.DBTX) error
}

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type AuthService struct {
	UserRepo UserRepo
	Hasher   Hasher
}

func NewAuthService(userRepo UserRepo, hasher Hasher) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
		Hasher:   hasher,
	}
}

func (s *AuthService) RegisterUser(username, email, password string, ctx context.Context, db repos.DBTX) error {
	// Check username not in use
	existingUser, err := s.UserRepo.GetUserByUsername(username, ctx, db)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("username %s is already taken", username)
	}

	// Check email not in use
	existingUser, err = s.UserRepo.GetUserByEmail(email, ctx, db)
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
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword,
	}
	return s.UserRepo.CreateUser(newUser, ctx, db)
}

func (s *AuthService) Login(username, password string, ctx context.Context, db repos.DBTX) (string, error) {
	// Implementation for user login
	return "", nil
}

func (s *AuthService) RefreshToken(refreshToken string, ctx context.Context, db repos.DBTX) (string, error) {
	// Implementation for refreshing JWT token
	return "", nil
}
