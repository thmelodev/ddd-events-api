package mocks

import (
	"time"

	"github.com/thmelodev/ddd-events-api/src/modules/events/domain"
)

func MockEvent(id string, userId string) *domain.EventAggregate {
	event, _ := domain.LoadEvent(domain.EventProps{
		Name:        "Event Name",
		Description: "Event Description",
		Location:    "Event Location",
		DateTime:    time.Now(),
		UserId:      userId,
	}, id)

	return event
}
