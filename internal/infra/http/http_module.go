package http

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	infra_cryptography "github.com/nitoba/poll-voting/internal/infra/cryptography"
	infra_repositories "github.com/nitoba/poll-voting/internal/infra/database/prisma/repositories"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
	"github.com/nitoba/poll-voting/internal/infra/module"
	"github.com/nitoba/poll-voting/prisma/db"
	"github.com/sarulabs/di"
)

type HttpModule struct {
	module.Module
}

func NewHttpModule(moduleOptions ...module.ModuleConfigOptions) *HttpModule {
	deps := []di.Def{
		{
			Name:  "hasher",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return infra_cryptography.CreateBCryptHasher(), nil
			},
		},
		{
			Name:  "voterRepository",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return infra_repositories.NewVotersRepositoryPrisma(ctn.Get("db").(*db.PrismaClient)), nil
			},
		},
		{
			Name:  "registerVoterUseCase",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return usecases.NewRegisterVoterUseCase(ctn.Get("voterRepository").(*infra_repositories.VotersRepositoryPrisma), ctn.Get("hasher").(*infra_cryptography.BCryptHasher)), nil
			},
		},
		{
			Name:  "registerController",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return controllers.NewRegisterVoterController(ctn.Get("registerVoterUseCase").(*usecases.RegisterVoterUseCase)), nil
			},
		},
	}

	if len(moduleOptions) > 0 {
		deps = moduleOptions[0].Dependencies
	}

	return &HttpModule{
		Module: *module.NewModule(module.ModuleConfigOptions{
			Dependencies: deps,
		}),
	}
}
