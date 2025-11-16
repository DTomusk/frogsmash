package services

type EmailClient interface{}

type EmailService struct {
	EmailClient EmailClient
}

func NewEmailService(emailClient EmailClient) *EmailService {
	return &EmailService{
		EmailClient: emailClient,
	}
}

func (s *EmailService) SendVerificationEmail(toEmail, verificationCode string) error {
	return nil
}
