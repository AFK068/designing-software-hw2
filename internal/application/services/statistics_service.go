package services

import "github.com/AFK068/hsezoo/internal/domain/repositories"

type StatisticsService struct {
	animalRepo          repositories.AnimalRepository
	enclosureRepo       repositories.EnclosureRepository
	feedingScheduleRepo repositories.FeedingScheduleRepository
}

func NewStatisticsService(
	animalRepo repositories.AnimalRepository,
	enclosureRepo repositories.EnclosureRepository,
	feedingScheduleRepo repositories.FeedingScheduleRepository,
) *StatisticsService {
	return &StatisticsService{
		animalRepo:          animalRepo,
		enclosureRepo:       enclosureRepo,
		feedingScheduleRepo: feedingScheduleRepo,
	}
}

func (s *StatisticsService) GetAnimalCount() (int, error) {
	return s.animalRepo.GetCountOfAnimals()
}

func (s *StatisticsService) GetFreeEnclosureCount() (int, error) {
	return s.enclosureRepo.GetCountFreeEnclosures()
}

func (s *StatisticsService) GetEnclosureCount() (int, error) {
	return s.enclosureRepo.GetCountOfEnclosures()
}

func (s *StatisticsService) GetFeedingScheduleCount() (int, error) {
	return s.feedingScheduleRepo.GetCountOfFeedingSchedules()
}
