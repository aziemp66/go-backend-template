package util_mail

import (
	gomail "gopkg.in/gomail.v2"
)

type mailManager struct {
	dialer *gomail.Dialer
}

func NewMailManager(host string, portNumber int, email string, password string) MailManager {
	dialer := gomail.NewDialer(host, portNumber, email, password)

	return &mailManager{
		dialer: dialer,
	}
}

func (m *mailManager) SentMessage(msg *gomail.Message) error {
	err := m.dialer.DialAndSend(msg)

	return err
}
