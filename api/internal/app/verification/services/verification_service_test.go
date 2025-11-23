package services

import (
	"context"
	"database/sql"
	"errors"
	"frogsmash/internal/app/shared"
	sharedMocks "frogsmash/internal/app/shared/mocks"
	userMocks "frogsmash/internal/app/user/mocks"
	"frogsmash/internal/app/verification/mocks"
	"frogsmash/internal/app/verification/models"
	emailMocks "frogsmash/internal/infrastructure/email/mocks"
	"testing"
	"time"
)

func TestVerifyUser_LoggedIn_AlreadyVerified(t *testing.T) {
	// Arrange
	svc := NewVerificationService(nil, nil, nil, 6, 15)

	// Act
	err := svc.VerifyUser("some-code", "user-id", true, nil, nil)

	// Assert
	if errors.Is(err, ErrAlreadyVerified) == false {
		t.Errorf("Expected error for already verified user, got %v", err)
	}
}

func TestVerifyUser_LoggedIn_NotVerified_InvalidCode(t *testing.T) {
	// Arrange
	mockVerificationRepo := &mocks.MockVerificationRepo{
		GetVerificationCodeFunc: func(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error) {
			return nil, nil
		},
	}

	svc := NewVerificationService(nil, mockVerificationRepo, nil, 6, 15)

	// Act
	err := svc.VerifyUser("invalid-code", "user-id", false, nil, nil)

	// Assert
	if errors.Is(err, ErrInvalidVerificationCode) == false {
		t.Errorf("Expected error for invalid verification code, got %v", err)
	}
}

func TestVerifyUser_LoggedIn_NotVerified_ExpiredCode(t *testing.T) {
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
	err := svc.VerifyUser("expired-code", "user-id", false, nil, nil)

	// Assert
	if errors.Is(err, ErrInvalidVerificationCode) == false {
		t.Errorf("Expected error for expired verification code, got %v", err)
	}
}

func TestVerifyUser_LoggedIn_NotVerified_ValidCode_Success(t *testing.T) {
	// Arrange
	userId := "user-id"
	now := time.Now()
	validCode := &models.VerificationCode{
		Code:      "valid-code",
		UserID:    userId,
		ExpiresAt: now.Add(10 * time.Minute),
	}

	mockVerificationRepo := &mocks.MockVerificationRepo{
		GetVerificationCodeFunc: func(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error) {
			return validCode, nil
		},
		DeleteVerificationCodesForUserFunc: func(userID string, ctx context.Context, db shared.DBTX) error {
			return nil
		},
	}

	mockEmailService := &emailMocks.MockEmailService{
		SendVerificationEmailFunc: func(toEmail, verificationCode string) error {
			return nil
		},
	}

	mockUserService := &userMocks.MockUserService{
		SetUserIsVerifiedFunc: func(userID string, isVerified bool, ctx context.Context, db shared.DBTX) error {
			return nil
		},
	}

	mockDB := &sharedMocks.MockDBWithTxStarter{
		BeginTxFunc: func(ctx context.Context, opts *sql.TxOptions) (shared.Tx, error) {
			return &sharedMocks.MockTx{
				CommitFunc: func() error {
					return nil
				},
				RollbackFunc: func() error {
					return nil
				},
			}, nil
		},
	}

	svc := NewVerificationService(mockUserService, mockVerificationRepo, mockEmailService, 6, 15)

	// Act
	err := svc.VerifyUser("valid-code", userId, false, nil, mockDB)

	// Assert
	if err != nil {
		t.Errorf("Expected no error for valid verification code, got %v", err)
	}
}

func TestVerifyUser_LoggedIn_NotVerified_ValidCode_DifferentUser_Error(t *testing.T) {
	// Arrange
	loggedInUserID := "logged-in-user-id"
	codeUserId := "code-user-id"
	now := time.Now()
	validCode := &models.VerificationCode{
		Code:      "valid-code",
		UserID:    codeUserId,
		ExpiresAt: now.Add(10 * time.Minute),
	}

	mockVerificationRepo := &mocks.MockVerificationRepo{
		GetVerificationCodeFunc: func(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error) {
			return validCode, nil
		},
		DeleteVerificationCodesForUserFunc: func(userID string, ctx context.Context, db shared.DBTX) error {
			return nil
		},
	}

	mockEmailService := &emailMocks.MockEmailService{
		SendVerificationEmailFunc: func(toEmail, verificationCode string) error {
			return nil
		},
	}

	mockUserService := &userMocks.MockUserService{
		SetUserIsVerifiedFunc: func(userID string, isVerified bool, ctx context.Context, db shared.DBTX) error {
			return nil
		},
	}

	mockDB := &sharedMocks.MockDBWithTxStarter{
		BeginTxFunc: func(ctx context.Context, opts *sql.TxOptions) (shared.Tx, error) {
			return &sharedMocks.MockTx{
				CommitFunc: func() error {
					return nil
				},
				RollbackFunc: func() error {
					return nil
				},
			}, nil
		},
	}

	svc := NewVerificationService(mockUserService, mockVerificationRepo, mockEmailService, 6, 15)

	// Act
	err := svc.VerifyUser("valid-code", loggedInUserID, false, nil, mockDB)

	// Assert
	if errors.Is(err, ErrInvalidVerificationCode) == false {
		t.Errorf("Expected invalid code error for different user, got %v", err)
	}
}

func TestVerifyUser_Anonymous_InvalidCode(t *testing.T) {
	// Arrange
	mockVerificationRepo := &mocks.MockVerificationRepo{
		GetVerificationCodeFunc: func(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error) {
			return nil, nil
		},
	}

	svc := NewVerificationService(nil, mockVerificationRepo, nil, 6, 15)

	// Act
	err := svc.VerifyUser("invalid-code", "", false, nil, nil)

	// Assert
	if errors.Is(err, ErrInvalidVerificationCode) == false {
		t.Errorf("Expected error for invalid verification code, got %v", err)
	}
}

func TestVerifyUser_Anonymous_ExpiredCode(t *testing.T) {
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
	err := svc.VerifyUser("expired-code", "", false, nil, nil)

	// Assert
	if errors.Is(err, ErrInvalidVerificationCode) == false {
		t.Errorf("Expected error for expired verification code, got %v", err)
	}
}

func TestVerifyUser_Anonymous_ValidCode_Success(t *testing.T) {
	// Arrange
	now := time.Now()
	validCode := &models.VerificationCode{
		Code:      "valid-code",
		UserID:    "user-id",
		ExpiresAt: now.Add(10 * time.Minute),
	}

	mockVerificationRepo := &mocks.MockVerificationRepo{
		GetVerificationCodeFunc: func(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error) {
			return validCode, nil
		},
		DeleteVerificationCodesForUserFunc: func(userID string, ctx context.Context, db shared.DBTX) error {
			return nil
		},
	}

	mockEmailService := &emailMocks.MockEmailService{
		SendVerificationEmailFunc: func(toEmail, verificationCode string) error {
			return nil
		},
	}

	mockUserService := &userMocks.MockUserService{
		SetUserIsVerifiedFunc: func(userID string, isVerified bool, ctx context.Context, db shared.DBTX) error {
			return nil
		},
	}

	mockDB := &sharedMocks.MockDBWithTxStarter{
		BeginTxFunc: func(ctx context.Context, opts *sql.TxOptions) (shared.Tx, error) {
			return &sharedMocks.MockTx{
				CommitFunc: func() error {
					return nil
				},
				RollbackFunc: func() error {
					return nil
				},
			}, nil
		},
	}

	svc := NewVerificationService(mockUserService, mockVerificationRepo, mockEmailService, 6, 15)

	// Act
	err := svc.VerifyUser("valid-code", "", false, nil, mockDB)

	// Assert
	if err != nil {
		t.Errorf("Expected no error for valid verification code, got %v", err)
	}
}
