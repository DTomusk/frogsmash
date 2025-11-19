package services

import (
	"context"
	"frogsmash/internal/app/factories"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
)

type VerificationRepo interface {
	SaveVerificationCode(code *models.VerificationCode, ctx context.Context, db repos.DBTX) error
	DeleteVerificationCodesForUser(userID string, ctx context.Context, db repos.DBTX) error
}

type EmailService interface {
	SendVerificationEmail(toEmail, verificationCode string) error
}

type VerificationService struct {
	userRepo                        UserRepo
	verificationRepo                VerificationRepo
	emailService                    EmailService
	verificationCodeLength          int
	verificationCodeLifetimeMinutes int
}

func NewVerificationService(userRepo UserRepo, verificationRepo VerificationRepo, emailService EmailService, verificationCodeLength int, verificationCodeLifetimeMinutes int) *VerificationService {
	return &VerificationService{
		userRepo:                        userRepo,
		verificationRepo:                verificationRepo,
		emailService:                    emailService,
		verificationCodeLength:          verificationCodeLength,
		verificationCodeLifetimeMinutes: verificationCodeLifetimeMinutes,
	}
}

func (s *VerificationService) ResendVerificationEmail(userID string, ctx context.Context, db repos.DBWithTxStarter) error {
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

func (s *VerificationService) GenerateAndSend(user *models.User, ctx context.Context, db repos.DBTX) error {
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
