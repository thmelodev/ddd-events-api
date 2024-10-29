package domain

import (
	"github.com/google/uuid"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

type UserAggregate struct {
	id       string
	email    *EmailValueObject
	password string
}

type UserProps struct {
	Email    string
	Password string
}

func NewUser(props UserProps) (*UserAggregate, error) {
	user := &UserAggregate{id: uuid.New().String()}

	err := user.build(UserProps{
		Email:    props.Email,
		Password: props.Password,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func LoadUser(props UserProps, id string) (*UserAggregate, error) {
	user := &UserAggregate{}

	err := user.setId(id)

	if err != nil {
		return nil, err
	}

	err = user.build(props)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (e *UserAggregate) setId(id string) error {
	_, err := uuid.Parse(id)

	if err != nil {
		return apiErrors.NewInvalidPropsError("id is invalid")
	}

	e.id = id

	return nil
}

func (u *UserAggregate) build(props UserProps) error {
	err := u.setEmail(props.Email)

	if err != nil {
		return err
	}

	err = u.setPassword(props.Password)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserAggregate) setEmail(e string) error {

	emailValueObject, err := NewEmail(e)

	if err != nil {
		return err
	}

	u.email = emailValueObject

	return nil
}

func (u *UserAggregate) setPassword(password string) error {
	if len(password) < 6 {
		return apiErrors.NewInvalidPropsError("password must be at least 6 characters")
	}

	u.password = password
	return nil
}

func (u *UserAggregate) GetId() string {
	return u.id
}

func (u *UserAggregate) GetEmail() string {
	return u.email.Value()
}

func (u *UserAggregate) GetPassword() string {
	return u.password
}
