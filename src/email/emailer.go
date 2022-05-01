package email

import (
	"strings"
)

type Emailer interface {
	Send(email Email) error
}

type emailer struct {
	smtpWrapper SmtpWrapper
	config      Config
}

type Config struct {
	Host string
	Port string
	User string
	Pass string
	From string
}

type Email struct {
	To      []string
	Subject string
	Body    string
}

func NewEmailer(smtpWrapper SmtpWrapper, config Config) *emailer {
	return &emailer{smtpWrapper: smtpWrapper, config: config}
}

func (e *emailer) Send(email Email) error {
	message := []byte("From: " + e.config.From + "\r\n" +
		"To: " + strings.Join(email.To, ", ") + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"Email Body: " + email.Body + "\r\n")

	auth := e.smtpWrapper.PlainAuth("", e.config.User, e.config.Pass, e.config.Host)
	return e.smtpWrapper.SendMail(e.config.Host+":"+e.config.Port, auth, e.config.From, email.To, message)
}
