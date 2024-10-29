package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain/mocks"
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

func TestCreateEventUsecaseSucess(t *testing.T) {
	usecase := NewCreateEventUsecase(mocks.MockEventRepository)

	props := MockCreateEventUsecaseProps()

	mocks.MockEventRepository.On("Save", mock.Anything).Return(nil).Once()

	result, err := usecase.Execute(context.Background(), props)

	id, exists := result.(gin.H)["id"]

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, id)
	assert.True(t, exists)
}
