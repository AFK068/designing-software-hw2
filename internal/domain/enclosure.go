package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type (
	EnclosureID uuid.UUID
)

type Enclosure struct {
	ID             EnclosureID
	Animals        map[AnimalID]*Animal
	Type           string
	Size           int
	CurrentAnimals int
	MaxCapacity    int
}

func NewEnclosure(enclosureType string, size, maxCapacity int) *Enclosure {
	return &Enclosure{
		ID:             EnclosureID(uuid.New()),
		Animals:        make(map[AnimalID]*Animal),
		Type:           enclosureType,
		Size:           size,
		CurrentAnimals: 0,
		MaxCapacity:    maxCapacity,
	}
}

func (e *Enclosure) AddAnimal(animal *Animal) error {
	if e.CurrentAnimals == e.MaxCapacity {
		return errors.New("enclosure is full")
	}

	if _, exists := e.Animals[animal.ID]; exists {
		return errors.New("animal already in enclosure")
	}

	e.Animals[animal.ID] = animal
	e.CurrentAnimals++

	return nil
}

func (e *Enclosure) RemoveAnimal(animalID AnimalID) error {
	if _, exists := e.Animals[animalID]; !exists {
		return errors.New("animal not found in enclosure")
	}

	delete(e.Animals, animalID)

	e.CurrentAnimals--

	return nil
}

func (e *Enclosure) GetAnimals() []*Animal {
	animals := make([]*Animal, 0, len(e.Animals))
	for _, animal := range e.Animals {
		animals = append(animals, animal)
	}

	return animals
}

func (e *Enclosure) Clean() {
	fmt.Println("Cleaning enclosure...")
}
