package mappers

import (
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/models"
)

type EventMapper struct{}

func NewEventMapper() *EventMapper {
	return &EventMapper{}
}

func (m *EventMapper) ToModel(event *domain.EventAggregate) *models.EventModel {
	return &models.EventModel{
		Id:          event.GetId(),
		Name:        event.GetName(),
		Description: event.GetDescription(),
		Location:    event.GetLocation(),
		DateTime:    event.GetDateTime(),
		UserID:      event.GetUserId(),
	}
}

func (m *EventMapper) ToDomain(e *models.EventModel) (*domain.EventAggregate, error) {
	domain, err := domain.LoadEvent(domain.EventProps{
		Name:        e.Name,
		Description: e.Description,
		Location:    e.Location,
		DateTime:    e.DateTime,
		UserId:      e.UserID,
	}, e.Id)

	if err != nil {
		return nil, err
	}

	return domain, nil
}
