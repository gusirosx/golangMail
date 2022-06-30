package types

// Mail struct to hold the submitted form.
type Mail struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"emailName"`
	Phone     string `json:"phone"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

// type  struct {
// 	Sender  string
// 	To      []string
// 	Cc      []string
// 	Bcc     []string
// 	Subject string
// 	Body    string
// }
// func main2() {

// 	from := "gsrodrigues280@example.com"

// 	to := []string{
// 		"gusirosx@example.com",
// 	}

// 	addr := host + port

// 	msg := []byte("From: gsrodrigues280@gmail.com\r\n" +
// 		"To: gusirosx@gmail.com\r\n" +
// 		"Subject: Test mail\r\n\r\n" +
// 		"Email body\r\n")

// 	auth := smtp.PlainAuth("", user, pass, host)

// 	err := smtp.SendMail(addr, auth, from, to, msg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Email sent successfully")
// }

//https://zetcode.com/golang/email-smtp/
