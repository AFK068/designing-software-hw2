package repositories

import "github.com/AFK068/hsezoo/internal/domain"

type FeedingScheduleRepository interface {
	AddFeedingSchedule(schedule *domain.FeedingSchedule) error
	DeleteFeedingSchedule(id domain.FeedingScheduleID) error

	GetFeedingSchedule(id domain.FeedingScheduleID) (*domain.FeedingSchedule, error)
	GetCountOfFeedingSchedules() (int, error)
	GetAllFeedingSchedules() ([]*domain.FeedingSchedule, error)
}
