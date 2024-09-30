package util_mail

type (
	EmailVerification struct {
		Token string
	}

	PasswordReset struct {
		Token string
	}
)
