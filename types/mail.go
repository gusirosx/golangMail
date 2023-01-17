package types

import (
	"regexp"
	"strings"
)

// Mail struct to hold the submitted form.
type Mail struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Errors    map[string]string
}

func (mail *Mail) Validate() bool {
	// Map will not be empty if errors are present.
	mail.Errors = make(map[string]string)
	// If the first name is empty, add to errors map.
	if strings.TrimSpace(mail.FirstName) == "" {
		mail.Errors["FirstName"] = "First name cannot be empty."
	}
	if strings.TrimSpace(mail.LastName) == "" {
		mail.Errors["LastName"] = "Last name cannot be empty."
	}
	// If email is not empty, check if it's the correct format.
	if strings.TrimSpace(mail.Email) != "" {
		re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		matched := re.Match([]byte(mail.Email))
		if !matched {
			mail.Errors["EmailFormat"] = "Email is not valid."
		}
	} else {
		mail.Errors["Email"] = "Email cannot be empty."
	}

	if strings.TrimSpace(mail.Subject) == "" {
		mail.Errors["Subject"] = "Subject cannot be empty."
	}
	if strings.TrimSpace(mail.Body) == "" {
		mail.Errors["Body"] = "Body cannot be empty."
	}
	// Return empty map (assuming it's empty/no errors)
	return len(mail.Errors) == 0
}
