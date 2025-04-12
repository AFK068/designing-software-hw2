package domain

import "time"

type Event interface {
	Name() string
}

type EventHandler interface {
	Handle(event Event)
}

type AnimalMovedEvent struct {
	AnimalID       AnimalID
	OldEnclosureID EnclosureID
	NewEnclosureID EnclosureID
	OccurredAt     time.Time
}

func (e AnimalMovedEvent) Name() string {
	return "AnimalMoved"
}

type FeedingTimeEvent struct {
	FeedingScheduleID FeedingScheduleID
	AnimalID          AnimalID
	FeedingTime       time.Time
	FoodType          Food
}

func (e FeedingTimeEvent) Name() string {
	return "FeedingTime"
}
