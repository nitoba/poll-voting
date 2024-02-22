package repositories

import (
	"fmt"

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

	pollTx := r.db.Poll.CreateOne(
		db.Poll.Title.Set(poll.Title),
		db.Poll.Owner.Link(db.Voter.ID.Equals(poll.OwnerId.String())),
		db.Poll.ID.Set(poll.Id.String()),
	).Tx()

	baseQuery := "INSERT INTO poll_options (id, title, poll_id) VALUES "
	values := []interface{}{}

	for i, option := range poll.Options {
		if i != 0 {
			baseQuery += ", "
		}
		baseQuery += fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3)
		values = append(values, option.Id.String(), option.Title, poll.Id.String())
	}

	optionsTx := r.db.Prisma.ExecuteRaw(baseQuery, values...).Tx()

	if err := r.db.Prisma.Transaction(pollTx, optionsTx).Exec(ctx); err != nil {
		return err
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

func (r *PollsRepositoryPrisma) FindManyByOwnerId(ownerId string) ([]*entities.Poll, error) {
	ctx := configs.GetConfig().Ctx
	polls, err := r.db.Poll.FindMany(db.Poll.OwnerID.Equals(ownerId)).With(db.Poll.Options.Fetch()).Exec(ctx)

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
