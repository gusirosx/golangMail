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

func (fv *Mail) Validate() bool {
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
		re := regexp.MustCompile(`.+@.+\\..+`)
		matched := re.Match([]byte(fv.Email))
		if !matched {
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
