package email

import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "net/smtp"
    "testing"
)

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

func TestNewSmtpWrapper_Any_ReturnsSmtpWrapper(t *testing.T) {
    _ = NewSmtpWrapper()
}

func TestPlainAuth_Any_CallsPlainAuth(t *testing.T) {
    isCalled := false
    mockedPlainAuth := func(identity string, username string, password string, hostname string) smtp.Auth {
        isCalled = true
        return nil
    }
    w := SmtpWrapper{smtpPlainAuth: mockedPlainAuth}

    w.PlainAuth("", "", "", "")

    assert.True(t, isCalled)
}

func TestSendMail_Any_CallsSendMail(t *testing.T) {
    isCalled := false
    mockedSendMail := func(string, smtp.Auth, string, []string, []byte) error {
        isCalled = true
        return nil
    }
    w := SmtpWrapper{smtpSendMail: mockedSendMail}

    _ = w.SendMail("", nil, "", nil, nil)

    assert.True(t, isCalled)
}
