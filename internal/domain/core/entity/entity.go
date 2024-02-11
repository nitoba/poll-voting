package entity

import "github.com/nitoba/poll-voting/internal/domain/core"

type Entity struct {
	Id core.UniqueEntityId
	// Equals(other BaseEntity) bool
}
