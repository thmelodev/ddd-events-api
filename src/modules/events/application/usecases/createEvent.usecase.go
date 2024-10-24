package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/repositories"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
	"github.com/thmelodev/ddd-events-api/src/utils/interfaces"
)

var _ interfaces.IUsecase = (*CreateEventUsecase)(nil)

type CreateEventUsecase struct {
	eventRepository repositories.IEventRepository
}

type CreateEventDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      string    `json:"-"`
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
		return nil, apiErrors.NewRepositoryError(fmt.Errorf("invalid props: %v, invalid type: %t", dto, dto).Error())
	}

	e, err := domain.NewEvent(domain.EventProps{
		Name:        eventDTO.Name,
		Description: eventDTO.Description,
		Location:    eventDTO.Location,
		DateTime:    eventDTO.DateTime,
		UserId:      eventDTO.UserId,
	})

	if err != nil {
		return nil, err
	}

	if err = u.eventRepository.Save(e); err != nil {
		return nil, err
	}

	return gin.H{"id": e.GetId()}, nil
}
