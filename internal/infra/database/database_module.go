package database

import (
	"slices"

	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	"github.com/nitoba/poll-voting/pkg/module"
)

type DatabaseModule struct {
	Imports   []module.Module
	Providers []module.Provider
}

func (m *DatabaseModule) Build() {
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

func (m *DatabaseModule) GetDependencies() []module.Provider {
	return m.Providers
}

func NewDatabaseModule() *DatabaseModule {
	m := &DatabaseModule{
		Providers: []module.Provider{
			{
				Name: "db",
				Provide: func(ctn module.Container) (interface{}, error) {
					return prisma.GetDB(), nil
				},
			},
		},
	}
	return m
}
