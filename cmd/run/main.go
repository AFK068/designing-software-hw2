package main

import (
	"log"

	"github.com/AFK068/hsezoo/internal/application/services"
	"github.com/AFK068/hsezoo/internal/domain/repositories"
	"github.com/AFK068/hsezoo/internal/infrastructure/eventhandlers"
	"github.com/AFK068/hsezoo/internal/infrastructure/events"
	"github.com/AFK068/hsezoo/internal/infrastructure/httpapi/hsezooapi"
	"github.com/AFK068/hsezoo/internal/infrastructure/repository/inmemoryrepo"
	"github.com/AFK068/hsezoo/internal/servers"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			func() events.Dispatcher {
				dispatcher := events.NewEventDispatcher()

				animalMovedHandler := eventhandlers.NewAnimalMovedHandler()
				feedingHandler := eventhandlers.NewFeedingHandler()

				dispatcher.RegisterHandler("AnimalMoved", animalMovedHandler)
				dispatcher.RegisterHandler("FeedingTime", feedingHandler)

				return dispatcher
			},

			func() (repositories.AnimalRepository, repositories.EnclosureRepository, repositories.FeedingScheduleRepository) {
				return inmemoryrepo.NewInMemoryAnimalRepository(),
					inmemoryrepo.NewInMemoryEnclosureRepository(),
					inmemoryrepo.NewInMemoryFeedingScheduleRepository()
			},

			func(
				animalRepo repositories.AnimalRepository,
				enclosureRepo repositories.EnclosureRepository,
				feedingScheduleRepo repositories.FeedingScheduleRepository,
			) (hsezooapi.ZooStatisticsProvider, hsezooapi.ZooTransferProvider) {
				statisticsService := services.NewStatisticsService(animalRepo, enclosureRepo, feedingScheduleRepo)
				animalTransferService := services.NewAnimalTransferService(animalRepo, enclosureRepo, events.NewEventDispatcher())

				return statisticsService, animalTransferService
			},

			func(
				feedingScheduleRepo repositories.FeedingScheduleRepository,
				dispatcher events.Dispatcher,
			) *services.FeedingService {
				feedingService, err := services.NewFeedingService(feedingScheduleRepo, dispatcher)
				if err != nil {
					log.Fatal("Failed to create feeding service")
				}
				return feedingService
			},

			hsezooapi.NewHandler,

			servers.NewHseZoo,
		),
		fx.Invoke(
			func(s *servers.HseZoo, lc fx.Lifecycle) {
				s.RegisterHooks(lc)
			},
		),
	).Run()
}
