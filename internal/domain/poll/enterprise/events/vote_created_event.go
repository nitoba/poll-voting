package events

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type VoteCreatedEvent struct {
	ocurredAt time.Time
	Vote      entities.Vote
}

func (e *VoteCreatedEvent) Name() string {
	return "event.vote.created"
}

func (e *VoteCreatedEvent) GetAggregateId() core.UniqueEntityId {
	return e.Vote.Id
}

func (e *VoteCreatedEvent) OcurredAt() time.Time {
	return e.ocurredAt
}

func NewVoteCreatedEvent(vote entities.Vote) *VoteCreatedEvent {
	return &VoteCreatedEvent{
		Vote:      vote,
		ocurredAt: time.Now(),
	}
}
