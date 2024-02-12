package core

type EventHandler interface {
	// Notify(event DomainEvent)
	SetupSubscriptions()
}
