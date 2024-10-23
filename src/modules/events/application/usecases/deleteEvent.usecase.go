package usecases

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/repositories"
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

func (u DeleteEventUsecase) Execute(ctx context.Context, dto any) (any, error) {
	id := fmt.Sprint(dto)

	event, err := u.eventRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	if err = u.eventRepository.Delete(event); err != nil {
		return nil, err
	}

	return gin.H{}, nil
}
