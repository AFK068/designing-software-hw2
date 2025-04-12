package events

import (
	"sync"

	"github.com/AFK068/hsezoo/internal/domain"
)

type Dispatcher interface {
	RegisterHandler(eventType string, handler domain.EventHandler)
	Dispatch(event domain.Event)
}

type EventDispatcher struct {
	handlers map[string][]domain.EventHandler
	mu       sync.RWMutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]domain.EventHandler),
	}
}

func (d *EventDispatcher) RegisterHandler(eventType string, handler domain.EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.handlers[eventType]; !exists {
		d.handlers[eventType] = []domain.EventHandler{}
	}

	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

func (d *EventDispatcher) Dispatch(event domain.Event) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if handlers, ok := d.handlers[event.Name()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
}
