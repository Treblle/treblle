package dispatcher

import (
	"sync"

	"github.com/treblle/treblle/pkg/event"
)

type Dispatcher struct {
	Events map[string]*event.Event
	Lock   sync.RWMutex
}

func New() *Dispatcher {
	return &Dispatcher{
		Events: map[string]*event.Event{},
	}
}

func (d *Dispatcher) Event(name string) *event.Event {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	if event, exists := d.Events[name]; exists {
		return event
	}

	event := event.New(name)
	d.Events[name] = event
	return event
}
