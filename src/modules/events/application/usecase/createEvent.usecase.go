package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/repositories"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

var _ IUsecase = (*CreateEventUsecase)(nil)

type CreateEventUsecase struct {
	eventRepository repositories.IEventRepository
}

type CreateEventDTO struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserID      string    `json:"userID"`
}

func NewCreateEventUsecase(
	eventRepository repositories.IEventRepository,
) *CreateEventUsecase {
	return &CreateEventUsecase{
		eventRepository: eventRepository,
	}
}

func (u CreateEventUsecase) Execute(ctx context.Context, dto any) (any, error) {
	eventDTO, ok := dto.(*CreateEventDTO)
	if !ok {
		return nil, apiErrors.NewInvalidPropsException(fmt.Errorf("invalid props: %v", dto).Error())
	}

	event, err := domain.NewEvent(domain.EventProps{
		Name:        eventDTO.Name,
		Description: eventDTO.Description,
		Location:    eventDTO.Location,
		DateTime:    eventDTO.DateTime,
		UserId:      eventDTO.UserID,
	})

	if err != nil {
		return nil, err
	}

	if err = u.eventRepository.Create(event); err != nil {
		return nil, err
	}

	return nil, nil
}
