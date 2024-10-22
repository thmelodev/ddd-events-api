package repositories

import (
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/mappers"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/models"
	"github.com/thmelodev/ddd-events-api/src/providers/db"
)

type IEventRepository interface {
	Create(event *domain.EventAggregate) error
	// FindById(id string) (*domain.EventAggregate, error)
	FindAll() ([]*domain.EventAggregate, error)
	// Update(event *domain.EventAggregate) error
	// Delete(id string) error
}

type EventRepository struct {
	db          *db.GormDatabase
	eventMapper *mappers.EventMapper
}

func NewEventRepository(db *db.GormDatabase, eventMapper *mappers.EventMapper) *EventRepository {
	return &EventRepository{db: db, eventMapper: eventMapper}
}

func (r *EventRepository) Create(event *domain.EventAggregate) error {
	model := r.eventMapper.ToModel(event)
	if err := r.db.DB.Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) FindAll() ([]*domain.EventAggregate, error) {
	var models []*models.Event
	if err := r.db.DB.Find(&models).Error; err != nil {
		return nil, err
	}

	var events []*domain.EventAggregate
	for _, model := range models {
		event, err := r.eventMapper.ToDomain(model)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
