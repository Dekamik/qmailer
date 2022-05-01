package email

import "net/smtp"

type SmtpWrapper interface {
	PlainAuth(string, string, string, string) smtp.Auth
	SendMail(string, smtp.Auth, string, []string, []byte) error
}

type wrapper struct {
	smtpPlainAuth func(string, string, string, string) smtp.Auth
	smtpSendMail  func(string, smtp.Auth, string, []string, []byte) error
}

func NewSmtpWrapper() *wrapper {
	return &wrapper{
		smtpPlainAuth: smtp.PlainAuth,
		smtpSendMail:  smtp.SendMail,
	}
}

func (w *wrapper) PlainAuth(identity string, username string, password string, hostname string) smtp.Auth {
	return w.smtpPlainAuth(identity, username, password, hostname)
}

func (w *wrapper) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return w.smtpSendMail(addr, a, from, to, msg)
}
