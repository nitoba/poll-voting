package factories

import (
	"time"

	"github.com/nitoba/poll-voting/internal/domain/core/entity"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type OptionalVoteParams struct {
	Id        *entity.UniqueEntityId
	PollId    *entity.UniqueEntityId
	OptionId  *entity.UniqueEntityId
	VoterId   *entity.UniqueEntityId
	CreatedAt *time.Time
}

func MakeVote(props ...OptionalVoteParams) *entities.Vote {

	var id entity.UniqueEntityId = entity.NewUniqueEntityId()
	var pollIdVo entity.UniqueEntityId = entity.NewUniqueEntityId()
	var optionIdVo entity.UniqueEntityId = entity.NewUniqueEntityId()
	var voterIdVo entity.UniqueEntityId = entity.NewUniqueEntityId()
	var createdAt time.Time = time.Now()

	if len(props) > 0 && props[0].Id != nil {
		id = entity.NewUniqueEntityId(props[0].Id.String())
	}
	if len(props) > 0 && props[0].PollId != nil {
		pollIdVo = entity.NewUniqueEntityId(props[0].PollId.String())
	}
	if len(props) > 0 && props[0].OptionId != nil {
		optionIdVo = entity.NewUniqueEntityId(props[0].OptionId.String())
	}
	if len(props) > 0 && props[0].VoterId != nil {
		voterIdVo = entity.NewUniqueEntityId(props[0].VoterId.String())
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
