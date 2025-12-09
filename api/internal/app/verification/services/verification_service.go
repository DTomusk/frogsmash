package services

import (
	"context"
	"errors"
	"frogsmash/internal/app/shared"
	user "frogsmash/internal/app/user/models"
	"frogsmash/internal/app/verification/factories"
	"frogsmash/internal/app/verification/models"
	"frogsmash/internal/infrastructure/messages"
	"time"
)

var (
	ErrInvalidVerificationCode = errors.New("invalid verification code")
	ErrAlreadyVerified         = errors.New("user is already verified")
)

type UserService interface {
	GetUserEmail(userID, tenantID string, ctx context.Context, db shared.DBTX) (string, error)
	GetUserByEmail(email, tenantID string, ctx context.Context, db shared.DBTX) (*user.User, error)
	SetUserIsVerified(userID, tenantID string, isVerified bool, ctx context.Context, db shared.DBTX) error
}

type VerificationRepo interface {
	SaveVerificationCode(code *models.VerificationCode, ctx context.Context, db shared.DBTX) error
	DeleteVerificationCodesForUser(userID string, ctx context.Context, db shared.DBTX) error
	GetVerificationCode(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error)
	IsUserVerified(userID string, ctx context.Context, db shared.DBTX) (bool, error)
}

type EmailService interface {
	SendVerificationEmail(toEmail, verificationCode string) error
}

type VerificationService interface {
	ResendVerificationEmail(userID, tenantID string, ctx context.Context, db shared.DBWithTxStarter) error
	ResendVerificationEmailToEmail(email, tenantID string, ctx context.Context, db shared.DBWithTxStarter) error
	GenerateAndSend(userID, email string, ctx context.Context, db shared.DBTX) error
	VerifyUser(code, userID, tenantID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error
	IsUserVerified(userID string, ctx context.Context, db shared.DBTX) (bool, error)
	messages.MessageHandler
}

type verificationService struct {
	userService                     UserService
	verificationRepo                VerificationRepo
	emailService                    EmailService
	verificationCodeLength          int
	verificationCodeLifetimeMinutes int
}

func NewVerificationService(userService UserService, verificationRepo VerificationRepo, emailService EmailService, verificationCodeLength int, verificationCodeLifetimeMinutes int) VerificationService {
	return &verificationService{
		userService:                     userService,
		verificationRepo:                verificationRepo,
		emailService:                    emailService,
		verificationCodeLength:          verificationCodeLength,
		verificationCodeLifetimeMinutes: verificationCodeLifetimeMinutes,
	}
}

func (s *verificationService) ResendVerificationEmail(userID, tenantID string, ctx context.Context, db shared.DBWithTxStarter) error {
	email, err := s.userService.GetUserEmail(userID, tenantID, ctx, db)
	if err != nil {
		return err
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	err = s.GenerateAndSend(userID, email, ctx, tx)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *verificationService) ResendVerificationEmailToEmail(email, tenantID string, ctx context.Context, db shared.DBWithTxStarter) error {
	user, err := s.userService.GetUserByEmail(email, tenantID, ctx, db)
	if err != nil || user == nil {
		return err
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	err = s.GenerateAndSend(user.ID, email, ctx, tx)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// TODO: this might know too much about messaging infrastructure
func (s *verificationService) HandleMessage(ctx context.Context, values map[string]interface{}, db shared.DBWithTxStarter) error {
	userID := values["user_id"].(string)
	email := values["email"].(string)
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	err = s.GenerateAndSend(userID, email, ctx, tx)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *verificationService) GenerateAndSend(userID, email string, ctx context.Context, db shared.DBTX) error {
	// Create verification code and send verification email
	verificationCode, err := factories.GenerateVerificationCode(userID, s.verificationCodeLength, s.verificationCodeLifetimeMinutes)
	if err != nil {
		return err
	}

	err = s.verificationRepo.DeleteVerificationCodesForUser(userID, ctx, db)
	if err != nil {
		return err
	}

	err = s.verificationRepo.SaveVerificationCode(verificationCode, ctx, db)
	if err != nil {
		return err
	}
	// TODO: Save verification code to database and send email
	// For now, send email directly
	err = s.emailService.SendVerificationEmail(email, verificationCode.Code)
	return err
}

func (s *verificationService) VerifyUser(code, loggedInUserID, tenantID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error {
	if loggedInUserID == "" {
		return s.verifyAnonymous(code, tenantID, ctx, db)
	}
	return s.verifyLoggedIn(code, loggedInUserID, tenantID, isVerified, ctx, db)
}

func (s *verificationService) verifyAnonymous(code, tenantID string, ctx context.Context, db shared.DBWithTxStarter) error {
	codeModel, err := s.verificationRepo.GetVerificationCode(code, ctx, db)
	if err != nil {
		return err
	}
	if codeModel == nil || codeModel.ExpiresAt.Before(time.Now()) {
		return ErrInvalidVerificationCode
	}

	return s.verificationTransaction(codeModel.UserID, tenantID, ctx, db)
}

func (s *verificationService) verifyLoggedIn(code, loggedInUserID, tenantID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error {
	// Do nothing if the calling user is verified already
	if isVerified {
		return ErrAlreadyVerified
	}

	codeModel, err := s.verificationRepo.GetVerificationCode(code, ctx, db)
	if err != nil {
		return err
	}
	if codeModel == nil || codeModel.ExpiresAt.Before(time.Now()) {
		return ErrInvalidVerificationCode
	}

	// If the logged-in user is the same as the code user, verify directly and expose any errors
	if loggedInUserID == codeModel.UserID {
		return s.verificationTransaction(codeModel.UserID, tenantID, ctx, db)
	}

	// If the logged-in user is different, just verify the code user without exposing errors
	_ = s.verificationTransaction(codeModel.UserID, tenantID, ctx, db)
	// In the case that the logged in user is different, always return invalid code to avoid information leakage
	return ErrInvalidVerificationCode
}

func (s *verificationService) verificationTransaction(userID, tenantID string, ctx context.Context, db shared.DBWithTxStarter) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.userService.SetUserIsVerified(userID, tenantID, true, ctx, tx)
	if err != nil {
		return err
	}

	err = s.verificationRepo.DeleteVerificationCodesForUser(userID, ctx, tx)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *verificationService) IsUserVerified(userID string, ctx context.Context, db shared.DBTX) (bool, error) {
	return s.verificationRepo.IsUserVerified(userID, ctx, db)
}
