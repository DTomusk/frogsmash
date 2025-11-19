package services

import (
	"context"
	"frogsmash/internal/app/auth/factories"
	"frogsmash/internal/app/auth/models"
	"frogsmash/internal/app/shared"
)

type VerificationRepo interface {
	SaveVerificationCode(code *models.VerificationCode, ctx context.Context, db shared.DBTX) error
	DeleteVerificationCodesForUser(userID string, ctx context.Context, db shared.DBTX) error
}

type EmailService interface {
	SendVerificationEmail(toEmail, verificationCode string) error
}

type VerificationService interface {
	ResendVerificationEmail(userID string, ctx context.Context, db shared.DBWithTxStarter) error
	GenerateAndSend(user *models.User, ctx context.Context, db shared.DBTX) error
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
