package inmemoryrepo

import (
	"errors"
	"sync"

	"github.com/AFK068/hsezoo/internal/domain"
)

type InMemoryFeedingScheduleRepository struct {
	feedingSchedules map[domain.FeedingScheduleID]*domain.FeedingSchedule
	mu               sync.RWMutex
}

func NewInMemoryFeedingScheduleRepository() *InMemoryFeedingScheduleRepository {
	return &InMemoryFeedingScheduleRepository{
		feedingSchedules: make(map[domain.FeedingScheduleID]*domain.FeedingSchedule),
	}
}

func (r *InMemoryFeedingScheduleRepository) AddFeedingSchedule(schedule *domain.FeedingSchedule) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.feedingSchedules[schedule.ID] = schedule

	return nil
}

func (r *InMemoryFeedingScheduleRepository) DeleteFeedingSchedule(id domain.FeedingScheduleID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.feedingSchedules[id]; !exists {
		return errors.New("feeding schedule not found")
	}

	delete(r.feedingSchedules, id)

	return nil
}

func (r *InMemoryFeedingScheduleRepository) GetFeedingSchedule(id domain.FeedingScheduleID) (*domain.FeedingSchedule, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	schedule, exists := r.feedingSchedules[id]
	if !exists {
		return nil, errors.New("feeding schedule not found")
	}

	return schedule, nil
}

func (r *InMemoryFeedingScheduleRepository) GetCountOfFeedingSchedules() (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.feedingSchedules), nil
}

func (r *InMemoryFeedingScheduleRepository) GetAllFeedingSchedules() ([]*domain.FeedingSchedule, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	schedules := make([]*domain.FeedingSchedule, 0, len(r.feedingSchedules))
	for _, schedule := range r.feedingSchedules {
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}
