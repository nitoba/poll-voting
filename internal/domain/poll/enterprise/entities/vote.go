package entities

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core/entity"
)

type Vote struct {
	Id        entity.UniqueEntityId
	PollId    entity.UniqueEntityId
	OptionId  entity.UniqueEntityId
	VoterId   entity.UniqueEntityId
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
	Id        *entity.UniqueEntityId
	CreatedAt *time.Time
	PollId    *entity.UniqueEntityId
	OptionId  *entity.UniqueEntityId
	VoterId   *entity.UniqueEntityId
}

func (v *Vote) IsTheSameVoteOption(optionId string) bool {
	return v.OptionId.String() == optionId
}
func (v *Vote) ChangeVoteOption(optionId string) {
	v.OptionId = entity.NewUniqueEntityId(optionId)
}

func NewVote(pollId entity.UniqueEntityId, optionId entity.UniqueEntityId, voterId entity.UniqueEntityId, optional ...OptionalVoteParams) (*Vote, error) {
	var id entity.UniqueEntityId
	var optionIdVo entity.UniqueEntityId
	var pollIdVo entity.UniqueEntityId
	var voterIdVo entity.UniqueEntityId
	var createdAt time.Time
	if len(optional) > 0 && optional[0].Id != nil {
		id = entity.NewUniqueEntityId(optional[0].Id.String())
	} else {
		id = entity.NewUniqueEntityId()
	}

	if len(optional) > 0 && optional[0].PollId != nil {
		pollIdVo = entity.NewUniqueEntityId(optional[0].PollId.String())
	} else {
		pollIdVo = pollId
	}

	if len(optional) > 0 && optional[0].OptionId != nil {
		optionIdVo = entity.NewUniqueEntityId(optional[0].OptionId.String())
	} else {
		optionIdVo = optionId
	}

	if len(optional) > 0 && optional[0].VoterId != nil {
		voterIdVo = entity.NewUniqueEntityId(optional[0].VoterId.String())
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
