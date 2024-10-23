package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/thmelodev/ddd-events-api/src/utils/apiErrors"
)

type EventProps struct {
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      string
}

type EventAggregate struct {
	id          string
	name        string
	description string
	location    string
	dateTime    time.Time
	userId      string
}

func New(props EventProps) (*EventAggregate, error) {
	event := &EventAggregate{id: uuid.New().String()}

	if err := event.build(props); err != nil {
		return nil, err
	}

	return event, nil
}

func Load(props EventProps, id string) (*EventAggregate, error) {

	event := &EventAggregate{}

	err := event.setId(id)

	if err != nil {
		return nil, err
	}

	err = event.build(props)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (e *EventAggregate) Update(props EventProps) error {
	return e.build(props)
}

func (e *EventAggregate) build(props EventProps) error {
	if err := e.setName(props.Name); err != nil {
		return err
	}

	if err := e.setDescription(props.Description); err != nil {
		return err
	}

	if err := e.setLocation(props.Location); err != nil {
		return err
	}

	if err := e.setDateTime(props.DateTime); err != nil {
		return err
	}

	if err := e.setUserId(props.UserId); err != nil {
		return err
	}

	return nil
}

func (e *EventAggregate) setId(id string) error {
	_, err := uuid.Parse(id)

	if err != nil {
		return apiErrors.NewInvalidPropsError("id is invalid")
	}

	e.id = id

	return nil
}

func (e *EventAggregate) setName(name string) error {
	if name == "" {
		return apiErrors.NewInvalidPropsError("name is required")
	}
	e.name = name

	return nil
}

func (e *EventAggregate) setDescription(description string) error {
	if description == "" {
		return apiErrors.NewInvalidPropsError("description is required")
	}
	e.description = description

	return nil
}

func (e *EventAggregate) setLocation(location string) error {
	if location == "" {
		return apiErrors.NewInvalidPropsError("location is required")
	}
	e.location = location

	return nil
}

func (e *EventAggregate) setDateTime(dateTime time.Time) error {
	if dateTime.IsZero() {
		return apiErrors.NewInvalidPropsError("dateTime is required")
	}
	e.dateTime = dateTime

	return nil
}

func (e *EventAggregate) setUserId(userId string) error {
	if userId == "" {
		return apiErrors.NewInvalidPropsError("userId is required")
	}
	e.userId = userId

	return nil
}

func (e *EventAggregate) GetId() string {
	return e.id
}

func (e *EventAggregate) GetName() string {
	return e.name
}

func (e *EventAggregate) GetDescription() string {
	return e.description
}

func (e *EventAggregate) GetLocation() string {
	return e.location
}

func (e *EventAggregate) GetDateTime() time.Time {
	return e.dateTime
}

func (e *EventAggregate) GetUserId() string {
	return e.userId
}
