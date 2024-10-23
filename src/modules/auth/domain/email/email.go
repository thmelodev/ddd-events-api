package email

import (
	"regexp"

	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

type EmailValueObject struct {
	value string
}

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

func New(value string) (*EmailValueObject, error) {
	if !isValidEmail(value) {
		return nil, apiErrors.NewInvalidPropsError("invalid email format")
	}
	return &EmailValueObject{value: value}, nil
}

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func (e *EmailValueObject) Value() string {
	return e.value
}
