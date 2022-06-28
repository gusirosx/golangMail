package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// The smtp package implements the Simple Mail Transfer Protocol.
/*
   Note: Gmail is not ideal for testing applications. You should
   use an online service such as Mailtrap as an SMTP server
*/

var host = os.Getenv("SMTP_HOST")
var port = os.Getenv("SMTP_PORT")
var user = os.Getenv("SMTP_USER")
var pass = os.Getenv("SMTP_PASS")

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

func main() {

	from := "gsrodrigues280@example.com"

	to := []string{
		"gusirosx@example.com",
	}

	addr := host + port

	msg := []byte("From: gsrodrigues280@gmail.com\r\n" +
		"To: gusirosx@gmail.com\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")

	auth := smtp.PlainAuth("", user, pass, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}

//https://zetcode.com/golang/email-smtp/
