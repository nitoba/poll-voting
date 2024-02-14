package database

import (
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	"github.com/nitoba/poll-voting/pkg/module"
)

type DatabaseModule struct {
	module.Module
}

func NewDatabaseModule() *DatabaseModule {
	m := &DatabaseModule{
		Module: module.Module{
			Providers: module.Providers{
				{
					Name: "db",
					Provide: func(ctn module.Container) (interface{}, error) {
						return prisma.GetDB(), nil
					},
				},
			},
		},
	}
	m.Build()
	return m
}
