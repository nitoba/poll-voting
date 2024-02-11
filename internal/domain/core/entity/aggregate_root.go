package entity

import (
	"github.com/nitoba/poll-voting/internal/domain/core/events"
)

type AggregateRoot struct {
	Entity
	domainEvents []events.DomainEvent
}
