package services

import (
	"context"
	"fmt"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/app/user/models"
)

type UserFactory interface {
	CreateNewUser(email, password string) (*models.User, error)
}

type UserRepo interface {
	GetUserByUserID(userID, tenantID string, ctx context.Context, db shared.DBTX) (*models.User, error)
	GetUserByEmail(email, tenantID string, ctx context.Context, db shared.DBTX) (*models.User, error)
	CreateUser(user *models.User, tenantID string, ctx context.Context, db shared.DBTX) error
	SetUserIsVerified(userID, tenantID string, isVerified bool, ctx context.Context, db shared.DBTX) error
}

type UserService interface {
	CreateNewUser(email, password, tenantID string, ctx context.Context, db shared.DBWithTxStarter) (string, error)
	GetUserByEmail(email, tenantID string, ctx context.Context, db shared.DBTX) (*models.User, error)
	GetUserByUserID(userID, tenantID string, ctx context.Context, db shared.DBTX) (*models.User, error)
	GetUserEmail(userID, tenantID string, ctx context.Context, db shared.DBTX) (string, error)
	SetUserIsVerified(userID, tenantID string, isVerified bool, ctx context.Context, db shared.DBTX) error
}

type userService struct {
	userFactory UserFactory
	userRepo    UserRepo
}

func NewUserService(userFactory UserFactory, userRepo UserRepo) UserService {
	return &userService{
		userFactory: userFactory,
		userRepo:    userRepo,
	}
}

func (s *userService) CreateNewUser(email, hashedPassword, tenantID string, ctx context.Context, db shared.DBWithTxStarter) (string, error) {
	// Check email not in use
	existingUser, err := s.userRepo.GetUserByEmail(email, tenantID, ctx, db)
	if err != nil {
		return "", err
	}
	if existingUser != nil {
		return "", fmt.Errorf("email %s is already taken", email)
	}

	newUser, err := s.userFactory.CreateNewUser(email, hashedPassword)
	if err != nil {
		return "", err
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer tx.Rollback()

	err = s.userRepo.CreateUser(newUser, tenantID, ctx, tx)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return newUser.ID, nil
}

func (s *userService) GetUserByEmail(email, tenantID string, ctx context.Context, db shared.DBTX) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email, tenantID, ctx, db)
}

func (s *userService) GetUserByUserID(userID, tenantID string, ctx context.Context, db shared.DBTX) (*models.User, error) {
	return s.userRepo.GetUserByUserID(userID, tenantID, ctx, db)
}

func (s *userService) GetUserEmail(userID, tenantID string, ctx context.Context, db shared.DBTX) (string, error) {
	user, err := s.userRepo.GetUserByUserID(userID, tenantID, ctx, db)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("user with ID %s not found", userID)
	}
	return user.Email, nil
}

func (s *userService) SetUserIsVerified(userID, tenantID string, isVerified bool, ctx context.Context, db shared.DBTX) error {
	return s.userRepo.SetUserIsVerified(userID, tenantID, isVerified, ctx, db)
}
