package factories

import (
	"github.com/jaswdr/faker"
	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type OptionalPollOptionParams struct {
	Id    *core.UniqueEntityId
	Title *string
}

func MakePoolOption(props ...OptionalPollOptionParams) *entities.PollOption {
	fake := faker.New()
	id := core.NewUniqueEntityId()
	title := fake.Lorem().Word()
	if len(props) > 0 && props[0].Id != nil {
		id = core.NewUniqueEntityId(props[0].Id.String())
	}

	if len(props) > 0 && props[0].Title != nil {
		title = *props[0].Title
	}

	return &entities.PollOption{
		Id:    id,
		Title: title,
	}
}
