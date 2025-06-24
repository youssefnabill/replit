
package user

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrInvalidName  = errors.New("invalid name")
)

// Email represents an email value object
type Email struct {
	value string
}

// NewEmail creates a new email with validation
func NewEmail(email string) (*Email, error) {
	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}
	return &Email{value: email}, nil
}

// String returns the email value
func (e *Email) String() string {
	return e.value
}

// isValidEmail validates email format
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}
