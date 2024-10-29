package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser(UserProps{
		Email:    "teste@gmail.com",
		Password: "123456",
	})

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, &UserAggregate{}, user)
	assert.Equal(t, user.GetEmail(), "teste@gmail.com")
	assert.Equal(t, user.GetPassword(), "123456")
}

func TestNewUserErrorEmail(t *testing.T) {
	user, err := NewUser(UserProps{
		Email:    "teste",
		Password: "123456",
	})

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err, apiErrors.NewInvalidPropsError("invalid email format"))
}

func TestNewUserErrorPassword(t *testing.T) {
	user, err := NewUser(UserProps{
		Email:    "teste@gmail.com",
		Password: "123",
	})

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err, apiErrors.NewInvalidPropsError("password must be at least 6 characters"))
}

func TestLoadUser(t *testing.T) {
	id := uuid.New().String()

	user, err := LoadUser(UserProps{
		Email:    "teste2@gmail.com",
		Password: "12345678",
	}, id)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, user.GetId(), id)
	assert.Equal(t, user.GetEmail(), "teste2@gmail.com")
	assert.Equal(t, user.GetPassword(), "12345678")
}

func TestLoadUserErrorId(t *testing.T) {
	user, err := LoadUser(UserProps{
		Email:    "teste@gmail.com",
		Password: "123456",
	}, "123")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err, apiErrors.NewInvalidPropsError("id is invalid"))
}

func TestLoadUserErrorEmail(t *testing.T) {
	user, err := LoadUser(UserProps{
		Email:    "a_gmail.com",
		Password: "123456",
	}, uuid.New().String())

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, err, apiErrors.NewInvalidPropsError("invalid email format"))
}
