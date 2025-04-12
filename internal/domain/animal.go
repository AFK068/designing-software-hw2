package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	AnimalStatus string
	AnimalID     uuid.UUID

	Gender string
)

const (
	Healthy AnimalStatus = "Healthy"
	Sick    AnimalStatus = "Sick"

	Male   Gender = "Male"
	Female Gender = "Female"
)

type Animal struct {
	ID           AnimalID
	EnclosureID  EnclosureID
	Species      string
	Name         string
	BirthDate    time.Time
	Gender       Gender
	FavoriteFood Food
	Status       AnimalStatus
}

func NewAnimal(
	enclosureID EnclosureID,
	species, name string,
	birthDate time.Time,
	gender Gender,
	favoriteFood Food,
	status AnimalStatus,
) *Animal {
	return &Animal{
		ID:           AnimalID(uuid.New()),
		EnclosureID:  enclosureID,
		Species:      species,
		Name:         name,
		BirthDate:    birthDate,
		Gender:       gender,
		FavoriteFood: favoriteFood,
		Status:       status,
	}
}

func (a *Animal) Move(enclosureID EnclosureID) {
	a.EnclosureID = enclosureID
}

func (a *Animal) Feed(food Food) {
	fmt.Println("Feeding", a.Name, "with", food)
}

func (a *Animal) Treat() {
	a.Status = Healthy
}
