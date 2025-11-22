package mocks

import (
	"context"
	"frogsmash/internal/app/shared"
)

type MockUserService struct {
	GetUserEmailFunc      func(userID string, ctx context.Context, db shared.DBTX) (string, error)
	SetUserIsVerifiedFunc func(userID string, isVerified bool, ctx context.Context, db shared.DBTX) error
}

func (s *MockUserService) GetUserEmail(userID string, ctx context.Context, db shared.DBTX) (string, error) {
	return s.GetUserEmailFunc(userID, ctx, db)
}

func (s *MockUserService) SetUserIsVerified(userID string, isVerified bool, ctx context.Context, db shared.DBTX) error {
	return s.SetUserIsVerifiedFunc(userID, isVerified, ctx, db)
}
