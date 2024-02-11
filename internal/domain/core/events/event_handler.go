package events

type EventHandler interface {
	Notify(event DomainEvent)
	SetupSubscriptions()
}
