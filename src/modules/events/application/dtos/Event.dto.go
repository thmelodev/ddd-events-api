package dtos

import "time"

type EventDTO struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description" `
	Location    string    `json:"location" `
	DateTime    time.Time `json:"dateTime" `
	UserId      string    `json:"userId" `
}
