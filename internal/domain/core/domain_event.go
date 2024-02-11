package core

import (
	"time"
)

type DomainEvent interface {
	Name() string
	OcurredAt() time.Time
	GetAggregateId() UniqueEntityId
}
