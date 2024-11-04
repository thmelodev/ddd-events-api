package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

func TestNewEvent(t *testing.T) {

	t.Run("create event success", func(t *testing.T) {
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
		assert.Equal(t, eventProps.Name, event.GetName())
		assert.Equal(t, eventProps.Description, event.GetDescription())
		assert.Equal(t, eventProps.Location, event.GetLocation())
		assert.Equal(t, eventProps.DateTime, event.GetDateTime())
		assert.Equal(t, eventProps.UserId, event.GetUserId())
	})

	t.Run("create event error because userId length === 0", func(t *testing.T) {
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
	})

	t.Run("create event error because name length === 0", func(t *testing.T) {
		eventProps := EventProps{
			Name:        "",
			Description: "Test Description",
			Location:    "Test Location",
			DateTime:    time.Now(),
			UserId:      uuid.New().String(),
		}

		event, err := NewEvent(eventProps)

		assert.Error(t, err)
		assert.Nil(t, event)
		assert.Equal(t, apiErrors.NewInvalidPropsError("name is required"), err)
	})

	t.Run("create event error because description length === 0", func(t *testing.T) {
		eventProps := EventProps{
			Name:        "Test Event",
			Description: "",
			Location:    "Test Location",
			DateTime:    time.Now(),
			UserId:      uuid.New().String(),
		}

		event, err := NewEvent(eventProps)

		assert.Error(t, err)
		assert.Nil(t, event)
		assert.Equal(t, apiErrors.NewInvalidPropsError("description is required"), err)
	})

	t.Run("create event error because location length === 0", func(t *testing.T) {
		eventProps := EventProps{
			Name:        "Test Event",
			Description: "Test Description",
			Location:    "",
			DateTime:    time.Now(),
			UserId:      uuid.New().String(),
		}

		event, err := NewEvent(eventProps)

		assert.Error(t, err)
		assert.Nil(t, event)
		assert.Equal(t, apiErrors.NewInvalidPropsError("location is required"), err)
	})

	t.Run("create event error because dateTime is zero", func(t *testing.T) {
		eventProps := EventProps{
			Name:        "Test Event",
			Description: "Test Description",
			Location:    "Test Location",
			DateTime:    time.Time{},
			UserId:      uuid.New().String(),
		}

		event, err := NewEvent(eventProps)

		assert.Error(t, err)
		assert.Nil(t, event)
		assert.Equal(t, apiErrors.NewInvalidPropsError("dateTime is required"), err)
	})

	t.Run("create event error because id is invalid", (func(t *testing.T) {
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
	}))

	t.Run("update event success", func(t *testing.T) {
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
	})

	t.Run("Load event sucess", func(t *testing.T) {
		id := uuid.New().String()

		eventProps := EventProps{
			Name:        "Test Event",
			Description: "Test Description",
			Location:    "Test Location",
			DateTime:    time.Now(),
			UserId:      "test",
		}

		event, err := LoadEvent(eventProps, id)

		assert.NoError(t, err)
		assert.NotNil(t, event)
		assert.Equal(t, event.GetId(), event.GetId())
		assert.Equal(t, eventProps.Name, event.GetName())
		assert.Equal(t, eventProps.Description, event.GetDescription())
		assert.Equal(t, eventProps.Location, event.GetLocation())
		assert.Equal(t, eventProps.DateTime, event.GetDateTime())
		assert.Equal(t, eventProps.UserId, event.GetUserId())
	})

	t.Run("Load event error build", (func(t *testing.T) {
		eventProps := EventProps{
			Name:        "",
			Description: "Test Description",
			Location:    "Test Location",
			DateTime:    time.Now(),
			UserId:      "test",
		}

		event, err := LoadEvent(eventProps, uuid.NewString())

		assert.Error(t, err)
		assert.Nil(t, event)
		assert.Equal(t, apiErrors.NewInvalidPropsError("name is required"), err)
	}))

}
