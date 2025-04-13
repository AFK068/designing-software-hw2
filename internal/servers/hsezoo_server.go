package servers

import (
	"context"
	"log"
	"time"

	"github.com/AFK068/hsezoo/internal/application/services"
	"github.com/AFK068/hsezoo/internal/infrastructure/httpapi/hsezooapi"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	hsezootypes "github.com/AFK068/hsezoo/internal/api/openapi/hsezoo/v1"
)

type HseZoo struct {
	Handler  *hsezooapi.Handler
	Sheduler *services.FeedingService
	Echo     *echo.Echo
}

func NewHseZoo(
	hanler *hsezooapi.Handler,
	sheduler *services.FeedingService,
) *HseZoo {
	return &HseZoo{
		Handler:  hanler,
		Sheduler: sheduler,
		Echo:     echo.New(),
	}
}

func (h *HseZoo) Start() error {
	hsezootypes.RegisterHandlers(h.Echo, h.Handler)

	h.Sheduler.Run(services.DefaultJobDuration)

	return h.Echo.Start(":8080")
}

func (h *HseZoo) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return h.Echo.Shutdown(ctx)
}

func (h *HseZoo) RegisterHooks(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Println("Starting compressor server")

			go func() {
				if err := h.Start(); err != nil {
					log.Fatalf("Failed to start compressor server: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("Stopping compressor server")

			if err := h.Stop(); err != nil {
				log.Fatalf("Failed to stop compressor server: %v", err)
			}

			return nil
		},
	})
}
