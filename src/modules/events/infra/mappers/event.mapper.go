package mappers

import (
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/models"
)

type IEventMapper interface {
	ToModel(*domain.EventAggregate) *models.Event
	ToDomain(*models.Event) *domain.EventAggregate
}

type EventMapper struct{}

func NewEventMapper() *EventMapper {
	return &EventMapper{}
}

func (m *EventMapper) ToModel(event *domain.EventAggregate) *models.Event {
	return &models.Event{
		Id:          event.GetId(),
		Name:        event.GetName(),
		Description: event.GetDescription(),
		Location:    event.GetLocation(),
		DateTime:    event.GetDateTime(),
		UserID:      event.GetUserId(),
	}
}

func (m *EventMapper) ToDomain(event *models.Event) (*domain.EventAggregate, error) {
	domain, err := domain.LoadEvent(domain.EventProps{
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    event.DateTime,
		UserId:      event.UserID,
	}, event.Id)

	if err != nil {
		return nil, err
	}

	return domain, nil
}
