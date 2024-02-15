package repositories

import (
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma/mappers"
	"github.com/nitoba/poll-voting/prisma/db"
)

type VoteRepositoryPrisma struct {
	db *db.PrismaClient
}

func (r *VoteRepositoryPrisma) Create(vote *entities.Vote) error {
	ctx := configs.GetConfig().Ctx
	_, err := r.db.Votes.CreateOne(
		db.Votes.Poll.Link(db.Poll.ID.Equals(vote.PollId.String())),
		db.Votes.PollOption.Link(db.PollOption.ID.Equals(vote.OptionId.String())),
		db.Votes.Voter.Link(db.Voter.ID.Equals(vote.VoterId.String())),
		db.Votes.ID.Set(vote.Id.String()),
	).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
func (r *VoteRepositoryPrisma) Delete(vote *entities.Vote) error {
	ctx := configs.GetConfig().Ctx
	_, err := r.db.Votes.FindUnique(db.Votes.ID.Equals(vote.Id.String())).Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *VoteRepositoryPrisma) FindByPollIdAndVoterId(pollId string, voterId string) (*entities.Vote, error) {
	ctx := configs.GetConfig().Ctx

	vote, err := r.db.Votes.FindUnique(
		db.Votes.PollIDVoterID(
			db.Votes.PollID.Equals(pollId),
			db.Votes.VoterID.Equals(voterId),
		),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return mappers.VoteToDomain(vote), nil
}

func NewVoteRepositoryPrisma(db *db.PrismaClient) *VoteRepositoryPrisma {
	return &VoteRepositoryPrisma{
		db: db,
	}
}
