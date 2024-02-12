package mappers

import (
	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/prisma/db"
)

func PollToDomain(poll *db.PollModel) *entities.Poll {

	var options []*entities.PollOption

	for _, option := range poll.Options() {
		options = append(options, &entities.PollOption{
			Entity: *core.NewEntity(core.NewUniqueEntityId(option.ID)),
			Title:  option.Title,
		})
	}

	return &entities.Poll{
		Entity:    *core.NewEntity(core.NewUniqueEntityId(poll.ID)),
		Title:     poll.Title,
		Options:   options,
		OwnerId:   core.NewUniqueEntityId(poll.OwnerID),
		CreatedAt: poll.CreatedAt,
	}
}

func PollToDomainList(polls []db.PollModel) []*entities.Poll {
	var list []*entities.Poll
	for _, poll := range polls {
		list = append(list, PollToDomain(&poll))
	}
	return list
}
