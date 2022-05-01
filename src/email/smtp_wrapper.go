package email

import "net/smtp"

type SmtpWrapper interface {
	PlainAuth(string, string, string, string) smtp.Auth
	SendMail(string, smtp.Auth, string, []string, []byte) error
}

type Wrapper struct {
	smtpPlainAuth func(string, string, string, string) smtp.Auth
	smtpSendMail  func(string, smtp.Auth, string, []string, []byte) error
}

func NewSmtpWrapper() *Wrapper {
	return &Wrapper{
		smtpPlainAuth: smtp.PlainAuth,
		smtpSendMail:  smtp.SendMail,
	}
}

func (w *Wrapper) PlainAuth(identity string, username string, password string, hostname string) smtp.Auth {
	return w.smtpPlainAuth(identity, username, password, hostname)
}

func (w *Wrapper) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return w.smtpSendMail(addr, a, from, to, msg)
}
