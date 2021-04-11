package helpers

import (
	"net/smtp"

	"github.com/northern-dehacks/northern-dehacks-server/constants"
	"github.com/northern-dehacks/northern-dehacks-server/keys"
)

var auth smtp.Auth

func AuthEmail() {
	from := constants.EMAIL
	pass := keys.EMAIL_PASS

	auth = smtp.PlainAuth("", from, pass, "smtp.gmail.com")
}
