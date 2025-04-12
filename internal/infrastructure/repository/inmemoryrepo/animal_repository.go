package inmemoryrepo

import (
	"errors"
	"sync"

	"github.com/AFK068/hsezoo/internal/domain"
)

type InMemoryAnimalRepository struct {
	animals map[domain.AnimalID]*domain.Animal
	mu      sync.RWMutex
}

func NewInMemoryAnimalRepository() *InMemoryAnimalRepository {
	return &InMemoryAnimalRepository{
		animals: make(map[domain.AnimalID]*domain.Animal),
	}
}

func (r *InMemoryAnimalRepository) AddAnimal(animal *domain.Animal) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.animals[animal.ID] = animal

	return nil
}

func (r *InMemoryAnimalRepository) DeleteAnimal(id domain.AnimalID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.animals[id]; !exists {
		return errors.New("animal not found")
	}

	delete(r.animals, id)

	return nil
}

func (r *InMemoryAnimalRepository) GetAnimal(id domain.AnimalID) (*domain.Animal, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	animal, exists := r.animals[id]
	if !exists {
		return nil, errors.New("animal not found")
	}

	return animal, nil
}

func (r *InMemoryAnimalRepository) GetAllAnimals() ([]*domain.Animal, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	animals := make([]*domain.Animal, 0, len(r.animals))
	for _, animal := range r.animals {
		animals = append(animals, animal)
	}

	return animals, nil
}

func (r *InMemoryAnimalRepository) GetCountOfAnimals() (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.animals), nil
}
