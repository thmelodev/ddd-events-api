package mappers

import (
	"github.com/thmelodev/ddd-events-api/src/modules/auth/domain/user"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/infra/models"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) ToModel(user *user.UserAggregate) *models.UserModel {
	return &models.UserModel{
		Id:       user.GetId(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}
}

func (m *UserMapper) ToDomain(u *models.UserModel) (*user.UserAggregate, error) {
	domain, err := user.Load(user.UserProps{
		Email:    u.Email,
		Password: u.Password,
	}, u.Id)

	if err != nil {
		return nil, err
	}

	return domain, nil
}
