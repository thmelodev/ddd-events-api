package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/repositories"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

var _ IUsecase = (*UpdateEventUsecase)(nil)

type UpdateEventUsecase struct {
	eventRepository repositories.IEventRepository
}

type UpdateEventDTO struct {
	Id          string    `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      string    `json:"userId"`
}

func NewUpdateEventUsecase(
	eventRepository repositories.IEventRepository,
) *UpdateEventUsecase {
	return &UpdateEventUsecase{
		eventRepository: eventRepository,
	}
}

func (u UpdateEventUsecase) Execute(ctx context.Context, dto any) (any, error) {
	eventDTO, ok := dto.(*UpdateEventDTO)

	if !ok {
		return nil, apiErrors.NewInvalidPropsError(fmt.Errorf("invalid props: %v, invalid type: %t", dto, dto).Error())
	}

	event, err := u.eventRepository.FindById(eventDTO.Id)

	if err != nil {
		return nil, err
	}

	err = event.UpdateEvent(domain.EventProps{
		Name:        eventDTO.Name,
		Description: eventDTO.Description,
		Location:    eventDTO.Location,
		DateTime:    eventDTO.DateTime,
		UserId:      eventDTO.UserId,
	})

	if err != nil {
		return nil, err
	}

	if err = u.eventRepository.Save(event); err != nil {
		return nil, err
	}

	return gin.H{}, nil
}
