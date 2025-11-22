package services

import (
	"context"
	"errors"
	"frogsmash/internal/app/auth/factories"
	"frogsmash/internal/app/auth/models"
	"frogsmash/internal/app/shared"
	"time"
)

var (
	ErrInvalidVerificationCode = errors.New("invalid verification code")
	ErrAlreadyVerified         = errors.New("user is already verified")
)

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
	GenerateAndSend(user *models.User, ctx context.Context, db shared.DBTX) error
	VerifyAnonymous(code string, ctx context.Context, db shared.DBWithTxStarter) error
	VerifyLoggedIn(code, userID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error
}

type verificationService struct {
	userRepo                        UserRepo
	verificationRepo                VerificationRepo
	emailService                    EmailService
	verificationCodeLength          int
	verificationCodeLifetimeMinutes int
}

func NewVerificationService(userRepo UserRepo, verificationRepo VerificationRepo, emailService EmailService, verificationCodeLength int, verificationCodeLifetimeMinutes int) VerificationService {
	return &verificationService{
		userRepo:                        userRepo,
		verificationRepo:                verificationRepo,
		emailService:                    emailService,
		verificationCodeLength:          verificationCodeLength,
		verificationCodeLifetimeMinutes: verificationCodeLifetimeMinutes,
	}
}

func (s *verificationService) ResendVerificationEmail(userID string, ctx context.Context, db shared.DBWithTxStarter) error {
	user, err := s.userRepo.GetUserByUserID(userID, ctx, db)
	if err != nil {
		return err
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	err = s.GenerateAndSend(user, ctx, tx)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s *verificationService) GenerateAndSend(user *models.User, ctx context.Context, db shared.DBTX) error {
	// Create verification code and send verification email
	verificationCode, err := factories.GenerateVerificationCode(user.ID, s.verificationCodeLength, s.verificationCodeLifetimeMinutes)
	if err != nil {
		return err
	}

	err = s.verificationRepo.DeleteVerificationCodesForUser(user.ID, ctx, db)
	if err != nil {
		return err
	}

	err = s.verificationRepo.SaveVerificationCode(verificationCode, ctx, db)
	if err != nil {
		return err
	}
	// TODO: Save verification code to database and send email
	// For now, send email directly
	err = s.emailService.SendVerificationEmail(user.Email, verificationCode.Code)
	return err
}

func (s *verificationService) VerifyAnonymous(code string, ctx context.Context, db shared.DBWithTxStarter) error {
	codeModel, err := s.verificationRepo.GetVerificationCode(code, ctx, db)
	if err != nil {
		return err
	}
	if codeModel == nil || codeModel.ExpiresAt.Before(time.Now()) {
		return ErrInvalidVerificationCode
	}

	user, err := s.userRepo.GetUserByUserID(codeModel.UserID, ctx, db)
	if err != nil {
		return err
	}

	return s.verifyUser(user.ID, ctx, db)
}

func (s *verificationService) VerifyLoggedIn(code, loggedInUserID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error {
	codeModel, err := s.verificationRepo.GetVerificationCode(code, ctx, db)
	if err != nil {
		return err
	}
	if codeModel == nil || codeModel.ExpiresAt.Before(time.Now()) {
		return ErrInvalidVerificationCode
	}

	codeUser, err := s.userRepo.GetUserByUserID(codeModel.UserID, ctx, db)
	if err != nil {
		return err
	}

	// Do nothing if the calling user is verified already
	if isVerified {
		return ErrAlreadyVerified
	}

	// If the logged-in user is the same as the code user, verify directly and expose any errors
	if loggedInUserID == codeUser.ID {
		return s.verifyUser(codeUser.ID, ctx, db)
	}

	// If the logged-in user is different, just verify the code user without exposing errors
	_ = s.verifyUser(codeUser.ID, ctx, db)

	// In the case that the logged in user is different, always return invalid code to avoid information leakage
	return ErrInvalidVerificationCode
}

func (s *verificationService) verifyUser(userID string, ctx context.Context, db shared.DBWithTxStarter) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.userRepo.SetUserIsVerified(userID, true, ctx, tx)
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
