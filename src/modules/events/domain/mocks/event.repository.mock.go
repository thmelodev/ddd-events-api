package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
)

type EventRepositoryMock struct {
	mock.Mock
}

func (m *EventRepositoryMock) Save(event *domain.EventAggregate) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *EventRepositoryMock) FindById(id string) (*domain.EventAggregate, error) {
	args := m.Called(id)
	var domainEvent *domain.EventAggregate
	if args.Get(0) != nil {
		domainEvent = args.Get(0).(*domain.EventAggregate)
	}

	return domainEvent, args.Error(1)
}

func (m *EventRepositoryMock) FindByUserId(id string) ([]*domain.EventAggregate, error) {
	args := m.Called(id)
	var domainEvent []*domain.EventAggregate
	if args.Get(0) != nil {
		domainEvent = args.Get(0).([]*domain.EventAggregate)
	}

	return domainEvent, args.Error(1)
}

func (m *EventRepositoryMock) FindAll() ([]*domain.EventAggregate, error) {
	args := m.Called()
	var domainEvent []*domain.EventAggregate
	if args.Get(0) != nil {
		domainEvent = args.Get(0).([]*domain.EventAggregate)
	}

	return domainEvent, args.Error(1)
}

func (m *EventRepositoryMock) Delete(event *domain.EventAggregate) error {
	args := m.Called(event)
	return args.Error(0)
}
