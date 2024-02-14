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
			Name: "voterRepository",
			Provide: func(ctn module.Container) (interface{}, error) {
				return infra_repositories.NewVotersRepositoryPrisma(ctn.Get("db").(*db.PrismaClient)), nil
			},
		},
		{
			Name: "registerVoterUseCase",
			Provide: func(ctn module.Container) (interface{}, error) {
				return usecases.NewRegisterVoterUseCase(ctn.Get("voterRepository").(*infra_repositories.VotersRepositoryPrisma), ctn.Get("hasher").(*infra_cryptography.BCryptHasher)), nil
			},
		},
		{
			Name: "registerController",
			Provide: func(ctn module.Container) (interface{}, error) {
				return controllers.NewRegisterVoterController(ctn.Get("registerVoterUseCase").(*usecases.RegisterVoterUseCase)), nil
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
