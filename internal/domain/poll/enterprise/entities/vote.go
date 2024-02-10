package entities

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
)

type Vote struct {
	Id        core.UniqueEntityId
	PollId    core.UniqueEntityId
	OptionId  core.UniqueEntityId
	VoterId   core.UniqueEntityId
	CreatedAt time.Time
}

func (v *Vote) Equals(other *Vote) bool {
	if v == nil || other == nil {
		return false
	}

	if v.PollId.String() != other.PollId.String() || v.OptionId.String() != other.OptionId.String() || v.VoterId.String() != other.VoterId.String() {
		return false
	}

	if v == other {
		return true
	}

	if v.Id.String() == other.Id.String() {
		return true
	}

	return false
}

type OptionalVoteParams struct {
	Id        *core.UniqueEntityId
	CreatedAt *time.Time
	PollId    *core.UniqueEntityId
	OptionId  *core.UniqueEntityId
	VoterId   *core.UniqueEntityId
}

func NewVote(pollId core.UniqueEntityId, optionId core.UniqueEntityId, voterId core.UniqueEntityId, optional ...OptionalVoteParams) (*Vote, error) {
	var id core.UniqueEntityId
	var optionIdVo core.UniqueEntityId
	var pollIdVo core.UniqueEntityId
	var voterIdVo core.UniqueEntityId
	var createdAt time.Time
	if len(optional) > 0 && optional[0].Id != nil {
		id = core.NewUniqueEntityId(optional[0].Id.String())
	} else {
		id = core.NewUniqueEntityId()
	}

	if len(optional) > 0 && optional[0].PollId != nil {
		pollIdVo = core.NewUniqueEntityId(optional[0].PollId.String())
	} else {
		pollIdVo = pollId
	}

	if len(optional) > 0 && optional[0].OptionId != nil {
		optionIdVo = core.NewUniqueEntityId(optional[0].OptionId.String())
	} else {
		optionIdVo = optionId
	}

	if len(optional) > 0 && optional[0].VoterId != nil {
		voterIdVo = core.NewUniqueEntityId(optional[0].VoterId.String())
	} else {
		voterIdVo = voterId
	}

	if len(optional) > 0 && optional[0].CreatedAt != nil {
		createdAt = *optional[0].CreatedAt
	} else {
		createdAt = time.Now()
	}

	return &Vote{
		Id:        id,
		PollId:    pollIdVo,
		OptionId:  optionIdVo,
		VoterId:   voterIdVo,
		CreatedAt: createdAt,
	}, nil
}
