package mappers

import (
	"github.com/thmelodev/ddd-events-api/src/modules/events/application/dtos"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
)

type EventMapper struct{}

func NewEventMapper() *EventMapper {
	return &EventMapper{}
}

func (m *EventMapper) ToDTO(event *domain.EventAggregate) *dtos.EventDTO {
	return &dtos.EventDTO{
		Id:          event.GetId(),
		Name:        event.GetName(),
		Description: event.GetDescription(),
		Location:    event.GetLocation(),
		DateTime:    event.GetDateTime(),
		UserId:      event.GetUserId(),
	}
}
