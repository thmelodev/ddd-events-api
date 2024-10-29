package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

func TestNewEmailSucess(t *testing.T) {
	email, err := NewEmail("teste@gmail.com")

	assert.Nil(t, err)
	assert.Equal(t, email.Value(), "teste@gmail.com")
}

func TestNewEmailError(t *testing.T) {
	email, err := NewEmail("teste")

	assert.Error(t, err)
	assert.Nil(t, email)
	assert.Equal(t, err, apiErrors.NewInvalidPropsError("invalid email format"))
}
