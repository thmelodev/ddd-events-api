package usecases

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/repositories"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
	"github.com/thmelodev/ddd-events-api/src/utils/interfaces"
)

var _ interfaces.IUsecase = (*DeleteEventUsecase)(nil)

type DeleteEventUsecase struct {
	eventRepository repositories.IEventRepository
}

func NewDeleteEventUsecase(
	eventRepository repositories.IEventRepository,
) *DeleteEventUsecase {
	return &DeleteEventUsecase{
		eventRepository: eventRepository,
	}
}

type DeleteEventDTO struct {
	Id     string `json:"-"`
	UserId string `json:"-"`
}

func (u DeleteEventUsecase) Execute(ctx context.Context, dto any) (any, error) {
	eventDTO, ok := dto.(*DeleteEventDTO)

	if !ok {
		return nil, apiErrors.NewInvalidPropsError(fmt.Errorf("invalid props: %v, invalid type: %t", dto, dto).Error())
	}

	event, err := u.eventRepository.FindById(eventDTO.Id)

	if err != nil {
		return nil, err
	}

	if event.GetUserId() != eventDTO.UserId {
		return nil, apiErrors.NewInvalidPropsError("user is not the owner of this event")
	}

	if err = u.eventRepository.Delete(event); err != nil {
		return nil, err
	}

	return gin.H{}, nil
}
