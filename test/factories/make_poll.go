package factories

import (
	"time"

	"github.com/jaswdr/faker"
	"github.com/nitoba/poll-voting/internal/domain/core/entity"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type OptionalPollParams struct {
	Id        *entity.UniqueEntityId
	Title     *string
	CreatedAt *time.Time
	Options   []*entities.PollOption
}

func MakePool(props ...OptionalPollParams) *entities.Poll {
	fake := faker.New()
	id := entity.NewUniqueEntityId()
	title := fake.Lorem().Text(100)
	pollOptions := []*entities.PollOption{
		MakePoolOption(),
		MakePoolOption(),
		MakePoolOption(),
		MakePoolOption(),
	}
	createdAt := time.Now()

	if len(props) > 0 && props[0].Id != nil {
		id = entity.NewUniqueEntityId(props[0].Id.String())
	}

	if len(props) > 0 && props[0].Title != nil {
		title = *props[0].Title
	}

	if len(props) > 0 && props[0].CreatedAt != nil {
		createdAt = *props[0].CreatedAt
	}

	if len(props) > 0 && len(props[0].Options) > 0 {
		pollOptions = props[0].Options
	}

	return &entities.Poll{
		Entity: entity.Entity{
			Id: id,
		},
		Title:     title,
		Options:   pollOptions,
		CreatedAt: createdAt,
	}
}
