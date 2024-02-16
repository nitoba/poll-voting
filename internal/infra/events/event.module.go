package events

import (
	"github.com/nitoba/poll-voting/internal/domain/notification/application/subscribers"
	"github.com/nitoba/poll-voting/internal/domain/notification/application/usecases"
	"github.com/nitoba/poll-voting/internal/infra/database/redis/repositories"
	"github.com/nitoba/poll-voting/pkg/di"
	"github.com/nitoba/poll-voting/pkg/module"
)

type EventModule struct {
	module.Module
}

func NewEventModule() *EventModule {
	ctn := di.GetContainer()
	usecase := ctn.Get("updateVotingCountUseCase").(*usecases.UpdateVotingCountUseCase)
	repo := ctn.Get("countingVotingRepository").(*repositories.CountingVotingRepositoryRedis)
	subscribers.NewOnVoteCreatedHandler(usecase, repo)
	subscribers.NewOnVoteChangedHandler(usecase, repo)
	return &EventModule{
		Module: module.Module{
			Providers: module.Providers{},
		},
	}
}
