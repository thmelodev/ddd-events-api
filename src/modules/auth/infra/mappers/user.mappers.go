package mappers

import (
	"github.com/thmelodev/ddd-events-api/src/modules/auth/domain"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/infra/models"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) ToModel(user *domain.UserAggregate) *models.UserModel {
	return &models.UserModel{
		Id:       user.GetId(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}
}

func (m *UserMapper) ToDomain(u *models.UserModel) (*domain.UserAggregate, error) {
	domain, err := domain.LoadUser(domain.UserProps{
		Email:    u.Email,
		Password: u.Password,
	}, u.Id)

	if err != nil {
		return nil, err
	}

	return domain, nil
}
