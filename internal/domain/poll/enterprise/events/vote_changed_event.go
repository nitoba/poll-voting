package events

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
)

type VoteChangedEvent struct {
	ocurredAt time.Time
	VoteId    core.UniqueEntityId
	PollId    core.UniqueEntityId
	OptionId  core.UniqueEntityId
}

func (e *VoteChangedEvent) Name() string {
	return "event.vote.changed"
}

func (e *VoteChangedEvent) GetAggregateId() core.UniqueEntityId {
	return e.VoteId
}

func (e *VoteChangedEvent) OcurredAt() time.Time {
	return e.ocurredAt
}

func NewVoteChangedEvent(voteId, pollId, optionId core.UniqueEntityId) *VoteChangedEvent {
	return &VoteChangedEvent{
		VoteId:    voteId,
		PollId:    pollId,
		OptionId:  optionId,
		ocurredAt: time.Now(),
	}
}
