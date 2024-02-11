package events

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core/entity"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type VoteChangedEvent struct {
	ocurredAt time.Time
	Vote      entities.Vote
}

func (e *VoteChangedEvent) Name() string {
	return "event.vote.changed"
}

func (e *VoteChangedEvent) GetAggregateId() entity.UniqueEntityId {
	return e.Vote.Id
}

func (e *VoteChangedEvent) OcurredAt() time.Time {
	return e.ocurredAt
}

func NewVoteChangedEvent(vote entities.Vote) *VoteChangedEvent {
	return &VoteChangedEvent{
		Vote:      vote,
		ocurredAt: time.Now(),
	}
}
