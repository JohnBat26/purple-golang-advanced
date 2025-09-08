// Package email for sending email
package email

import (
	"net/smtp"

	"demo/3-validation-api/configs"

	mail "github.com/jordan-wright/email"
)

type EmailSender interface {
	SendEmail(to, subject, body string, config configs.Config) error
}

type SMTPEmailSender struct {
	SMTPAddr string
	Auth     smtp.Auth
}

func (s *SMTPEmailSender) SendEmail(to, subject, body string, config configs.Config) error {
	e := mail.NewEmail()

	e.From = config.Email
	e.To = []string{to}
	e.Subject = subject
	e.HTML = []byte(body)

	return e.Send(s.SMTPAddr, s.Auth)
}
