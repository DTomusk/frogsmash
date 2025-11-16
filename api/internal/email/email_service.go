package email

type EmailClient interface {
	SendEmail(toEmail, subject, body string) error
}

type EmailService struct {
	emailClient      EmailClient
	templateRenderer *TemplateRenderer
	appUrl           string
}

func NewEmailService(emailClient EmailClient, templateRenderer *TemplateRenderer, appUrl string) *EmailService {
	return &EmailService{
		emailClient:      emailClient,
		templateRenderer: templateRenderer,
		appUrl:           appUrl,
	}
}

func (s *EmailService) SendVerificationEmail(toEmail, verificationCode string) error {
	link := s.appUrl + "/verify?code=" + verificationCode
	subject := "FrogSmash - Verify your email"

	body, err := s.templateRenderer.RenderTemplate("verification_email", map[string]string{
		"Subject": subject,
		"Link":    link,
	})
	if err != nil {
		return err
	}

	return s.emailClient.SendEmail(toEmail, subject, body)
}
