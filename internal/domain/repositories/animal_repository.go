package repositories

import "github.com/AFK068/hsezoo/internal/domain"

type AnimalRepository interface {
	AddAnimal(animal *domain.Animal) error
	DeleteAnimal(id domain.AnimalID) error

	GetAnimal(id domain.AnimalID) (*domain.Animal, error)
	GetAllAnimals() ([]*domain.Animal, error)
	GetCountOfAnimals() (int, error)
}
