package service

import (
	"crypto/tls"
	"fmt"
	emailConfig "rest_api/internal/apps/email/config"

	gomail "gopkg.in/mail.v2"
)

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}

type EmailServiceImpl struct {
	config *emailConfig.EmailConfig
}

func NewEmailService(config *emailConfig.EmailConfig) EmailService {
	return &EmailServiceImpl{config: config}
}

func (e *EmailServiceImpl) SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", e.config.EmailFrom)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer(e.config.SMTPHost, e.config.SMTPPort, e.config.SMTPUser, e.config.SMTPPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
