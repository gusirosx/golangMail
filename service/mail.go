package service

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// Sender email address info. Email address, password and smtp server.
const (
	SENDER_EMAIL     string = "EmailAddress@email.com"
	SENDER_EMAIL_PWD string = "PasswordHere"
	SENDER_SMTP      string = "smtp.gmail.com"
)

// formData struct to hold the submitted form.
type formData struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Subject   string
	Body      string
}

// The smtp package implements the Simple Mail Transfer Protocol.
/*
   Note: Gmail is not ideal for testing applications. You should
   use an online service such as Mailtrap as an SMTP server
*/

var host = os.Getenv("SMTP_HOST")
var port = os.Getenv("SMTP_PORT")
var user = os.Getenv("SMTP_USER")
var pass = os.Getenv("SMTP_PASS")

// Global variable tpl holds all the parsed templates.
var tpl *template.Template
var err error

// Parse all .gohtml templates in the templates folder.
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func Index(res http.ResponseWriter, req *http.Request) {

	// Form is submitted.
	if req.Method == http.MethodPost {
		// Collect form input and validate
		fd := &formValidate{
			FirstName: req.FormValue("frmFirstName"),
			LastName:  req.FormValue("frmLastName"),
			Email:     req.FormValue("frmEmail"),
			Phone:     req.FormValue("frmPhone"),
			Subject:   req.FormValue("frmSubject"),
			Body:      req.FormValue("frmBody"),
			SentFlag:  false,
		}
		// Check for errors.
		if fd.validate() == false {
			err := tpl.ExecuteTemplate(res, "index.gohtml", fd)
			if err != nil {
				log.Fatalln(err)
			}
			return
		}
		// Send email
		sendEmail(*fd)

		/*
			SentFlag variable will be set to true once the email sendEmail function runs.
			It is then used in the template to check if the message sent html should be displayed.
		*/
		fd.SentFlag = true
		err := tpl.ExecuteTemplate(res, "index.gohtml", fd)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// sendEmail function takes a formData object and uses the info to send an email.
func sendEmail(fd formValidate) {
	// Authentication information for the sender email address. Configurable above.
	//auth := smtp.PlainAuth("", SENDER_EMAIL, SENDER_EMAIL_PWD, SENDER_SMTP)
	auth := smtp.PlainAuth("", user, pass, host)

	from := "gsrodrigues280@example.com"

	addr := host + port

	// msg := []byte("From: gsrodrigues280@gmail.com\r\n" +
	// 	"To: gusirosx@gmail.com\r\n" +
	// 	"Subject: Test mail\r\n\r\n" +
	// 	"Email body\r\n")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"gusirosx@example.com"}
	msg := []byte("To: Marck527@gmail.com\r\n" +
		"Subject:" + fd.Subject + "\r\n" +
		"\r\n" +
		"Name: " + fd.FirstName + " " + fd.LastName + "\n" +
		"Email: " + fd.Email + "\n" +
		"Mobile: " + fd.Phone + "\n" +
		"Subject: " + fd.Subject + "\n" +
		"-----------------------------------------------------------------------------" + "\n" +
		fd.Body + "\n" + "\r\n")
	//err := smtp.SendMail("smtp.gmail.com:587", auth, SENDER_EMAIL, to, msg)
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email successfully sent to: ", to)
}

// ================================================================

type formValidate struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Subject   string
	Body      string
	Errors    map[string]string
	SentFlag  bool
}

func (fv *formValidate) validate() bool {
	// Map will not be empty if errors are present.
	fv.Errors = make(map[string]string)
	// If the first name is empty, add to errors map.
	if strings.TrimSpace(fv.FirstName) == "" {
		fv.Errors["FirstName"] = "First name cannot be empty."
	}
	if strings.TrimSpace(fv.LastName) == "" {
		fv.Errors["LastName"] = "Last name cannot be empty."
	}
	// If email is not empty, check if it's the correct format.
	if strings.TrimSpace(fv.Email) != "" {
		re := regexp.MustCompile(".+@.+\\..+")
		matched := re.Match([]byte(fv.Email))

		if matched == false {
			fv.Errors["EmailFormat"] = "Email is not valid."
		}
	} else {
		fv.Errors["Email"] = "Email cannot be empty."
	}

	if strings.TrimSpace(fv.Subject) == "" {
		fv.Errors["Subject"] = "Subject cannot be empty."
	}
	if strings.TrimSpace(fv.Body) == "" {
		fv.Errors["Body"] = "Body cannot be empty."
	}
	// Return empty map (assuming it's empty/no errors)
	return len(fv.Errors) == 0
}
