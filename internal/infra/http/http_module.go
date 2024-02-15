package http

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	infra_cryptography "github.com/nitoba/poll-voting/internal/infra/cryptography"
	"github.com/nitoba/poll-voting/internal/infra/database"
	infra_repositories "github.com/nitoba/poll-voting/internal/infra/database/prisma/repositories"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
	"github.com/nitoba/poll-voting/pkg/module"
	"github.com/nitoba/poll-voting/prisma/db"
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
			Name: "voterRepository",
			Provide: func(ctn module.Container) (interface{}, error) {
				return infra_repositories.NewVotersRepositoryPrisma(ctn.Get("db").(*db.PrismaClient)), nil
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
