package usecases

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain/mocks"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

var id = uuid.New().String()
var userId = uuid.New().String()

func MockDeleteEventUsecaseProps() DeleteEventProps {
	return DeleteEventProps{
		Id:     id,
		UserId: userId,
	}
}

func sutFactory() (*DeleteEventUsecase, *mocks.EventRepositoryMock) {
	eventRepositoryMock := new(mocks.EventRepositoryMock)

	sut := NewDeleteEventUsecase(eventRepositoryMock)

	return sut, eventRepositoryMock
}

func TestDeleteUsecaseSucess(t *testing.T) {

	t.Run("delete event Sucess", func(t *testing.T) {
		usecase, eventRepositoryMock := sutFactory()

		props := MockDeleteEventUsecaseProps()

		event := mocks.MockEvent(id, userId)

		eventRepositoryMock.On("FindById", props.Id).Return(event, nil).Once()
		eventRepositoryMock.On("Delete", event).Return(nil).Once()

		result, err := usecase.Execute(context.Background(), &props)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("dto convert error", func(t *testing.T) {
		usecase, _ := sutFactory()

		props := struct{}{}

		result, err := usecase.Execute(context.Background(), props)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, apiErrors.NewInvalidPropsError(fmt.Errorf("invalid props: %v, invalid type: %t", props, props).Error()), err)
	})

	t.Run("event find by id error", func(t *testing.T) {
		usecase, eventRepositoryMock := sutFactory()

		props := MockDeleteEventUsecaseProps()

		eventRepositoryMock.On("FindById", props.Id).Return(nil, apiErrors.NewRepositoryError("")).Once()

		result, err := usecase.Execute(context.Background(), &props)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, apiErrors.NewRepositoryError(""), err)
	})

	t.Run("user not owner of event", func(t *testing.T) {
		usecase, eventRepositoryMock := sutFactory()

		props := MockDeleteEventUsecaseProps()
		props.UserId = uuid.New().String()

		event := mocks.MockEvent(id, userId)

		eventRepositoryMock.On("FindById", props.Id).Return(event, nil).Once()

		result, err := usecase.Execute(context.Background(), &props)

		assert.Error(t, err)
		assert.Nil(t, result)

		assert.Equal(t, apiErrors.NewInvalidPropsError("user is not the owner of this event"), err)
	})

	t.Run("delete event error", func(t *testing.T) {
		usecase, eventRepositoryMock := sutFactory()

		props := MockDeleteEventUsecaseProps()

		event := mocks.MockEvent(id, userId)

		eventRepositoryMock.On("FindById", props.Id).Return(event, nil).Once()
		eventRepositoryMock.On("Delete", event).Return(apiErrors.NewRepositoryError("")).Once()

		result, err := usecase.Execute(context.Background(), &props)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, apiErrors.NewRepositoryError(""), err)
	})
}
