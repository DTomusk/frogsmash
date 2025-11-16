package email

import (
	mj "github.com/mailjet/mailjet-apiv3-go/v4"
)

type MailjetClient struct {
	mjClient *mj.Client
	sender   string
}

func NewMailjetClient(apiKey, secretKey, sender string) *MailjetClient {
	return &MailjetClient{
		mjClient: mj.NewMailjetClient(apiKey, secretKey),
		sender:   sender,
	}
}

func (c *MailjetClient) SendEmail(toEmail, subject, htmlBody, textBody string) error {
	messagesInfo := []mj.InfoMessagesV31{
		{
			From: &mj.RecipientV31{
				Email: c.sender,
				Name:  "FrogSmash",
			},
			To: &mj.RecipientsV31{
				mj.RecipientV31{
					Email: toEmail,
				},
			},
			Subject:  subject,
			HTMLPart: htmlBody,
			TextPart: textBody,
		},
	}

	messages := mj.MessagesV31{Info: messagesInfo}
	_, err := c.mjClient.SendMailV31(&messages)
	return err
}
