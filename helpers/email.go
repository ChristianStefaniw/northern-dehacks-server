package helpers

import (
	"github.com/northern-dehacks/northern-dehacks-server/constants"
	"net/smtp"
)

func SendEmail(msg, to string) error {
	return smtp.SendMail("smtp.gmail.com:587", auth, constants.EMAIL, []string{to}, []byte(msg))
}
