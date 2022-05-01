package email

import (
	"net/smtp"
	"testing"
)

var config = Config{
	Host: "smtp.any.com",
	Port: "25",
	User: "AnyUser",
	Pass: "AnyPassword",
	From: "no-reply@any.com",
}

var email = Email{
	To:      []string{"any1@any.com", "any2@any.com"},
	Subject: "AnySubject",
	Body:    "AnyBody",
}

func mockPlainAuth(identity string, username string, password string, host string) smtp.Auth {
	return nil
}

func mockSendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return nil
}

func TestSendEmail(t *testing.T) {
	e := NewEmailer(mockPlainAuth, mockSendMail)
	err := e.Send(email, config)
	if err != nil {
		t.FailNow()
	}
}
