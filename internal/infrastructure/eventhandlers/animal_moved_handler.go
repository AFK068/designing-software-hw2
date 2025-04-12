package eventhandlers

import (
	"log"

	"github.com/AFK068/hsezoo/internal/domain"
)

type AnimalMovedHandler struct{}

func NewAnimalMovedHandler() *AnimalMovedHandler {
	return &AnimalMovedHandler{}
}

func (h *AnimalMovedHandler) Handle(event domain.Event) {
	switch e := event.(type) {
	case domain.AnimalMovedEvent:
		log.Printf("Animal %s moved: %s to enclosure %d", e.AnimalID, e.OldEnclosureID, e.NewEnclosureID)
	default:
		log.Printf("Unknown event type: %T", e)
	}
}
