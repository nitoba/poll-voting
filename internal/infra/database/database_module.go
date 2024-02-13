package database

import (
	"github.com/nitoba/poll-voting/internal/infra/database/prisma"
	"github.com/sarulabs/di"
)

type DatabaseModule struct {
	PrismaDB di.Def
}

func NewDatabaseModule() *DatabaseModule {
	return &DatabaseModule{
		PrismaDB: di.Def{
			Name:  "db",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return prisma.GetDB(), nil
			},
		},
	}
}
