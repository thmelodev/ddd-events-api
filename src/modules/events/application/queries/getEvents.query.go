package queries

import (
	"context"

	"github.com/thmelodev/ddd-events-api/src/modules/events/application/dtos"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain/mappers"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/repositories"
)

var _ IQuery = (*GetEventsQuery)(nil)

type GetEventsQuery struct {
	eventRepository repositories.IEventRepository
	eventMapper     mappers.EventMapper
}

func NewGetEventsQuery(eventRepository repositories.IEventRepository) *GetEventsQuery {
	return &GetEventsQuery{
		eventRepository: eventRepository,
	}
}

func (query GetEventsQuery) Execute(ctx context.Context, dto any) (any, error) {
	events, err := query.eventRepository.FindAll()
	if err != nil {
		return nil, err
	}

	dtos := make([]*dtos.EventDTO, len(events))

	for index, event := range events {
		eventDTO := query.eventMapper.ToDTO(event)
		dtos[index] = eventDTO
	}

	return dtos, nil
}
