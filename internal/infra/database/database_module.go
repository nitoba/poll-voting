package database

import (
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	"github.com/nitoba/poll-voting/internal/infra/module"
	"github.com/sarulabs/di"
)

type DatabaseModule struct {
	module.Module
}

func NewDatabaseModule(moduleOptions ...module.ModuleConfigOptions) *DatabaseModule {
	deps := []di.Def{
		{
			Name:  "db",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return prisma.GetDB(), nil
			},
			Close: func(_ interface{}) error {
				prisma.Disconnect()
				return nil
			},
		},
	}
	if len(moduleOptions) > 0 {
		deps = moduleOptions[0].Dependencies
	}
	return &DatabaseModule{
		Module: *module.NewModule(module.ModuleConfigOptions{
			Dependencies: deps,
		}),
	}
}
