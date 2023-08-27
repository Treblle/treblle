package event

import "sync"

type EventData interface{}

type EventListener func(data EventData)

type Event struct {
	Name      string
	Listeners []EventListener
	Lock      sync.RWMutex
}

func New(name string) *Event {
	return &Event{
		Name:      name,
		Listeners: []EventListener{},
	}
}

func (e *Event) Listen(listener EventListener) *Event {
	e.Lock.Lock()
	defer e.Lock.Unlock()

	e.Listeners = append(e.Listeners, listener)
	return e
}

func (e *Event) Dispatch(data EventData) {
	e.Lock.RLock()
	defer e.Lock.RUnlock()

	var wg sync.WaitGroup

	for _, listener := range e.Listeners {
		wg.Add(1)
		go func(l EventListener) {
			defer wg.Done()
			l(data)
		}(listener)
	}

	wg.Wait()
}
