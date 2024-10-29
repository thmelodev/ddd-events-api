package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

func TestNewEventSucess(t *testing.T) {
	userId := uuid.New().String()

	eventProps := EventProps{
		Name:        "Test Event",
		Description: "Test Description",
		Location:    "Test Location",
		DateTime:    time.Now(),
		UserId:      userId,
	}

	event, err := NewEvent(eventProps)

	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, eventProps.Name, event.name)
	assert.Equal(t, eventProps.Description, event.description)
	assert.Equal(t, eventProps.Location, event.location)
	assert.Equal(t, eventProps.DateTime, event.dateTime)
	assert.Equal(t, eventProps.UserId, event.userId)
}

func TestNewEventError(t *testing.T) {
	eventProps := EventProps{
		Name:        "Test Event",
		Description: "Test Description",
		Location:    "Test Location",
		DateTime:    time.Now(),
		UserId:      "",
	}

	event, err := NewEvent(eventProps)

	assert.Error(t, err)
	assert.Nil(t, event)
	assert.Equal(t, apiErrors.NewInvalidPropsError("userId is required"), err)
}

func TestLoadEventErrorId(t *testing.T) {
	eventProps := EventProps{
		Name:        "Test Event",
		Description: "Test Description",
		Location:    "Test Location",
		DateTime:    time.Now(),
		UserId:      "test",
	}

	event, err := LoadEvent(eventProps, "")

	assert.Error(t, err)
	assert.Nil(t, event)
	assert.Equal(t, apiErrors.NewInvalidPropsError("id is invalid"), err)
}

func TestLoadEventErrorName(t *testing.T) {
	eventProps := EventProps{
		Name:        "",
		Description: "Test Description",
		Location:    "Test Location",
		DateTime:    time.Now(),
		UserId:      "test",
	}

	event, err := LoadEvent(eventProps, uuid.New().String())

	assert.Error(t, err)
	assert.Nil(t, event)
	assert.Equal(t, apiErrors.NewInvalidPropsError("name is required"), err)
}

func TestUpdateEventSuccess(t *testing.T) {
	eventProps := EventProps{
		Name:        "Test Event",
		Description: "Test Description",
		Location:    "Test Location",
		DateTime:    time.Now(),
		UserId:      "test",
	}

	event, _ := NewEvent(eventProps)

	updateEventProps := EventProps{
		Name:        "Test Event Updated",
		Description: "Test Description Updated",
		Location:    "Test Location Updated",
		DateTime:    time.Now(),
	}

	err := event.UpdateEvent(updateEventProps)

	assert.NoError(t, err)
	assert.Equal(t, updateEventProps.Name, event.name)
	assert.Equal(t, updateEventProps.Description, event.description)
	assert.Equal(t, updateEventProps.Location, event.location)
	assert.Equal(t, updateEventProps.DateTime, event.dateTime)
}
