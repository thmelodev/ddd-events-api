package models

import "time"

type Event struct {
	Id          string    `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      string    `json:"user_id" binding:"required"`
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func createEvent() {

}

func GetAllEvents() []Event {
	return events
}
