package controllers

import (
	"fmt"
	"net/http"

	"github.com/northern-dehacks/northern-dehacks-server/constants"
	"github.com/northern-dehacks/northern-dehacks-server/helpers"
)

type contact struct {
	Name    string
	Subject string
	Email   string
	Message string
}

func Contact(w http.ResponseWriter, r *http.Request) {
	var c contact

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Name = r.FormValue("name")
	c.Subject = r.FormValue("subject")
	c.Email = r.FormValue("email")
	c.Message = r.FormValue("message")

	if err := send(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func send(c contact) error {
	body := fmt.Sprintf("Email: %s\nName: %s\nMessage: %s\n", c.Email, c.Name, c.Message)
	msg :=
		"From: " + c.Email + "\n" +
			"To: " + constants.EMAIL + "\n" +
			"Subject: " + c.Subject + "\n\n" +
			body

	return helpers.SendEmail(msg, constants.EMAIL)
}
