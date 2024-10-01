package util_mail

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
)

//go:embed templates/*.html
var templates embed.FS

var parsedTemplates = template.Must(template.ParseFS(templates, "templates/*.html"))

func RenderEmailVerificationTemplate(token string, frontEndUrl string) (string, error) {
	var buf bytes.Buffer
	err := parsedTemplates.ExecuteTemplate(&buf, "email_verification.html", map[string]string{
		"Link": fmt.Sprintf("%s/verify-email?token=%s", frontEndUrl, token),
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderPasswordResetTemplate(token string, frontEndUrl string) (string, error) {
	var buf bytes.Buffer
	err := parsedTemplates.ExecuteTemplate(&buf, "reset_password.html", map[string]string{
		"Link": fmt.Sprintf("%s/reset-password?token=%s", frontEndUrl, token),
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
