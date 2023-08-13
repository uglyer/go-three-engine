package core

import (
	"sync"
	"sync/atomic"
)

type EventDispatcher struct {
	listeners map[string]map[int]func(Event)
	mu        sync.Mutex
}

type Event struct {
	Type   string
	Target *EventDispatcher
	Data   interface{}
}

func (ed *EventDispatcher) AddEventListener(eventType string, listener func(Event)) int {
	ed.mu.Lock()
	defer ed.mu.Unlock()

	if ed.listeners == nil {
		ed.listeners = make(map[string]map[int]func(Event))
	}

	eventId := generateEventId()
	if _, found := ed.listeners[eventType]; !found {
		ed.listeners[eventType] = make(map[int]func(Event))
	}
	ed.listeners[eventType][eventId] = listener

	return eventId
}

func (ed *EventDispatcher) RemoveEventListener(eventType string, eventId int) {
	ed.mu.Lock()
	defer ed.mu.Unlock()

	if eventTypeListeners, found := ed.listeners[eventType]; found {
		delete(eventTypeListeners, eventId)
	}
}

func (ed *EventDispatcher) DispatchEvent(eventType string, data interface{}) {
	ed.mu.Lock()
	defer ed.mu.Unlock()

	if eventTypeListeners, found := ed.listeners[eventType]; found {
		event := Event{
			Type:   eventType,
			Target: ed,
			Data:   data,
		}

		// Make a copy of listeners to avoid issues if listeners are removed while iterating.
		listenersCopy := make([]func(Event), len(eventTypeListeners))
		index := 0
		for _, listener := range eventTypeListeners {
			listenersCopy[index] = listener
			index++
		}

		for _, listener := range listenersCopy {
			listener(event)
		}
	}
}

var eventId atomic.Uint32

func generateEventId() int {
	return int(eventId.Add(1))
}
