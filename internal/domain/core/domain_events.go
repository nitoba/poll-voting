package core

import (
	"slices"
	"sync"
)

type DomainEventCallback func(event interface{}, wg *sync.WaitGroup)

type domainEvents struct {
	handlers         map[string][]DomainEventCallback
	markedAggregates []interface{}
}

var instance *domainEvents = &domainEvents{
	handlers: map[string][]DomainEventCallback{},
}

func (e *domainEvents) MarkAggregateForDispatch(aggregate interface{}) {
	aggregateFound := e.findMarkedAggregateById(aggregate.(*AggregateRoot).Id)

	if aggregateFound == nil {
		e.markedAggregates = append(e.markedAggregates, aggregate)
	}
}

func (e *domainEvents) DispatchEventsForAggregate(id UniqueEntityId) {
	aggregate := e.findMarkedAggregateById(id)
	if aggregate != nil {
		e.dispatchAggregateEvents(aggregate)
		aggregate.(*AggregateRoot).ClearEvents()
		e.removeAggregateFromMarkedDispatchList(aggregate)
	}
}

func (e *domainEvents) Register(handler DomainEventCallback, eventName string) {
	wasEventRegisteredBefore := e.handlers[eventName] != nil

	if !wasEventRegisteredBefore {
		e.handlers[eventName] = []DomainEventCallback{}
	}

	e.handlers[eventName] = append(e.handlers[eventName], handler)
}

func (e *domainEvents) ClearHandlers() {
	e.handlers = map[string][]DomainEventCallback{}
}

func (e *domainEvents) ClearMarkedAggregates() {
	e.markedAggregates = []interface{}{}
}

func (e *domainEvents) findMarkedAggregateById(id UniqueEntityId) interface{} {
	for _, aggregate := range e.markedAggregates {
		if aggregate.(*AggregateRoot).Id.Equals(id) {
			return aggregate
		}
	}

	return nil
}

func (e *domainEvents) dispatchAggregateEvents(aggregate interface{}) {
	domainEvents := aggregate.(*AggregateRoot).GetEvents()
	for _, event := range domainEvents {
		e.dispatch(event)
	}
}

func (e *domainEvents) removeAggregateFromMarkedDispatchList(aggregate interface{}) {
	aggregateIndex := slices.Index(e.markedAggregates, aggregate)
	if aggregateIndex != -1 {
		slices.DeleteFunc(e.markedAggregates, func(a interface{}) bool {
			return a.(*AggregateRoot).Id == aggregate.(*AggregateRoot).Id
		})
	}
}

func (e *domainEvents) dispatch(event DomainEvent) {
	eventName := event.Name()
	if _, ok := e.handlers[eventName]; ok {
		wg := &sync.WaitGroup{}
		handlers := e.handlers[eventName]
		for _, handler := range handlers {
			wg.Add(1)
			handler(event, wg)
		}
		wg.Wait()
	}
}

// DomainEvents returns the singleton instance of the domain events.
func DomainEvents() *domainEvents {
	return instance
}
