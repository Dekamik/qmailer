package email

import (
	"github.com/stretchr/testify/mock"
	"net/smtp"
	"strings"
	"testing"
)

type MockedSmtpWrapper struct {
	mock.Mock
}

func (m *MockedSmtpWrapper) PlainAuth(identity string, username string, password string, hostname string) smtp.Auth {
	args := m.Called(identity, username, password, hostname)
	return args.Get(0).(smtp.Auth)
}

func (m *MockedSmtpWrapper) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	args := m.Called(addr, a, from, to, msg)
	return args.Error(0)
}

type MockedAuth struct {
	mock.Mock
}

func (m *MockedAuth) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	args := m.Called(server)
	return args.String(0), args.Get(1).([]byte), args.Error(2)
}

func (m *MockedAuth) Next(fromServer []byte, more bool) (toServer []byte, err error) {
	args := m.Called(fromServer, more)
	return args.Get(0).([]byte), args.Error(1)
}

func TestSend_Any_CallsPlainAuthAndSendMail(t *testing.T) {
	// Arrange
	w := &MockedSmtpWrapper{}
	a := &MockedAuth{}

	config := Config{
		Host: "smtp.any.com",
		Port: "25",
		User: "AnyUser",
		Pass: "AnyPassword",
		From: "no-reply@any.com",
	}
	testEmail := Email{
		To:      []string{"any1@any.com", "any2@any.com"},
		Subject: "AnySubject",
		Body:    "AnyBody",
	}

	addr := config.Host + ":" + config.Port
	message := []byte("From: " + config.From + "\r\n" +
		"To: " + strings.Join(testEmail.To, ", ") + "\r\n" +
		"Subject: " + testEmail.Subject + "\r\n" +
		"Email Body: " + testEmail.Body + "\r\n")
	w.Mock.On(
		"PlainAuth",
		"",
		config.User,
		config.Pass,
		config.Host,
	).Return(a)
	w.Mock.On(
		"SendMail",
		addr,
		a,
		config.From,
		testEmail.To,
		message,
	).Return(nil)
	e := NewEmailer(w, config)

	// Act
	_ = e.Send(testEmail)

	// Assert
	w.AssertCalled(t, "PlainAuth", "", config.User, config.Pass, config.Host)
	w.AssertCalled(t, "SendMail", addr, a, config.From, testEmail.To, message)
}
