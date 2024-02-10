package value_objects

import (
	"errors"
	"regexp"
)

var ErrInvalidEmail = errors.New("invalid email")

type Email struct {
	value string
}

func (email *Email) Value() string {
	return email.value
}

func (email *Email) Equals(other *Email) bool {
	if email == nil && other == nil {
		return false
	}
	if email == other {
		return true
	}
	return email.value == other.value
}

func validate(value string) error {
	if value == "" {
		return ErrInvalidEmail
	}
	pattern := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !pattern.MatchString(value) {
		return ErrInvalidEmail
	}

	return nil
}

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, ErrInvalidEmail
	}

	if err := validate(value); err != nil {
		return nil, err
	}

	return &Email{value: value}, nil
}
