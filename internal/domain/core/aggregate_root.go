package core

type AggregateRoot struct {
	Entity
	domainEvents []DomainEvent
}

func (ar AggregateRoot) GetEvents() []DomainEvent {
	return ar.domainEvents
}

func (ar *AggregateRoot) AddDomainEvent(event DomainEvent) {
	ar.domainEvents = append(ar.domainEvents, event)
	DomainEvents().MarkAggregateForDispatch(ar)
}

func (ar *AggregateRoot) ClearEvents() {
	ar.domainEvents = []DomainEvent{}
}
