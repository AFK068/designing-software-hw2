package eventhandlers

import (
	"log"

	"github.com/AFK068/hsezoo/internal/domain"
)

type FeedingHandler struct{}

func NewFeedingHandler() *FeedingHandler {
	return &FeedingHandler{}
}

func (h *FeedingHandler) Handle(event domain.Event) {
	switch e := event.(type) {
	case domain.FeedingTimeEvent:
		log.Printf("Feeding time event: Animal %s, Time %s", e.AnimalID, e.FeedingTime)
	default:
		log.Printf("Unknown event type: %T", e)
	}
}
