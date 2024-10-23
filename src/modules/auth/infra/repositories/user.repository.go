package repositories

import (
	"fmt"

	"github.com/thmelodev/ddd-events-api/src/modules/auth/domain/user"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/infra/mappers"
	"github.com/thmelodev/ddd-events-api/src/modules/auth/infra/models"
	"github.com/thmelodev/ddd-events-api/src/providers/db"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
	"gorm.io/gorm"
)

type IUserRepository interface{}

type UserRepository struct {
	db         *db.GormDatabase
	userMapper mappers.UserMapper
}

func NewUserRepository(db *db.GormDatabase, userMapper mappers.UserMapper) *UserRepository {
	return &UserRepository{db: db, userMapper: userMapper}
}

func (r *UserRepository) Save(u user.UserAggregate) error {
	model := r.userMapper.ToModel(&u)

	if err := r.db.DB.Save(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*user.UserAggregate, error) {
	model := &models.UserModel{}
	if err := r.db.DB.Where("email = ?", email).First(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apiErrors.NewNoDataFoundRepositoryError(fmt.Errorf("user with email %s not found", email).Error())
		}
		return nil, apiErrors.NewRepositoryError(fmt.Errorf("failed to user by email %s: %w", email, err).Error())
	}

	user, err := r.userMapper.ToDomain(model)
	if err != nil {
		return nil, err
	}

	return user, nil
}
