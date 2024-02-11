package events

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core/entity"
)

type DomainEvent interface {
	Name() string
	OcurredAt() time.Time
	GetAggregateId() entity.UniqueEntityId
}
