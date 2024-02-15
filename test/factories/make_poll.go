package factories

import (
	"time"

	"github.com/jaswdr/faker"
	configs "github.com/nitoba/poll-voting/config"
	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	"github.com/nitoba/poll-voting/prisma/db"
)

type OptionalPollParams struct {
	Id        *core.UniqueEntityId
	Title     string
	OwnerId   *core.UniqueEntityId
	CreatedAt *time.Time
	Options   []*entities.PollOption
}

func MakePool(props ...OptionalPollParams) *entities.Poll {
	fake := faker.New()
	id := core.NewUniqueEntityId()
	title := fake.Lorem().Text(100)
	pollOptions := []*entities.PollOption{
		MakePoolOption(),
		MakePoolOption(),
		MakePoolOption(),
		MakePoolOption(),
	}
	createdAt := time.Now()

	if len(props) > 0 && props[0].Id != nil {
		id = core.NewUniqueEntityId(props[0].Id.String())
	}

	if len(props) > 0 && props[0].Title != "" {
		title = props[0].Title
	}

	if len(props) > 0 && props[0].CreatedAt != nil {
		createdAt = *props[0].CreatedAt
	}

	if len(props) > 0 && len(props[0].Options) > 0 {
		pollOptions = props[0].Options
	}

	return &entities.Poll{
		Entity: core.Entity{
			Id: id,
		},
		Title:     title,
		Options:   pollOptions,
		CreatedAt: createdAt,
	}
}

func MakePrismaPoll(props ...OptionalPollParams) (*entities.Poll, error) {
	fake := faker.New()
	id := core.NewUniqueEntityId()
	ownerId := core.NewUniqueEntityId()
	title := fake.Lorem().Text(100)
	pollOptions := []*entities.PollOption{
		MakePoolOption(),
		MakePoolOption(),
		MakePoolOption(),
		MakePoolOption(),
	}
	createdAt := time.Now()

	if len(props) > 0 && props[0].Id != nil {
		id = core.NewUniqueEntityId(props[0].Id.String())
	}

	if len(props) > 0 && props[0].Title != "" {
		title = props[0].Title
	}

	if len(props) > 0 && props[0].CreatedAt != nil {
		createdAt = *props[0].CreatedAt
	}

	if len(props) > 0 && len(props[0].Options) > 0 {
		pollOptions = props[0].Options
	}

	if len(props) > 0 && props[0].OwnerId != nil {
		ownerId = *props[0].OwnerId
	}

	_, err := prisma.GetDB().Poll.CreateOne(
		db.Poll.Title.Set(title),
		db.Poll.Owner.Link(db.Voter.ID.Equals(ownerId.String())),
		db.Poll.ID.Set(id.String()),
		db.Poll.CreatedAt.Set(createdAt),
	).Exec(configs.GetConfig().Ctx)

	if err != nil {
		return nil, err
	}

	for _, option := range pollOptions {
		_, err := prisma.GetDB().PollOption.CreateOne(
			db.PollOption.Title.Set(option.Title),
			db.PollOption.Poll.Link(db.Poll.ID.Equals(id.String())),
		).Exec(configs.GetConfig().Ctx)

		if err != nil {
			return nil, err
		}
	}

	return &entities.Poll{
		Entity: core.Entity{
			Id: id,
		},
		OwnerId:   ownerId,
		Title:     title,
		Options:   pollOptions,
		CreatedAt: createdAt,
	}, nil
}
