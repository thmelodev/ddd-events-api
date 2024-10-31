package usecases

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain/mocks"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

func MockCreateEventUsecaseProps() *CreateEventUsecaseProps {
	props := CreateEventUsecaseProps{
		Name:        "Event Name",
		Description: "Event Description",
		Location:    "Event Location",
		DateTime:    time.Now(),
		UserId:      "user-id",
	}

	return &props
}

func sutFactoryCreateEventUsecase() (*CreateEventUsecase, *mocks.EventRepositoryMock) {
	eventRepositoryMock := new(mocks.EventRepositoryMock)

	sut := NewCreateEventUsecase(eventRepositoryMock)

	return sut, eventRepositoryMock
}

func TestCreateEventUsecase(t *testing.T) {
	t.Run("create event Sucess", func(t *testing.T) {
		usecase, eventRepositoryMock := sutFactoryCreateEventUsecase()

		props := MockCreateEventUsecaseProps()

		eventRepositoryMock.On("Save", mock.Anything).Return(nil).Once()

		result, err := usecase.Execute(context.Background(), props)

		id, exists := result.(gin.H)["id"]

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, id)
		assert.True(t, exists)
	})

	t.Run("create event dto convert error", func(t *testing.T) {
		usecase, _ := sutFactoryCreateEventUsecase()

		props := struct{}{}

		result, err := usecase.Execute(context.Background(), props)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, apiErrors.NewRepositoryError(fmt.Errorf("invalid props: %v, invalid type: %t", props, props).Error()), err)
	})

	t.Run("create event error", func(t *testing.T) {
		usecase, _ := sutFactoryCreateEventUsecase()

		props := MockCreateEventUsecaseProps()
		props.Name = ""

		result, err := usecase.Execute(context.Background(), props)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, apiErrors.NewInvalidPropsError("name is required"), err)
	})

	t.Run("save event error", func(t *testing.T) {
		usecase, eventRepositoryMock := sutFactoryCreateEventUsecase()

		props := MockCreateEventUsecaseProps()

		eventRepositoryMock.On("Save", mock.Anything).Return(apiErrors.NewRepositoryError("")).Once()

		result, err := usecase.Execute(context.Background(), props)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, apiErrors.NewRepositoryError(""), err)
	})

}
