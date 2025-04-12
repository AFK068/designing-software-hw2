package mapper

import (
	"github.com/AFK068/hsezoo/internal/domain"
	"github.com/oapi-codegen/runtime/types"

	hsezootypes "github.com/AFK068/hsezoo/internal/api/openapi/hsezoo/v1"
)

func MapToEnclosureListResponse(enclosures []*domain.Enclosure) hsezootypes.EnclosureListResponse {
	enclosureResponses := make([]hsezootypes.Enclosure, len(enclosures))

	for i, enclosure := range enclosures {
		animals := enclosure.GetAnimals()

		animalResponses := make([]hsezootypes.Animal, 0, len(animals))
		for _, animal := range animals {
			animalResponses = append(animalResponses, hsezootypes.Animal{
				Id:           types.UUID(animal.ID),
				EnclosureId:  types.UUID(enclosure.ID),
				Species:      animal.Species,
				Name:         animal.Name,
				BirthDate:    animal.BirthDate,
				Gender:       hsezootypes.AnimalGender(animal.Gender),
				FavoriteFood: string(animal.FavoriteFood),
				Status:       hsezootypes.AnimalStatus(animal.Status),
			})
		}

		enclosureResponses[i] = hsezootypes.Enclosure{
			Id:             types.UUID(enclosure.ID),
			Type:           enclosure.Type,
			Size:           enclosure.Size,
			CurrentAnimals: enclosure.CurrentAnimals,
			MaxCapacity:    enclosure.MaxCapacity,
			Animals:        &animalResponses,
		}
	}

	return hsezootypes.EnclosureListResponse{
		Enclosures: &enclosureResponses,
	}
}

func MapToTypesEnclosure(enclosure *domain.Enclosure) hsezootypes.Enclosure {
	animals := enclosure.GetAnimals()

	animalResponses := make([]hsezootypes.Animal, 0, len(animals))
	for _, animal := range animals {
		animalResponses = append(animalResponses, hsezootypes.Animal{
			Id:           types.UUID(animal.ID),
			EnclosureId:  types.UUID(enclosure.ID),
			Species:      animal.Species,
			Name:         animal.Name,
			BirthDate:    animal.BirthDate,
			Gender:       hsezootypes.AnimalGender(animal.Gender),
			FavoriteFood: string(animal.FavoriteFood),
			Status:       hsezootypes.AnimalStatus(animal.Status),
		})
	}

	return hsezootypes.Enclosure{
		Id:             types.UUID(enclosure.ID),
		Type:           enclosure.Type,
		Size:           enclosure.Size,
		CurrentAnimals: enclosure.CurrentAnimals,
		MaxCapacity:    enclosure.MaxCapacity,
		Animals:        &animalResponses,
	}
}
