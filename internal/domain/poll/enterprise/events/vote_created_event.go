package events

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
)

type VoteCreatedEvent struct {
	ocurredAt time.Time
	VoteId    core.UniqueEntityId
	PollId    core.UniqueEntityId
	OptionId  core.UniqueEntityId
}

func (e *VoteCreatedEvent) Name() string {
	return "event.vote.created"
}

func (e *VoteCreatedEvent) GetAggregateId() core.UniqueEntityId {
	return e.VoteId
}

func (e *VoteCreatedEvent) OcurredAt() time.Time {
	return e.ocurredAt
}

func NewVoteCreatedEvent(voteId, pollId, optionId core.UniqueEntityId) *VoteCreatedEvent {
	return &VoteCreatedEvent{
		VoteId:    voteId,
		PollId:    pollId,
		OptionId:  optionId,
		ocurredAt: time.Now(),
	}
}
