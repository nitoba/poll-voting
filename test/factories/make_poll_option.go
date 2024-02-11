package factories

import (
	"github.com/jaswdr/faker"
	"github.com/nitoba/poll-voting/internal/domain/core/entity"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type OptionalPollOptionParams struct {
	Id    *entity.UniqueEntityId
	Title *string
}

func MakePoolOption(props ...OptionalPollOptionParams) *entities.PollOption {
	fake := faker.New()
	id := entity.NewUniqueEntityId()
	title := fake.Lorem().Word()
	if len(props) > 0 && props[0].Id != nil {
		id = entity.NewUniqueEntityId(props[0].Id.String())
	}

	if len(props) > 0 && props[0].Title != nil {
		title = *props[0].Title
	}

	return &entities.PollOption{
		Entity: entity.Entity{
			Id: id,
		},
		Title: title,
	}
}
