package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/domain/repositories"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
	"github.com/thmelodev/ddd-events-api/src/utils/hash"
	"github.com/thmelodev/ddd-events-api/src/utils/interfaces"
	"github.com/thmelodev/ddd-events-api/src/utils/jwt"
)

var _ interfaces.IUsecase = (*LoginUserUsecase)(nil)

type LoginUserUsecase struct {
	userRepository repositories.IUserRepository
	config         *config.Config
}

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginUserUsecase(
	userRepository repositories.IUserRepository,
	config *config.Config,
) *LoginUserUsecase {
	return &LoginUserUsecase{
		userRepository: userRepository,
		config:         config,
	}
}

func (u LoginUserUsecase) Execute(context context.Context, dto any) (any, error) {
	userDTO, ok := dto.(*LoginUserDTO)

	if !ok {
		return nil, apiErrors.NewInvalidPropsError(fmt.Errorf("invalid props: %v, invalid type: %t", dto, dto).Error())
	}

	userExist, _ := u.userRepository.FindByEmail(userDTO.Email)

	if userExist == nil {
		return nil, apiErrors.NewInvalidPropsError("email or password is invalid")
	}

	if !hash.IsValidPassword(userExist.GetPassword(), userDTO.Password) {
		return nil, apiErrors.NewInvalidPropsError("email or password is invalid")
	}

	token, err := jwt.GenerateToken(userExist.GetEmail(), userExist.GetId(), u.config)

	if err != nil {
		return nil, errors.New("could not authenticate user")
	}

	return gin.H{"token": token}, nil
}
