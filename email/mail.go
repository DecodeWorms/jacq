package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"jacq/config"
	"jacq/model"
)

const (
	jacqEmail                    = "lekzy.csharp@gmail.com"
	emailVerifications           = "Email Verification"
	passswordChangedSuccessfully = "User password changed successfully"
	pinChangedSuccessfully       = "User Transaction pin changed successfully"
)

// SendEmail sends an email with the given subject and body to the specified recipient
func SendEmailVerification(data model.Email) error {

	c := config.ImportConfig(config.OSSource{})
	fmt.Println(c.GmailPassword)
	m := gomail.NewMessage()
	m.SetHeader("From", jacqEmail)
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", emailVerifications)
	m.SetBody("text/plain", data.Body)

	d := gomail.NewDialer(c.GomailName, c.GomailPort, jacqEmail, c.GmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// SendEmail sends an email with the given subject and body to the specified recipient
func SendPasswordChangedSuccessfully(data model.Email) error {

	c := config.ImportConfig(config.OSSource{})
	fmt.Println(c.GmailPassword)
	m := gomail.NewMessage()
	m.SetHeader("From", jacqEmail)
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", pinChangedSuccessfully)
	m.SetBody("text/plain", data.Body)

	d := gomail.NewDialer(c.GomailName, c.GomailPort, jacqEmail, c.GmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// SendEmail sends an email with the given subject and body to the specified recipient
func SendPinChangedSuccessfully(data model.Email) error {

	c := config.ImportConfig(config.OSSource{})
	fmt.Println(c.GmailPassword)
	m := gomail.NewMessage()
	m.SetHeader("From", jacqEmail)
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", passswordChangedSuccessfully)
	m.SetBody("text/plain", data.Body)

	d := gomail.NewDialer(c.GomailName, c.GomailPort, jacqEmail, c.GmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
