package config

type MailConfig struct {
	MailjetAPIKey       string
	MailjetSecretKey    string
	SenderEmail         string
	TemplateGlobPattern string
}

func NewMailConfig() *MailConfig {
	return &MailConfig{
		MailjetAPIKey:       getEnv("MAILJET_API_KEY"),
		MailjetSecretKey:    getEnv("MAILJET_SECRET_KEY"),
		SenderEmail:         getEnv("SENDER_EMAIL"),
		TemplateGlobPattern: getEnv("TEMPLATE_GLOB_PATTERN"),
	}
}
