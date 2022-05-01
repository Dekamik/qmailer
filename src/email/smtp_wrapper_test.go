package email

import (
	"github.com/stretchr/testify/assert"
	"net/smtp"
	"testing"
)

func TestNewSmtpWrapper_Any_ReturnsSmtpWrapper(t *testing.T) {
	_ = NewSmtpWrapper()
}

func TestPlainAuth_Any_CallsPlainAuth(t *testing.T) {
	wasCalled := false
	mockedPlainAuth := func(identity string, username string, password string, hostname string) smtp.Auth {
		wasCalled = true
		return nil
	}
	w := Wrapper{smtpPlainAuth: mockedPlainAuth}

	w.PlainAuth("", "", "", "")

	assert.True(t, wasCalled)
}

func TestSendMail_Any_CallsSendMail(t *testing.T) {
	wasCalled := false
	mockedSendMail := func(string, smtp.Auth, string, []string, []byte) error {
		wasCalled = true
		return nil
	}
	w := Wrapper{smtpSendMail: mockedSendMail}

	_ = w.SendMail("", nil, "", nil, nil)

	assert.True(t, wasCalled)
}
