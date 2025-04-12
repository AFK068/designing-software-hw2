package repositories

import "github.com/AFK068/hsezoo/internal/domain"

type EnclosureRepository interface {
	AddEnclosure(enclosure *domain.Enclosure) error
	DeleteEnclosure(id domain.EnclosureID) error

	GetEnclosure(id domain.EnclosureID) (*domain.Enclosure, error)
	GetAllEnclosures() ([]*domain.Enclosure, error)
	GetCountOfEnclosures() (int, error)
	GetCountFreeEnclosures() (int, error)

	SaveAnimalToEnclosure(enclosureID domain.EnclosureID, animal *domain.Animal) error
	GetAnimalsByEnclosureID(enclosureID domain.EnclosureID) ([]*domain.Animal, error)
}
