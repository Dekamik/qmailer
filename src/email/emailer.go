package email

import (
	"net/smtp"
	"strings"
)

type smtpPlainAuth func(string, string, string, string) smtp.Auth
type smtpSendMail func(string, smtp.Auth, string, []string, []byte) error

type Emailer struct {
	plainAuth smtpPlainAuth
	sendMail  smtpSendMail
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

func NewEmailer(plainAuth smtpPlainAuth, sendMail smtpSendMail) *Emailer {
	return &Emailer{plainAuth: plainAuth, sendMail: sendMail}
}

func (e *Emailer) Send(email Email, config Config) error {
	message := []byte("From: " + config.From + "\r\n" +
		"To: " + strings.Join(email.To, ", ") + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"Email Body: " + email.Body + "\r\n")

	auth := e.plainAuth("", config.User, config.Pass, config.Host)
	return e.sendMail(config.Host+":"+config.Port, auth, config.From, email.To, message)
}
