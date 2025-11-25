package mocks

import (
	"context"
	"frogsmash/internal/app/shared"
	user "frogsmash/internal/app/user/models"
)

type MockUserService struct {
	GetUserEmailFunc      func(userID string, ctx context.Context, db shared.DBTX) (string, error)
	SetUserIsVerifiedFunc func(userID string, isVerified bool, ctx context.Context, db shared.DBTX) error
	GetUserByEmailFunc    func(email string, ctx context.Context, db shared.DBTX) (*user.User, error)
}

func (s *MockUserService) GetUserEmail(userID string, ctx context.Context, db shared.DBTX) (string, error) {
	return s.GetUserEmailFunc(userID, ctx, db)
}

func (s *MockUserService) SetUserIsVerified(userID string, isVerified bool, ctx context.Context, db shared.DBTX) error {
	return s.SetUserIsVerifiedFunc(userID, isVerified, ctx, db)
}

func (s *MockUserService) GetUserByEmail(email string, ctx context.Context, db shared.DBTX) (*user.User, error) {
	return s.GetUserByEmailFunc(email, ctx, db)
}
