package http

import (
	usecase_notification "github.com/nitoba/poll-voting/internal/domain/notification/application/usecases"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	infra_cryptography "github.com/nitoba/poll-voting/internal/infra/cryptography"
	"github.com/nitoba/poll-voting/internal/infra/database"
	infra_repositories "github.com/nitoba/poll-voting/internal/infra/database/prisma/repositories"
	redis_repositories "github.com/nitoba/poll-voting/internal/infra/database/redis/repositories"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
	"github.com/nitoba/poll-voting/internal/infra/http/ws"
	"github.com/nitoba/poll-voting/pkg/module"
	"github.com/nitoba/poll-voting/prisma/db"
	"github.com/redis/go-redis/v9"
)

type HttpModule struct {
	module.Module
}

func NewHttpModule() *HttpModule {
	providers := module.Providers{
		{
			Name: "hasher",
			Provide: func(ctn module.Container) (interface{}, error) {
				return infra_cryptography.CreateBCryptHasher(), nil
			},
		},
		{
			Name: "encrypter",
			Provide: func(ctn module.Container) (interface{}, error) {
				return infra_cryptography.NewJWTEncrypter(), nil
			},
		},
		{
			Name: "updateCountingVotesPublisher",
			Provide: func(ctn module.Container) (interface{}, error) {
				return ws.NewUpdateCountingVotesPublisher(), nil
			},
		},
		{
			Name: "countingVotingRepository",
			Provide: func(ctn module.Container) (interface{}, error) {
				return redis_repositories.NewCountingVotingRepositoryRedis(ctn.Get("redis").(*redis.Conn)), nil
			},
		},
		{
			Name: "voterRepository",
			Provide: func(ctn module.Container) (interface{}, error) {
				return infra_repositories.NewVotersRepositoryPrisma(ctn.Get("db").(*db.PrismaClient)), nil
			},
		},
		{
			Name: "voteRepository",
			Provide: func(ctn module.Container) (interface{}, error) {
				return infra_repositories.NewVoteRepositoryPrisma(ctn.Get("db").(*db.PrismaClient)), nil
			},
		},
		{
			Name: "pollsRepository",
			Provide: func(ctn module.Container) (interface{}, error) {
				return infra_repositories.NewPollsRepositoryPrisma(ctn.Get("db").(*db.PrismaClient)), nil
			},
		},
		{
			Name: "registerVoterUseCase",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecases.NewRegisterVoterUseCase(
					ctn.Get("voterRepository").(*infra_repositories.VotersRepositoryPrisma),
					ctn.Get("hasher").(*infra_cryptography.BCryptHasher),
				), nil
			},
		},
		{
			Name: "authenticateVoterUseCase",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecases.NewAuthenticateUseCase(
					ctn.Get("voterRepository").(*infra_repositories.VotersRepositoryPrisma),
					ctn.Get("hasher").(*infra_cryptography.BCryptHasher),
					ctn.Get("encrypter").(*infra_cryptography.JWTEncrypter),
				), nil
			},
		},
		{
			Name: "createPollUseCase",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecases.NewCreatePollUseCase(
					ctn.Get("pollsRepository").(*infra_repositories.PollsRepositoryPrisma),
					ctn.Get("voterRepository").(*infra_repositories.VotersRepositoryPrisma),
				), nil
			},
		},
		{
			Name: "fetchPollsUseCase",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecases.NewFetchPollsUseCase(
					ctn.Get("pollsRepository").(*infra_repositories.PollsRepositoryPrisma),
				), nil
			},
		},
		{
			Name: "fetchPollsByOwnerUseCase",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecases.NewFetchPollsByOwnerUseCase(
					ctn.Get("pollsRepository").(*infra_repositories.PollsRepositoryPrisma),
				), nil
			},
		},
		{
			Name: "getPollByIdUseCase",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecases.NewGetPollByIdUseCase(
					ctn.Get("pollsRepository").(*infra_repositories.PollsRepositoryPrisma),
					ctn.Get("countingVotingRepository").(*redis_repositories.CountingVotingRepositoryRedis),
				), nil
			},
		},
		{
			Name: "voteOnPollUseCase",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecases.NewVoteOnPollUseCase(
					ctn.Get("voteRepository").(*infra_repositories.VoteRepositoryPrisma),
					ctn.Get("pollsRepository").(*infra_repositories.PollsRepositoryPrisma),
					ctn.Get("voterRepository").(*infra_repositories.VotersRepositoryPrisma),
					ctn.Get("countingVotingRepository").(*redis_repositories.CountingVotingRepositoryRedis),
				), nil
			},
		},
		{
			Name: "updateVotingCountUseCase",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecase_notification.NewUpdateVotingCountUseCase(ctn.Get("updateCountingVotesPublisher").(*ws.UpdateCountingVotesPublisher)), nil
			},
		},
		{
			Name: "onVoteChangedSubscriber",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecase_notification.NewUpdateVotingCountUseCase(ctn.Get("updateCountingVotesPublisher").(*ws.UpdateCountingVotesPublisher)), nil
			},
		},
		{
			Name: "registerController",
			Provide: func(ctn module.Container) (interface{}, error) {
				return controllers.NewRegisterVoterController(ctn.Get("registerVoterUseCase").(*usecases.RegisterVoterUseCase)), nil
			},
		},
		{
			Name: "authenticateController",
			Provide: func(ctn module.Container) (interface{}, error) {
				return controllers.NewAuthenticateVoterController(ctn.Get("authenticateVoterUseCase").(*usecases.AuthenticateUseCase)), nil
			},
		},
		{
			Name: "createPollController",
			Provide: func(ctn module.Container) (interface{}, error) {
				return controllers.NewCreatePollController(ctn.Get("createPollUseCase").(*usecases.CreatePollUseCase)), nil
			},
		},
		{
			Name: "fetchPollsController",
			Provide: func(ctn module.Container) (interface{}, error) {
				return controllers.NewFetchPollsController(ctn.Get("fetchPollsUseCase").(*usecases.FetchPollsUseCase)), nil
			},
		},
		{
			Name: "fetchPollsByOwnerController",
			Provide: func(ctn module.Container) (interface{}, error) {
				return controllers.NewFetchPollsByOwnerController(ctn.Get("fetchPollsByOwnerUseCase").(*usecases.FetchPollsByOwnerUseCase)), nil
			},
		},
		{
			Name: "getPollByIdController",
			Provide: func(ctn module.Container) (interface{}, error) {
				return controllers.NewGetPollByIdController(ctn.Get("getPollByIdUseCase").(*usecases.GetPollByIdUseCase)), nil
			},
		},
		{
			Name: "voteOnPollController",
			Provide: func(ctn module.Container) (interface{}, error) {
				return controllers.NewVoteOnPollController(ctn.Get("voteOnPollUseCase").(*usecases.VoteOnPollUseCase)), nil
			},
		},
		{
			Name: "updateCountingVotesController",
			Provide: func(ctn module.Container) (interface{}, error) {
				return controllers.NewUpdateCountingVotesController(), nil
			},
		},
	}

	m := &HttpModule{
		Module: module.Module{
			Imports: module.Imports{
				database.NewDatabaseModule(),
			},
			Providers: providers,
		},
	}

	m.Build()

	return m
}
