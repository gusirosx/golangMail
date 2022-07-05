package services

import (
	"fmt"
	"golangMail/types"
	"log"
	"net/smtp"
	"os"
	"strings"
)

// The smtp package implements the Simple Mail Transfer Protocol.
/*
   Note: Gmail is not ideal for testing applications. You should
   use an online service such as Mailtrap as an SMTP server
*/
// smtp server configuration.
var host = os.Getenv("SMTP_HOST")
var port = os.Getenv("SMTP_PORT")
var user = os.Getenv("SMTP_USER")
var password = os.Getenv("SMTP_PASS")

// ContactUs function takes a formData object and uses the info to send an email.
func ContactUs(mail types.Mail) error {

	// Sender data.
	from := "test@example.com"

	// Receiver email address.
	to := []string{
		"gusirosx@example.com",
		"gsrodrigues@example.com",
	}

	// Authentication information for the sender email address.
	auth := smtp.PlainAuth("", user, password, host)

	// Sending email.
	err := smtp.SendMail(host+port, auth, from, to, []byte(buildMessage(mail, to)))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("Email successfully sent to: ", to)
	//fmt.Println("Email Sent Successfully!")
	return nil
}

func buildMessage(mail types.Mail, toList []string) string {
	//msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg := fmt.Sprintf("From: %s\r\n", mail.Email)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(toList, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("Name: %s %s\r\n", mail.FirstName, mail.LastName)
	msg += fmt.Sprintf("Email: %s\r\n", mail.Email)
	msg += fmt.Sprintf("Mobile: %s\r\n", mail.Phone)
	msg += fmt.Sprintf("\r\n %s \r\n", mail.Body)
	return msg
}
