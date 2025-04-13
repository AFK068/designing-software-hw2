package services

import (
	"time"

	"github.com/AFK068/hsezoo/internal/domain"
	"github.com/AFK068/hsezoo/internal/domain/repositories"
	"github.com/AFK068/hsezoo/internal/infrastructure/events"
	"github.com/go-co-op/gocron/v2"
)

const (
	DefaultJobDuration = 3 * time.Second
)

type FeedingService struct {
	sheduler   gocron.Scheduler
	dispatcher events.Dispatcher
	repo       repositories.FeedingScheduleRepository
}

func NewFeedingService(repo repositories.FeedingScheduleRepository, dispatcher events.Dispatcher) (*FeedingService, error) {
	sheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &FeedingService{
		sheduler:   sheduler,
		repo:       repo,
		dispatcher: dispatcher,
	}, nil
}

func (fs *FeedingService) Run(jobDuration time.Duration) {
	_, err := fs.sheduler.NewJob(
		gocron.DurationJob(
			jobDuration,
		),
		gocron.NewTask(
			fs.feedAnimals,
		),
	)

	if err != nil {
		return
	}

	fs.sheduler.Start()
}

func (fs *FeedingService) feedAnimals() {
	schedules, err := fs.repo.GetAllFeedingSchedules()
	if err != nil {
		return
	}

	for _, schedule := range schedules {
		if schedule.FeedingTime.Before(time.Now()) && !schedule.Completed {
			event := domain.FeedingTimeEvent{
				FeedingScheduleID: schedule.ID,
				AnimalID:          schedule.Animal.ID,
				FeedingTime:       schedule.FeedingTime,
				FoodType:          schedule.FoodType,
			}

			fs.dispatcher.Dispatch(event)

			schedule.Animal.Feed(schedule.FoodType)
			schedule.MarkAsCompleted()
		}
	}
}
