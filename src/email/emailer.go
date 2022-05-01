package email

import (
    "strings"
)

type IEmailer interface {
    Send(email Email, config Config) error
}

type Emailer struct {
    smtpWrapper ISmtpWrapper
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

func NewEmailer(smtpWrapper ISmtpWrapper) *Emailer {
    return &Emailer{smtpWrapper: smtpWrapper}
}

func (e *Emailer) Send(email Email, config Config) error {
    message := []byte("From: " + config.From + "\r\n" +
        "To: " + strings.Join(email.To, ", ") + "\r\n" +
        "Subject: " + email.Subject + "\r\n" +
        "Email Body: " + email.Body + "\r\n")

    auth := e.smtpWrapper.PlainAuth("", config.User, config.Pass, config.Host)
    return e.smtpWrapper.SendMail(config.Host+":"+config.Port, auth, config.From, email.To, message)
}
