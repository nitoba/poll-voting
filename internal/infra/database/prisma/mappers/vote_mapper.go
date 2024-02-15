package mappers

import (
	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/prisma/db"
)

func VoteToDomain(vote *db.VotesModel) *entities.Vote {
	return &entities.Vote{
		AggregateRoot: core.AggregateRoot{
			Entity: *core.NewEntity(core.NewUniqueEntityId(vote.ID)),
		},
		PollId:    core.NewUniqueEntityId(vote.PollID),
		VoterId:   core.NewUniqueEntityId(vote.VoterID),
		OptionId:  core.NewUniqueEntityId(vote.PollOptionID),
		CreatedAt: vote.CreatedAt,
	}
}
