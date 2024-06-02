package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"jacq/config"
	"jacq/model"
)

const (
	jacqEmail = "lekzy.csharp@gmail.com"
	subject   = "Email Verification"
)

// SendEmail sends an email with the given subject and body to the specified recipient
func SendEmail(data model.Email) error {

	c := config.ImportConfig(config.OSSource{})
	fmt.Println(c.GmailPassword)
	m := gomail.NewMessage()
	m.SetHeader("From", jacqEmail)
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", data.Body)

	d := gomail.NewDialer(c.GomailName, c.GomailPort, jacqEmail, c.GmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
