package mapper

import (
	"github.com/AFK068/hsezoo/internal/domain"
	"github.com/oapi-codegen/runtime/types"

	hsezootypes "github.com/AFK068/hsezoo/internal/api/openapi/hsezoo/v1"
)

func MapToFeedingScheduleListResponse(schedules []*domain.FeedingSchedule) hsezootypes.FeedingScheduleListResponse {
	feedingScheduleResponses := make([]hsezootypes.FeedingSchedule, len(schedules))

	for i, schedule := range schedules {
		feedingScheduleResponses[i] = hsezootypes.FeedingSchedule{
			Id:          types.UUID(schedule.ID),
			Animal:      MapToTypesAnimal(schedule.Animal),
			FeedingTime: schedule.FeedingTime,
			FoodType:    string(schedule.FoodType),
			Completed:   schedule.Completed,
		}
	}

	return hsezootypes.FeedingScheduleListResponse{
		Schedules: &feedingScheduleResponses,
	}
}

func MapToTypesFeedingSchedule(schedule *domain.FeedingSchedule) hsezootypes.FeedingSchedule {
	return hsezootypes.FeedingSchedule{
		Id:          types.UUID(schedule.ID),
		Animal:      MapToTypesAnimal(schedule.Animal),
		FeedingTime: schedule.FeedingTime,
		FoodType:    string(schedule.FoodType),
		Completed:   schedule.Completed,
	}
}
