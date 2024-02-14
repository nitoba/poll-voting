package http

import (
	"slices"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	infra_cryptography "github.com/nitoba/poll-voting/internal/infra/cryptography"
	"github.com/nitoba/poll-voting/internal/infra/database"
	infra_repositories "github.com/nitoba/poll-voting/internal/infra/database/prisma/repositories"
	"github.com/nitoba/poll-voting/internal/infra/http/controllers"
	"github.com/nitoba/poll-voting/pkg/module"
	"github.com/nitoba/poll-voting/prisma/db"
)

type HttpModule struct {
	Imports   []module.Module
	Providers []module.Provider
}

func (m *HttpModule) Build() {
	for _, i := range m.Imports {
		i.Build()
		importDeps := i.GetDependencies()

		for _, dep := range importDeps {
			alreadyInProviders := slices.ContainsFunc(m.Providers, func(p module.Provider) bool {
				return p.Name == dep.Name
				// return reflect.TypeOf(dep.Provide).String() == reflect.TypeOf(p.Provide).String()
			})

			if !alreadyInProviders {
				m.Providers = append(m.Providers, dep)
			}
		}
	}
}

func (m *HttpModule) GetDependencies() []module.Provider {
	return m.Providers
}

func NewHttpModule() *HttpModule {
	deps := []module.Provider{
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
		Imports: []module.Module{
			database.NewDatabaseModule(),
		},
		Providers: deps,
	}

	return m
}
