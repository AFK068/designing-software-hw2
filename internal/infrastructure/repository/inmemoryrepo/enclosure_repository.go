package inmemoryrepo

import (
	"errors"
	"sync"

	"github.com/AFK068/hsezoo/internal/domain"
)

type InMemoryEnclosureRepository struct {
	enclosures map[domain.EnclosureID]*domain.Enclosure
	mu         sync.RWMutex
}

func NewInMemoryEnclosureRepository() *InMemoryEnclosureRepository {
	return &InMemoryEnclosureRepository{
		enclosures: make(map[domain.EnclosureID]*domain.Enclosure),
	}
}

func (r *InMemoryEnclosureRepository) AddEnclosure(enclosure *domain.Enclosure) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.enclosures[enclosure.ID] = enclosure

	return nil
}

func (r *InMemoryEnclosureRepository) DeleteEnclosure(id domain.EnclosureID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.enclosures[id]; !exists {
		return errors.New("enclosure not found")
	}

	delete(r.enclosures, id)

	return nil
}

func (r *InMemoryEnclosureRepository) GetEnclosure(id domain.EnclosureID) (*domain.Enclosure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	enclosure, exists := r.enclosures[id]
	if !exists {
		return nil, errors.New("enclosure not found")
	}

	return enclosure, nil
}

func (r *InMemoryEnclosureRepository) SaveAnimalToEnclosure(enclosureID domain.EnclosureID, animal *domain.Animal) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	enclosure, exists := r.enclosures[enclosureID]
	if !exists {
		return errors.New("enclosure not found")
	}

	if enclosure.CurrentAnimals >= enclosure.MaxCapacity {
		return errors.New("enclosure is full")
	}

	err := enclosure.AddAnimal(animal)
	if err != nil {
		return err
	}

	return nil
}

func (r *InMemoryEnclosureRepository) GetAllEnclosures() ([]*domain.Enclosure, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	enclosures := make([]*domain.Enclosure, 0, len(r.enclosures))
	for _, enclosure := range r.enclosures {
		enclosures = append(enclosures, enclosure)
	}

	return enclosures, nil
}

func (r *InMemoryEnclosureRepository) GetAnimalsByEnclosureID(enclosureID domain.EnclosureID) ([]*domain.Animal, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	animals := r.enclosures[enclosureID].GetAnimals()

	return animals, nil
}

func (r *InMemoryEnclosureRepository) GetCountOfEnclosures() (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.enclosures), nil
}

func (r *InMemoryEnclosureRepository) GetCountFreeEnclosures() (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	count := 0

	for _, enclosure := range r.enclosures {
		if enclosure.CurrentAnimals < enclosure.MaxCapacity {
			count++
		}
	}

	return count, nil
}
