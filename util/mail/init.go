package util_mail

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	gomail "gopkg.in/gomail.v2"
)

type mailManager struct {
	dialer      *gomail.Dialer
	frontEndUrl string
}

func NewMailManager(host string, portNumber int, email string, password string, frontEndUrl string) MailManager {
	dialer := gomail.NewDialer(host, portNumber, email, password)

	return &mailManager{
		dialer: dialer,
	}
}

func (m *mailManager) SentMessage(msg *gomail.Message) error {
	err := m.dialer.DialAndSend(msg)

	return err
}

//go:embed templates/*.html
var templates embed.FS

var parsedTemplates = template.Must(template.ParseFS(templates, "templates/*.html"))

func (m *mailManager) SentVerifyEmail(token string, email string) error {
	var buf bytes.Buffer
	err := parsedTemplates.ExecuteTemplate(&buf, "email_verification.html", map[string]string{
		"Link": fmt.Sprintf("%s/verify-email?token=%s", m.frontEndUrl, token),
	})
	if err != nil {
		return err
	}

	msg := newMessage(m.dialer.Username, email, "Email Verification", buf.String())

	return m.dialer.DialAndSend(msg)
}

func (m *mailManager) SentResetPassword(token, email string) error {
	var buf bytes.Buffer
	err := parsedTemplates.ExecuteTemplate(&buf, "reset_password.html", map[string]string{
		"Link": fmt.Sprintf("%s/reset-password?token=%s", m.frontEndUrl, token),
	})
	if err != nil {
		return err
	}

	msg := newMessage(m.dialer.Username, email, "Reset Password Link", buf.String())

	return m.dialer.DialAndSend(msg)
}

func newMessage(from, to, subject, body string) *gomail.Message {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Cc", from)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	return mailer
}

// func attachFileToMessage(msg *gomail.Message, attachment Attachment) {
// 	msg.Attach(attachment.FileName, gomail.SetCopyFunc(func(w io.Writer) error {
// 		_, err := w.Write(attachment.FileByte)
// 		return err
// 	}))
// }
