package factories

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type OptionalVoteParams struct {
	Id        *core.UniqueEntityId
	PollId    *core.UniqueEntityId
	OptionId  *core.UniqueEntityId
	VoterId   *core.UniqueEntityId
	CreatedAt *time.Time
}

func MakeVote(props ...OptionalVoteParams) *entities.Vote {

	var id core.UniqueEntityId = core.NewUniqueEntityId()
	var pollIdVo core.UniqueEntityId = core.NewUniqueEntityId()
	var optionIdVo core.UniqueEntityId = core.NewUniqueEntityId()
	var voterIdVo core.UniqueEntityId = core.NewUniqueEntityId()
	var createdAt time.Time = time.Now()

	if len(props) > 0 && props[0].Id != nil {
		id = core.NewUniqueEntityId(props[0].Id.String())
	}
	if len(props) > 0 && props[0].PollId != nil {
		pollIdVo = core.NewUniqueEntityId(props[0].PollId.String())
	}
	if len(props) > 0 && props[0].OptionId != nil {
		optionIdVo = core.NewUniqueEntityId(props[0].OptionId.String())
	}
	if len(props) > 0 && props[0].VoterId != nil {
		voterIdVo = core.NewUniqueEntityId(props[0].VoterId.String())
	}
	if len(props) > 0 && props[0].CreatedAt != nil {
		createdAt = *props[0].CreatedAt
	}

	return &entities.Vote{
		Id:        id,
		PollId:    pollIdVo,
		OptionId:  optionIdVo,
		VoterId:   voterIdVo,
		CreatedAt: createdAt,
	}
}
