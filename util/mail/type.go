package util_mail

import "gopkg.in/gomail.v2"

type MailManager interface {
	SentMessage(msg *gomail.Message) error
}
