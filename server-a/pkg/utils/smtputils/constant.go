package smtputils

import (
	"embed"
)

//go:embed templates/*.html
var EmailHTMLTemplates embed.FS

const (
	ResetPasswordSubject = "[clean-arch-v2] Please reset your password"
	VerificationSubject  = "[clean-arch-v2] Verify your account"
	PharmacistSubject    = "[clean-arch-v2] Pharmacist account"
)

type emailTemplate string

const (
	ResetPasswordTemplate emailTemplate = "templates/forgot-password.html"
	VerificationTemplate  emailTemplate = "templates/verification.html"
	PharmacistTemplate    emailTemplate = "templates/pharmacist.html"
)
