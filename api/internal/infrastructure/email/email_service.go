package email

type EmailClient interface {
	SendEmail(toEmail, subject, htmlBody, textBody string) error
}

type EmailService interface {
	SendVerificationEmail(toEmail, verificationCode string) error
}

type emailService struct {
	emailClient      EmailClient
	templateRenderer *TemplateRenderer
	appUrl           string
}

func NewEmailService(emailClient EmailClient, templateRenderer *TemplateRenderer, appUrl string) *emailService {
	return &emailService{
		emailClient:      emailClient,
		templateRenderer: templateRenderer,
		appUrl:           appUrl,
	}
}

func (s *emailService) SendVerificationEmail(toEmail, verificationCode string) error {
	link := s.appUrl + "/#/verify?code=" + verificationCode
	subject := "FrogSmash - Verify your email"

	data := map[string]string{
		"Subject": subject,
		"Link":    link,
	}

	htmlBody, err := s.templateRenderer.RenderTemplate("verification_email.html", data)
	if err != nil {
		return err
	}

	textBody, err := s.templateRenderer.RenderTemplate("verification_email.txt", data)
	if err != nil {
		return err
	}

	return s.emailClient.SendEmail(toEmail, subject, htmlBody, textBody)
}
