package events

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
)

type DomainEvent interface {
	Name() string
	OcurredAt() time.Time
	GetAggregateId() core.UniqueEntityId
}
