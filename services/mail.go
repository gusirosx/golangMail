package services

import (
	"fmt"
	"golangMail/types"
	"log"
	"net/smtp"
	"os"
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

	// Message.
	message := []byte(
		"Subject: " + mail.Subject + "\r\n" +
			"From:" + from + "\r\n" +
			"To:" + createStringOfSlice(to) + "\r\n" +
			"Subject:" + mail.Subject + "\r\n" +
			"\r\n" +
			"Name: " + mail.FirstName + " " + mail.LastName + "\n" +
			"Email: " + mail.Email + "\n" +
			"Mobile: " + mail.Phone + "\n" +
			"Subject: " + mail.Subject + "\n" +
			"----------------------------------------------------------------" + "\n" +
			mail.Body + "\n" + "\r\n",
	)

	// Authentication information for the sender email address.
	auth := smtp.PlainAuth("", user, password, host)

	// Sending email.
	err := smtp.SendMail(host+port, auth, from, to, message)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("Email successfully sent to: ", to)
	//fmt.Println("Email Sent Successfully!")
	return nil
}

func createStringOfSlice(emails []string) (toList string) {
	for _, email := range emails {
		str := fmt.Sprint(email + ", ")
		toList = toList + str
	}
	return toList[0 : len(toList)-2]
}
