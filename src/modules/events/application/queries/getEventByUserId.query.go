package queries

import (
	"context"
	"fmt"

	"github.com/thmelodev/ddd-events-api/src/modules/events/application/dtos"
	"github.com/thmelodev/ddd-events-api/src/modules/events/application/mappers"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain/repositories"
	"github.com/thmelodev/ddd-events-api/src/utils/interfaces"
)

var _ interfaces.IQuery = (*GetEventByUserIdQuery)(nil)

type GetEventByUserIdQuery struct {
	eventRepository repositories.IEventRepository
	eventMapper     mappers.EventMapper
}

func NewGetEventByUserIdQuery(eventRepository repositories.IEventRepository) *GetEventByUserIdQuery {
	return &GetEventByUserIdQuery{
		eventRepository: eventRepository,
	}
}

func (query GetEventByUserIdQuery) Execute(ctx context.Context, dto any) (any, error) {
	userId := fmt.Sprint(dto)

	events, err := query.eventRepository.FindByUserId(userId)
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
