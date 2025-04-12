package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	FeedingScheduleID uuid.UUID
	Food              string
)

type FeedingSchedule struct {
	ID          FeedingScheduleID
	Animal      *Animal
	FeedingTime time.Time
	FoodType    Food
	Completed   bool
}

func NewFeedingSchedule(animal *Animal, feedingTime time.Time, foodType Food) *FeedingSchedule {
	return &FeedingSchedule{
		ID:          FeedingScheduleID(uuid.New()),
		Animal:      animal,
		FeedingTime: feedingTime,
		FoodType:    foodType,
		Completed:   false,
	}
}

func (fs *FeedingSchedule) UpdateFeedingTime(newTime time.Time) {
	fs.FeedingTime = newTime
}

func (fs *FeedingSchedule) MarkAsCompleted() {
	fs.Completed = true
}
