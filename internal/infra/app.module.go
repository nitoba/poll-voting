package infra

import (
	"github.com/nitoba/poll-voting/internal/infra/database"
	"github.com/nitoba/poll-voting/internal/infra/http"
	"github.com/nitoba/poll-voting/pkg/module"
)

type AppModule struct {
	module.Module
}

func NewAppModule() *AppModule {
	m := &AppModule{
		Module: module.Module{
			Imports: module.Imports{
				database.NewDatabaseModule(),
				http.NewHttpModule(),
			},
		},
	}
	m.Build()
	return m
}
