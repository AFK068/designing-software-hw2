package services

import (
	"time"

	"github.com/AFK068/hsezoo/internal/domain"
	"github.com/AFK068/hsezoo/internal/domain/repositories"
	"github.com/AFK068/hsezoo/internal/infrastructure/events"
)

type timeGetter func() time.Time

type AnimalTransferService struct {
	animalRepository    repositories.AnimalRepository
	enclosureRepository repositories.EnclosureRepository
	timeGetter          timeGetter
	dispatcher          events.Dispatcher
}

func NewAnimalTransferService(
	animalRepo repositories.AnimalRepository,
	enclosureRepo repositories.EnclosureRepository,
	dispatcher events.Dispatcher,
) *AnimalTransferService {
	return &AnimalTransferService{
		animalRepository:    animalRepo,
		enclosureRepository: enclosureRepo,
		timeGetter:          time.Now,
		dispatcher:          dispatcher,
	}
}

func (s *AnimalTransferService) MoveAnimalToEnclosure(animalID domain.AnimalID, enclosureID domain.EnclosureID) error {
	animal, err := s.animalRepository.GetAnimal(animalID)
	if err != nil {
		return err
	}

	enclosure, err := s.enclosureRepository.GetEnclosure(enclosureID)
	if err != nil {
		return err
	}

	if err := enclosure.AddAnimal(animal); err != nil {
		return err
	}

	oldEnclosure, err := s.enclosureRepository.GetEnclosure(animal.EnclosureID)
	if err != nil {
		return err
	}

	if err := oldEnclosure.RemoveAnimal(animalID); err != nil {
		return err
	}

	animal.Move(enclosureID)

	event := domain.AnimalMovedEvent{
		AnimalID:       animal.ID,
		OldEnclosureID: animal.EnclosureID,
		NewEnclosureID: enclosure.ID,
		OccurredAt:     s.timeGetter(),
	}

	s.dispatcher.Dispatch(event)

	return nil
}
