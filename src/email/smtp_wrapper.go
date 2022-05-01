package email

import "net/smtp"

type ISmtpWrapper interface {
    PlainAuth(string, string, string, string) smtp.Auth
    SendMail(string, smtp.Auth, string, []string, []byte) error
}

type SmtpWrapper struct{}

func (w *SmtpWrapper) PlainAuth(identity string, username string, password string, hostname string) smtp.Auth {
    return smtp.PlainAuth(identity, username, password, hostname)
}

func (w *SmtpWrapper) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
    return smtp.SendMail(addr, a, from, to, msg)
}
