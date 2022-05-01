package email

import "net/smtp"

type ISmtpWrapper interface {
    PlainAuth(string, string, string, string) smtp.Auth
    SendMail(string, smtp.Auth, string, []string, []byte) error
}

type SmtpWrapper struct {
    smtpPlainAuth func(string, string, string, string) smtp.Auth
    smtpSendMail  func(string, smtp.Auth, string, []string, []byte) error
}

func NewSmtpWrapper() *SmtpWrapper {
    return &SmtpWrapper{
        smtpPlainAuth: smtp.PlainAuth,
        smtpSendMail:  smtp.SendMail,
    }
}

func (w *SmtpWrapper) PlainAuth(identity string, username string, password string, hostname string) smtp.Auth {
    return w.smtpPlainAuth(identity, username, password, hostname)
}

func (w *SmtpWrapper) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
    return w.smtpSendMail(addr, a, from, to, msg)
}
