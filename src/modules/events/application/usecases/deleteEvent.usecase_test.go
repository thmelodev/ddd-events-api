package usecases

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain/mocks"
)

var id = uuid.New().String()
var userId = uuid.New().String()

func MockDeleteEventUsecaseProps() DeleteEventProps {
	return DeleteEventProps{
		Id:     id,
		UserId: userId,
	}
}

func TestDeleteUsecaseSucess(t *testing.T) {

	usecase := NewDeleteEventUsecase(mocks.MockEventRepository)

	props := MockDeleteEventUsecaseProps()

	event := mocks.MockEvent(id, userId)

	mocks.MockEventRepository.On("FindById", props.Id).Return(event, nil).Once()
	mocks.MockEventRepository.On("Delete", event).Return(nil).Once()

	result, err := usecase.Execute(context.Background(), &props)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
