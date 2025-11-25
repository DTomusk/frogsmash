package services

import (
	"context"
	"errors"
	"frogsmash/internal/app/shared"
	user "frogsmash/internal/app/user/models"
	"frogsmash/internal/app/verification/factories"
	"frogsmash/internal/app/verification/models"
	"time"
)

var (
	ErrInvalidVerificationCode = errors.New("invalid verification code")
	ErrAlreadyVerified         = errors.New("user is already verified")
)

type UserService interface {
	GetUserEmail(userID string, ctx context.Context, db shared.DBTX) (string, error)
	GetUserByEmail(email string, ctx context.Context, db shared.DBTX) (*user.User, error)
	SetUserIsVerified(userID string, isVerified bool, ctx context.Context, db shared.DBTX) error
}

type VerificationRepo interface {
	SaveVerificationCode(code *models.VerificationCode, ctx context.Context, db shared.DBTX) error
	DeleteVerificationCodesForUser(userID string, ctx context.Context, db shared.DBTX) error
	GetVerificationCode(code string, ctx context.Context, db shared.DBTX) (*models.VerificationCode, error)
}

type EmailService interface {
	SendVerificationEmail(toEmail, verificationCode string) error
}

type VerificationService interface {
	ResendVerificationEmail(userID string, ctx context.Context, db shared.DBWithTxStarter) error
	ResendVerificationEmailToEmail(email string, ctx context.Context, db shared.DBWithTxStarter) error
	GenerateAndSend(userID, email string, ctx context.Context, db shared.DBTX) error
	VerifyUser(code, userID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error
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

func (s *verificationService) ResendVerificationEmail(userID string, ctx context.Context, db shared.DBWithTxStarter) error {
	email, err := s.userService.GetUserEmail(userID, ctx, db)
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

func (s *verificationService) ResendVerificationEmailToEmail(email string, ctx context.Context, db shared.DBWithTxStarter) error {
	user, err := s.userService.GetUserByEmail(email, ctx, db)
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

func (s *verificationService) VerifyUser(code, loggedInUserID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error {
	if loggedInUserID == "" {
		return s.verifyAnonymous(code, ctx, db)
	}
	return s.verifyLoggedIn(code, loggedInUserID, isVerified, ctx, db)
}

func (s *verificationService) verifyAnonymous(code string, ctx context.Context, db shared.DBWithTxStarter) error {
	codeModel, err := s.verificationRepo.GetVerificationCode(code, ctx, db)
	if err != nil {
		return err
	}
	if codeModel == nil || codeModel.ExpiresAt.Before(time.Now()) {
		return ErrInvalidVerificationCode
	}

	return s.verificationTransaction(codeModel.UserID, ctx, db)
}

func (s *verificationService) verifyLoggedIn(code, loggedInUserID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error {
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
		return s.verificationTransaction(codeModel.UserID, ctx, db)
	}

	// If the logged-in user is different, just verify the code user without exposing errors
	_ = s.verificationTransaction(codeModel.UserID, ctx, db)

	// In the case that the logged in user is different, always return invalid code to avoid information leakage
	return ErrInvalidVerificationCode
}

func (s *verificationService) verificationTransaction(userID string, ctx context.Context, db shared.DBWithTxStarter) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.userService.SetUserIsVerified(userID, true, ctx, tx)
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
