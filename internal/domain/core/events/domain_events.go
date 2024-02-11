package events

type DomainEventCallback func(event DomainEvent)

type domainEvents struct {
	handlers map[string][]DomainEventCallback
	// markedAggregates []interface{}
}

var instance *domainEvents = &domainEvents{
	handlers: map[string][]DomainEventCallback{},
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

func DomainEvents() *domainEvents {
	return instance
}
