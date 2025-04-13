package hsezooapi

import (
	"github.com/AFK068/hsezoo/internal/application/mapper"
	"github.com/AFK068/hsezoo/internal/domain"
	"github.com/AFK068/hsezoo/internal/domain/repositories"
	"github.com/labstack/echo/v4"

	hsezootypes "github.com/AFK068/hsezoo/internal/api/openapi/hsezoo/v1"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type ZooStatisticsProvider interface {
	GetAnimalCount() (int, error)
	GetFreeEnclosureCount() (int, error)
	GetEnclosureCount() (int, error)
	GetFeedingScheduleCount() (int, error)
}

type ZooTransferProvider interface {
	MoveAnimalToEnclosure(animalID domain.AnimalID, enclosureID domain.EnclosureID) error
}

type Handler struct {
	animalRepository          repositories.AnimalRepository
	enclosureRepository       repositories.EnclosureRepository
	feedingScheduleRepository repositories.FeedingScheduleRepository
	statisticsService         ZooStatisticsProvider
	animalTransferService     ZooTransferProvider
}

func NewHandler(
	animalRepo repositories.AnimalRepository,
	enclosureRepo repositories.EnclosureRepository,
	feedingScheduleRepo repositories.FeedingScheduleRepository,
	statisticsService ZooStatisticsProvider,
	animalTransferService ZooTransferProvider,
) *Handler {
	return &Handler{
		animalRepository:          animalRepo,
		enclosureRepository:       enclosureRepo,
		feedingScheduleRepository: feedingScheduleRepo,
		statisticsService:         statisticsService,
		animalTransferService:     animalTransferService,
	}
}

func (h *Handler) GetAnimals(ctx echo.Context) error {
	animals, err := h.animalRepository.GetAllAnimals()
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get animals")
	}

	if len(animals) == 0 {
		return SendNotFoundResponse(ctx, "No animals found", "No animals found")
	}

	response := mapper.MapToAnimalListResponse(animals)

	return SendSuccessResponse(ctx, response)
}

func (h *Handler) PostAnimals(ctx echo.Context) error {
	var request hsezootypes.AnimalInput
	if err := ctx.Bind(&request); err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to bind request")
	}

	animal := mapper.MapToDomainAnimal(&request)

	if err := h.animalRepository.AddAnimal(animal); err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to add animal")
	}

	if err := h.enclosureRepository.SaveAnimalToEnclosure(domain.EnclosureID(request.EnclosureId), animal); err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to save animal to enclosure")
	}

	response := mapper.MapToTypesAnimal(animal)

	return SendSuccessResponse(ctx, response)
}

func (h *Handler) DeleteAnimalsAnimalId(ctx echo.Context, animalId openapi_types.UUID) error { //nolint
	err := h.animalRepository.DeleteAnimal(domain.AnimalID(animalId))
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to delete animal")
	}

	return SendSuccessResponse(ctx, nil)
}

func (h *Handler) GetAnimalsAnimalId(ctx echo.Context, animalId openapi_types.UUID) error { //nolint
	animal, err := h.animalRepository.GetAnimal(domain.AnimalID(animalId))
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get animal")
	}

	if animal == nil {
		return SendNotFoundResponse(ctx, "Animal not found", "Animal not found")
	}

	response := mapper.MapToTypesAnimal(animal)

	return SendSuccessResponse(ctx, response)
}

func (h *Handler) PostAnimalsAnimalIdMove(ctx echo.Context, animalId openapi_types.UUID) error { //nolint
	var request hsezootypes.MoveAnimalInput
	if err := ctx.Bind(&request); err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to bind request")
	}

	err := h.animalTransferService.MoveAnimalToEnclosure(
		domain.AnimalID(animalId),
		domain.EnclosureID(request.NewEnclosureId),
	)
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to move animal")
	}

	return SendSuccessResponse(ctx, nil)
}

func (h *Handler) GetEnclosures(ctx echo.Context) error {
	enclosures, err := h.enclosureRepository.GetAllEnclosures()
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get enclosures")
	}

	if len(enclosures) == 0 {
		return SendNotFoundResponse(ctx, "No enclosures found", "No enclosures found")
	}

	response := mapper.MapToEnclosureListResponse(enclosures)

	return SendSuccessResponse(ctx, response)
}

func (h *Handler) PostEnclosures(ctx echo.Context) error {
	var request hsezootypes.EnclosureInput
	if err := ctx.Bind(&request); err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to bind request")
	}

	enclosure := domain.NewEnclosure(
		request.Type,
		request.Size,
		request.MaxCapacity,
	)

	if err := h.enclosureRepository.AddEnclosure(enclosure); err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to add enclosure")
	}

	response := mapper.MapToTypesEnclosure(enclosure)

	return SendSuccessResponse(ctx, response)
}

func (h *Handler) DeleteEnclosuresEnclosureId(ctx echo.Context, enclosureId openapi_types.UUID) error { //nolint
	err := h.enclosureRepository.DeleteEnclosure(domain.EnclosureID(enclosureId))
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to delete enclosure")
	}

	return SendSuccessResponse(ctx, nil)
}

func (h *Handler) GetEnclosuresEnclosureId(ctx echo.Context, enclosureId openapi_types.UUID) error { //nolint
	enclosure, err := h.enclosureRepository.GetEnclosure(domain.EnclosureID(enclosureId))
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get enclosure")
	}

	if enclosure == nil {
		return SendNotFoundResponse(ctx, "Enclosure not found", "Enclosure not found")
	}

	response := mapper.MapToTypesEnclosure(enclosure)

	return SendSuccessResponse(ctx, response)
}

func (h *Handler) GetFeedingSchedules(ctx echo.Context) error {
	schedules, err := h.feedingScheduleRepository.GetAllFeedingSchedules()
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get feeding schedules")
	}

	if len(schedules) == 0 {
		return SendNotFoundResponse(ctx, "No feeding schedules found", "No feeding schedules found")
	}

	response := mapper.MapToFeedingScheduleListResponse(schedules)

	return SendSuccessResponse(ctx, response)
}

func (h *Handler) PostFeedingSchedules(ctx echo.Context) error {
	var request hsezootypes.FeedingScheduleInput
	if err := ctx.Bind(&request); err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to bind request")
	}

	animal, err := h.animalRepository.GetAnimal(domain.AnimalID(request.AnimalId))
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get animal")
	}

	schedule := domain.NewFeedingSchedule(
		animal,
		request.FeedingTime,
		domain.Food(request.FoodType),
	)

	if err := h.feedingScheduleRepository.AddFeedingSchedule(schedule); err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to add feeding schedule")
	}

	response := mapper.MapToTypesFeedingSchedule(schedule)

	return SendSuccessResponse(ctx, response)
}

func (h *Handler) DeleteFeedingSchedulesScheduleId(ctx echo.Context, scheduleId openapi_types.UUID) error { //nolint
	err := h.feedingScheduleRepository.DeleteFeedingSchedule(domain.FeedingScheduleID(scheduleId))
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to delete feeding schedule")
	}

	return SendSuccessResponse(ctx, nil)
}

func (h *Handler) GetFeedingSchedulesScheduleId(ctx echo.Context, scheduleId openapi_types.UUID) error { //nolint
	feedingSchedule, err := h.feedingScheduleRepository.GetFeedingSchedule(domain.FeedingScheduleID(scheduleId))
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get feeding schedule")
	}

	if feedingSchedule == nil {
		return SendNotFoundResponse(ctx, "Feeding schedule not found", "Feeding schedule not found")
	}

	response := mapper.MapToTypesFeedingSchedule(feedingSchedule)

	return SendSuccessResponse(ctx, response)
}

func (h *Handler) GetStatistics(ctx echo.Context) error {
	animalCount, err := h.statisticsService.GetAnimalCount()
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get animal count")
	}

	freeEnclosureCount, err := h.statisticsService.GetFreeEnclosureCount()
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get free enclosure count")
	}

	shouldEnclosureCount, err := h.statisticsService.GetEnclosureCount()
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get enclosure count")
	}

	feedingScheduleCount, err := h.statisticsService.GetFeedingScheduleCount()
	if err != nil {
		return SendBadRequestResponse(ctx, "Error occurred", "Failed to get feeding schedule count")
	}

	response := hsezootypes.ZooStatistics{
		FeedingSchedulesCount: feedingScheduleCount,
		FreeEnclosures:        freeEnclosureCount,
		TotalAnimals:          animalCount,
		TotalEnclosures:       shouldEnclosureCount,
	}

	return SendSuccessResponse(ctx, response)
}
