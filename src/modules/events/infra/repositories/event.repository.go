package repositories

import (
	"fmt"

	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/mappers"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/models"
	"github.com/thmelodev/ddd-events-api/src/providers/db"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
	"gorm.io/gorm"
)

var _ IEventRepository = (*EventRepository)(nil)

type IEventRepository interface {
	Save(event *domain.EventAggregate) error
	FindById(id string) (*domain.EventAggregate, error)
	FindAll() ([]*domain.EventAggregate, error)
	Delete(event *domain.EventAggregate) error
}

type EventRepository struct {
	db          *db.GormDatabase
	eventMapper *mappers.EventMapper
}

func NewEventRepository(db *db.GormDatabase, eventMapper *mappers.EventMapper) *EventRepository {
	return &EventRepository{db: db, eventMapper: eventMapper}
}

func (r *EventRepository) Save(event *domain.EventAggregate) error {
	model := r.eventMapper.ToModel(event)

	if err := r.db.DB.Save(model).Error; err != nil {
		return apiErrors.NewRepositoryError(fmt.Errorf("failed to save event: %w", err).Error())
	}

	return nil
}

func (r *EventRepository) FindAll() ([]*domain.EventAggregate, error) {
	var models []*models.EventModel
	if err := r.db.DB.Find(&models).Error; err != nil {
		return nil, apiErrors.NewRepositoryError(fmt.Errorf("failed to find all events: %w", err).Error())
	}

	var events []*domain.EventAggregate
	for _, model := range models {
		e, err := r.eventMapper.ToDomain(model)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func (r *EventRepository) FindById(id string) (*domain.EventAggregate, error) {
	var model models.EventModel

	if err := r.db.DB.Where("id = ?", id).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apiErrors.NewNoDataFoundRepositoryError(fmt.Errorf("event with id %s not found", id).Error())
		}
		return nil, apiErrors.NewRepositoryError(fmt.Errorf("failed to find event by id %s: %w", id, err).Error())
	}

	eventAggregate, err := r.eventMapper.ToDomain(&model)

	if err != nil {
		return nil, apiErrors.NewRepositoryError(fmt.Errorf("failed to map event model to domain aggregate: %w", err).Error())
	}

	return eventAggregate, nil
}

func (r *EventRepository) Delete(event *domain.EventAggregate) error {
	model := r.eventMapper.ToModel(event)

	if err := r.db.DB.Delete(model).Error; err != nil {
		return apiErrors.NewInvalidPropsError(
			fmt.Errorf("failed to delete event with ID %s: %w", event.GetId(), err).Error(),
		)
	}

	return nil
}
