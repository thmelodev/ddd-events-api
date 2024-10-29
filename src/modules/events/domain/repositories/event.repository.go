package repositories

import (
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
)

type IEventRepository interface {
	Save(event *domain.EventAggregate) error
	FindById(id string) (*domain.EventAggregate, error)
	FindByUserId(id string) ([]*domain.EventAggregate, error)
	FindAll() ([]*domain.EventAggregate, error)
	Delete(event *domain.EventAggregate) error
}
