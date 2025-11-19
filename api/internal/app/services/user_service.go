package services

import (
	"context"
	"fmt"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
)

type UserFactory interface {
	CreateNewUser(email, password string) (*models.User, error)
}

type UserRepo interface {
	GetUserByUserID(userID string, ctx context.Context, db repos.DBTX) (*models.User, error)
	GetUserByEmail(email string, ctx context.Context, db repos.DBTX) (*models.User, error)
	CreateUser(user *models.User, ctx context.Context, db repos.DBTX) error
}

type UserService struct {
	userFactory         UserFactory
	userRepo            UserRepo
	verificationService *VerificationService
}

func NewUserService(userFactory UserFactory, userRepo UserRepo, verificationService *VerificationService) *UserService {
	return &UserService{
		userFactory:         userFactory,
		userRepo:            userRepo,
		verificationService: verificationService,
	}
}

func (s *UserService) RegisterUser(email, password string, ctx context.Context, db repos.DBWithTxStarter) error {
	// Check email not in use
	existingUser, err := s.userRepo.GetUserByEmail(email, ctx, db)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("email %s is already taken", email)
	}

	newUser, err := s.userFactory.CreateNewUser(email, password)
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = s.userRepo.CreateUser(newUser, ctx, tx)
	if err != nil {
		return err
	}

	err = s.verificationService.GenerateAndSend(newUser, ctx, tx)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
