package mocks

type MockEmailService struct {
	SendVerificationEmailFunc func(toEmail, verificationCode string) error
}

func (s *MockEmailService) SendVerificationEmail(toEmail, verificationCode string) error {
	return s.SendVerificationEmailFunc(toEmail, verificationCode)
}
