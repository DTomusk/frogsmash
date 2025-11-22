package services

import (
	"context"
	"errors"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/app/verification/mocks"
	"frogsmash/internal/app/verification/models"
	"testing"
	"time"
)

func TestVerifyLoggedIn_AlreadyVerified(t *testing.T) {
	// Arrange
	svc := NewVerificationService(nil, nil, nil, 6, 15)

	// Act
	err := svc.VerifyLoggedIn("some-code", "user-id", true, nil, nil)

	// Assert
	if errors.Is(err, ErrAlreadyVerified) == false {
		t.Errorf("Expected error for already verified user, got %v", err)
	}
}

func TestVerifyLoggedIn_NotVerified_InvalidCode(t *testing.T) {
	// Arrange
	mockVerificationRepo := &mocks.MockVerificationRepo{
		GetVerificationCodeFunc: func(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error) {
			return nil, nil
		},
	}

	svc := NewVerificationService(nil, mockVerificationRepo, nil, 6, 15)

	// Act
	err := svc.VerifyLoggedIn("invalid-code", "user-id", false, nil, nil)

	// Assert
	if errors.Is(err, ErrInvalidVerificationCode) == false {
		t.Errorf("Expected error for invalid verification code, got %v", err)
	}
}

func TestVerifyLoggedIn_NotVerified_ExpiredCode(t *testing.T) {
	// Arrange
	now := time.Now()
	expiredCode := &models.VerificationCode{
		Code:      "expired-code",
		UserID:    "user-id",
		ExpiresAt: now.Add(-1 * time.Minute),
	}

	mockVerificationRepo := &mocks.MockVerificationRepo{
		GetVerificationCodeFunc: func(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error) {
			return expiredCode, nil
		},
	}

	svc := NewVerificationService(nil, mockVerificationRepo, nil, 6, 15)

	// Act
	err := svc.VerifyLoggedIn("expired-code", "user-id", false, nil, nil)

	// Assert
	if errors.Is(err, ErrInvalidVerificationCode) == false {
		t.Errorf("Expected error for expired verification code, got %v", err)
	}
}
