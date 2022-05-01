package email

import (
    "github.com/stretchr/testify/mock"
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

var testEmail = Email{
    To:      []string{"any1@any.com", "any2@any.com"},
    Subject: "AnySubject",
    Body:    "AnyBody",
}

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
    return "", nil, nil
}

func (m *MockedAuth) Next(fromServer []byte, more bool) (toServer []byte, err error) {
    return nil, nil
}

func TestSend_Any_CallsPlainAuthAndSendMail(t *testing.T) {
    w := &MockedSmtpWrapper{}
    a := &MockedAuth{}
    w.Mock.On(
        "PlainAuth",
        "",
        config.User,
        config.Pass,
        config.Host,
    ).Return(a)
    w.Mock.On(
        "SendMail",
        config.Host+":"+config.Port,
        a,
        config.From,
        testEmail.To,
        mock.Anything,
    ).Return(nil)
    e := NewEmailer(w)

    _ = e.Send(testEmail, config)

    w.AssertCalled(t, "PlainAuth", "", config.User, config.Pass, config.Host)
}
