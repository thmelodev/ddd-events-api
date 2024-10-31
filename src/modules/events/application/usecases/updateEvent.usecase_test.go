package usecases

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain/mocks"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

func MockUpdateEventUseCaseDto() *UpdateEventDTO {
	return &UpdateEventDTO{
		Id:          uuid.NewString(),
		Name:        "Event Name",
		Description: "Event Description",
		Location:    "Event Location",
		DateTime:    time.Now(),
		UserId:      uuid.NewString(),
	}
}

func sutFacoryUpdateEventUsecase() (*UpdateEventUsecase, *mocks.EventRepositoryMock) {
	eventRepositoryMock := new(mocks.EventRepositoryMock)

	sut := NewUpdateEventUsecase(eventRepositoryMock)

	return sut, eventRepositoryMock
}

func TestUpdateEventUsecase(t *testing.T) {
	t.Run("create update event usecase", func(t *testing.T) {
		sut, eventRepositoryMock := sutFacoryUpdateEventUsecase()

		props := MockUpdateEventUseCaseDto()

		eventRepositoryMock.On("FindById", mock.Anything).Return(mocks.MockEvent(uuid.NewString(), props.UserId), nil)
		eventRepositoryMock.On("Save", mock.Anything).Return(nil)

		result, err := sut.Execute(context.Background(), props)

		require.Nil(t, err)
		require.NotNil(t, result)
	})

	t.Run("create update event usecase with invalid dto", func(t *testing.T) {
		sut, _ := sutFacoryUpdateEventUsecase()

		props := struct{}{}

		result, err := sut.Execute(context.Background(), props)

		require.Error(t, err)
		require.Nil(t, result)
		assert.Equal(t, apiErrors.NewInvalidPropsError(fmt.Errorf("invalid props: %v, invalid type: %t", props, props).Error()), err)
	})

	t.Run("event findById error", func(t *testing.T) {
		sut, eventRepositoryMock := sutFacoryUpdateEventUsecase()

		props := MockUpdateEventUseCaseDto()

		eventRepositoryMock.On("FindById", mock.Anything).Return(nil, apiErrors.NewRepositoryError(""))

		result, err := sut.Execute(context.Background(), props)

		require.Error(t, err)
		require.Nil(t, result)
		assert.Equal(t, apiErrors.NewRepositoryError(""), err)
	})

	t.Run("user not owner of event", func(t *testing.T) {
		sut, eventRepositoryMock := sutFacoryUpdateEventUsecase()

		props := MockUpdateEventUseCaseDto()

		eventRepositoryMock.On("FindById", mock.Anything).Return(mocks.MockEvent(uuid.NewString(), uuid.NewString()), nil)
		eventRepositoryMock.On("Save", mock.Anything).Return(nil)

		result, err := sut.Execute(context.Background(), props)

		require.Error(t, err)
		require.Nil(t, result)
		assert.Equal(t, apiErrors.NewUnauthorizedError("user is not the owner of this event"), err)
	})

	t.Run("update event error", func(t *testing.T) {
		sut, eventRepositoryMock := sutFacoryUpdateEventUsecase()

		props := MockUpdateEventUseCaseDto()
		props.Name = ""

		eventRepositoryMock.On("FindById", mock.Anything).Return(mocks.MockEvent(uuid.NewString(), props.UserId), nil)

		result, err := sut.Execute(context.Background(), props)

		require.Error(t, err)
		require.Nil(t, result)
		assert.Equal(t, apiErrors.NewInvalidPropsError("name is required"), err)
	})

	t.Run("save event error", func(t *testing.T) {
		sut, eventRepositoryMock := sutFacoryUpdateEventUsecase()

		props := MockUpdateEventUseCaseDto()

		eventRepositoryMock.On("FindById", mock.Anything).Return(mocks.MockEvent(uuid.NewString(), props.UserId), nil)
		eventRepositoryMock.On("Save", mock.Anything).Return(apiErrors.NewRepositoryError(""))

		result, err := sut.Execute(context.Background(), props)

		require.Error(t, err)
		require.Nil(t, result)
		assert.Equal(t, apiErrors.NewRepositoryError(""), err)
	})
}
