package repositories

import "github.com/thmelodev/ddd-events-api/src/modules/auth/domain"

type IUserRepository interface {
	Save(u *domain.UserAggregate) error
	FindByEmail(email string) (*domain.UserAggregate, error)
}
