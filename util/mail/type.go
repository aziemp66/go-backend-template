package util_mail

type MailManager interface {
	SentVerifyEmail(token, email string) error
	SentResetPassword(token, email string) error
}

type Attachment struct {
	FileName string
	FileByte []byte
}
