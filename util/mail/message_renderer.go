package util_mail

import (
	"io"

	gomail "gopkg.in/gomail.v2"
)

func NewMessage(from, to, subject, body string) *gomail.Message {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Cc", from)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	return mailer
}

func AttachFileToMessage(message *gomail.Message, filename string, fileBytes []byte) {
	message.Attach(filename, gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(fileBytes)
		return err
	}))
}
