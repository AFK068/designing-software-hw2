package mapper

import (
	"github.com/AFK068/hsezoo/internal/domain"
	"github.com/oapi-codegen/runtime/types"

	hsezootypes "github.com/AFK068/hsezoo/internal/api/openapi/hsezoo/v1"
)

func MapToAnimalListResponse(animals []*domain.Animal) hsezootypes.AnimalListResponse {
	animalList := make([]hsezootypes.Animal, 0, len(animals))

	for _, animal := range animals {
		typeAnimal := hsezootypes.Animal{
			BirthDate:    animal.BirthDate,
			EnclosureId:  types.UUID(animal.EnclosureID),
			FavoriteFood: string(animal.FavoriteFood),
			Gender:       hsezootypes.AnimalGender(animal.Gender),
			Id:           types.UUID(animal.ID),
			Name:         animal.Name,
			Species:      animal.Species,
			Status:       hsezootypes.AnimalStatus(animal.Status),
		}

		animalList = append(animalList, typeAnimal)
	}

	return hsezootypes.AnimalListResponse{
		Animals: &animalList,
	}
}

func MapToDomainAnimal(request *hsezootypes.AnimalInput) *domain.Animal {
	return domain.NewAnimal(
		domain.EnclosureID(request.EnclosureId),
		request.Species,
		request.Name,
		request.BirthDate,
		domain.Gender(request.Gender),
		domain.Food(request.FavoriteFood),
		domain.AnimalStatus(request.Status),
	)
}

func MapToTypesAnimal(animal *domain.Animal) hsezootypes.Animal {
	return hsezootypes.Animal{
		BirthDate:    animal.BirthDate,
		EnclosureId:  types.UUID(animal.EnclosureID),
		FavoriteFood: string(animal.FavoriteFood),
		Gender:       hsezootypes.AnimalGender(animal.Gender),
		Id:           types.UUID(animal.ID),
		Name:         animal.Name,
		Species:      animal.Species,
		Status:       hsezootypes.AnimalStatus(animal.Status),
	}
}
