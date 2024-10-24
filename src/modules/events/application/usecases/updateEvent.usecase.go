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

var _ interfaces.IUsecase = (*UpdateEventUsecase)(nil)

type UpdateEventUsecase struct {
	eventRepository repositories.IEventRepository
}

type UpdateEventDTO struct {
	Id          string    `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      string    `json:"-"`
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

	e, err := u.eventRepository.FindById(eventDTO.Id)

	if err != nil {
		return nil, err
	}

	if e.GetUserId() != eventDTO.UserId {
		return nil, apiErrors.NewUnauthorizedError("user is not the owner of this event")
	}

	err = e.UpdateEvent(domain.EventProps{
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

	return gin.H{}, nil
}
