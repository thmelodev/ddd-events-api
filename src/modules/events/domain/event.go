package domain

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

func NewEvent(props EventProps) (*EventAggregate, error) {

	event := &EventAggregate{id: uuid.New().String()}

	if err := event.setName(props.Name); err != nil {
		return nil, err
	}

	if err := event.setDescription(props.Description); err != nil {
		return nil, err
	}

	if err := event.setLocation(props.Location); err != nil {
		return nil, err
	}

	if err := event.setDateTime(props.DateTime); err != nil {
		return nil, err
	}

	if err := event.setUserId(props.UserId); err != nil {
		return nil, err
	}

	return event, nil
}

func LoadEvent(props EventProps, id string) (*EventAggregate, error) {

	event := &EventAggregate{}

	if err := event.setId(id); err != nil {
		return nil, err
	}

	if err := event.setName(props.Name); err != nil {
		return nil, err
	}

	if err := event.setDescription(props.Description); err != nil {
		return nil, err
	}

	if err := event.setLocation(props.Location); err != nil {
		return nil, err
	}

	if err := event.setDateTime(props.DateTime); err != nil {
		return nil, err
	}

	if err := event.setUserId(props.UserId); err != nil {
		return nil, err
	}

	return event, nil
}

func (e *EventAggregate) setId(id string) error {
	_, err := uuid.Parse(id)

	if err != nil {
		return apiErrors.NewInvalidPropsException("id is invalid")
	}

	e.id = id

	return nil
}

func (e *EventAggregate) setName(name string) error {
	if name == "" {
		return apiErrors.NewInvalidPropsException("name is required")
	}
	e.name = name

	return nil
}

func (e *EventAggregate) setDescription(description string) error {
	if description == "" {
		return apiErrors.NewInvalidPropsException("description is required")
	}
	e.description = description

	return nil
}

func (e *EventAggregate) setLocation(location string) error {
	if location == "" {
		return apiErrors.NewInvalidPropsException("location is required")
	}
	e.location = location

	return nil
}

func (e *EventAggregate) setDateTime(dateTime time.Time) error {
	if dateTime.IsZero() {
		return apiErrors.NewInvalidPropsException("dateTime is required")
	}
	e.dateTime = dateTime

	return nil
}

func (e *EventAggregate) setUserId(userId string) error {
	if userId == "" {
		return apiErrors.NewInvalidPropsException("userId is required")
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
