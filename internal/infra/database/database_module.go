package database

import (
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	"github.com/nitoba/poll-voting/pkg/module"
)

type DatabaseModule struct {
	Imports   []module.Module
	Providers []module.Provider
}

func (m *DatabaseModule) Build() {
	m.Providers = module.RevolveProvidersFromImports(m.Imports, m.Providers)
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
