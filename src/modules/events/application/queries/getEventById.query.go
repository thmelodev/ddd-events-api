package queries

import (
	"context"
	"fmt"

	"github.com/thmelodev/ddd-events-api/src/modules/events/application/mappers"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/repositories"
)

var _ IQuery = (*GetEventByIdQuery)(nil)

type GetEventByIdQuery struct {
	eventRepository repositories.IEventRepository
	eventMapper     mappers.EventMapper
}

func NewGetEventsQueryById(eventRepository repositories.IEventRepository) *GetEventByIdQuery {
	return &GetEventByIdQuery{
		eventRepository: eventRepository,
	}
}

func (query GetEventByIdQuery) Execute(ctx context.Context, dto any) (any, error) {
	id := fmt.Sprint(dto)

	event, err := query.eventRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	eventDTO := query.eventMapper.ToDTO(event)

	return eventDTO, nil
}
