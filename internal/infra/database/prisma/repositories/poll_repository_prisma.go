package repositories

import (
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma/mappers"
	"github.com/nitoba/poll-voting/prisma/db"
)

type PollsRepositoryPrisma struct {
	db *db.PrismaClient
}

func (r *PollsRepositoryPrisma) Create(poll *entities.Poll) error {
	ctx := configs.GetConfig().Ctx

	pollCreated, err := r.db.Poll.CreateOne(
		db.Poll.Title.Set(poll.Title),
		db.Poll.Owner.Link(db.Voter.ID.Equals(poll.OwnerId.String())),
		db.Poll.ID.Set(poll.Id.String()),
	).Exec(ctx)

	if err != nil {
		return err
	}

	for _, option := range poll.Options {
		_, err = r.db.PollOption.CreateOne(
			db.PollOption.Title.Set(option.Title),
			db.PollOption.Poll.Link(
				db.Poll.ID.Equals(pollCreated.ID),
			),
			db.PollOption.ID.Set(option.Id.String()),
		).Exec(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PollsRepositoryPrisma) FindById(id string) (*entities.Poll, error) {
	ctx := configs.GetConfig().Ctx
	poll, err := r.db.Poll.FindUnique(db.Poll.ID.Equals(id)).With(db.Poll.Options.Fetch()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return mappers.PollToDomain(poll), nil
}

func (r *PollsRepositoryPrisma) FindMany() ([]*entities.Poll, error) {
	ctx := configs.GetConfig().Ctx
	polls, err := r.db.Poll.FindMany().With(db.Poll.Options.Fetch()).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return mappers.PollToDomainList(polls), nil
}

func NewPollsRepositoryPrisma(db *db.PrismaClient) *PollsRepositoryPrisma {
	return &PollsRepositoryPrisma{
		db: db,
	}
}
