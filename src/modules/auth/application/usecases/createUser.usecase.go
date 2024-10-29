package usecases

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/domain"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/domain/repositories"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
	"github.com/thmelodev/ddd-events-api/src/utils/interfaces"
)

var _ interfaces.IUsecase = (*CreateUserUsecase)(nil)

type CreateUserUsecase struct {
	userRepository repositories.IUserRepository
	config         *config.Config
}

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserId   string `json:"-"`
}

func NewCreateUserUsecase(
	userRepository repositories.IUserRepository,
	config *config.Config,
) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepository: userRepository,
		config:         config,
	}
}

func (u CreateUserUsecase) Execute(context context.Context, dto any) (any, error) {
	userDTO, ok := dto.(*CreateUserDTO)

	if !ok {
		return nil, apiErrors.NewInvalidPropsError(fmt.Errorf("invalid props: %v, invalid type: %t", dto, dto).Error())
	}

	if userDTO.UserId != u.config.App.AdminId {
		return nil, apiErrors.NewUnauthorizedError("only admin can create users")
	}

	userExist, _ := u.userRepository.FindByEmail(userDTO.Email)

	if userExist != nil {
		return nil, apiErrors.NewInvalidPropsError(fmt.Errorf("user with email %s already exists", userDTO.Email).Error())
	}

	userAggregate, err := domain.NewUser(domain.UserProps{
		Email:    userDTO.Email,
		Password: userDTO.Password,
	})

	if err != nil {
		return nil, err
	}

	err = u.userRepository.Save(userAggregate)

	if err != nil {
		return nil, err
	}

	return gin.H{"id": userAggregate.GetId()}, nil
}
