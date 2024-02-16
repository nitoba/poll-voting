package infra

import (
	"github.com/nitoba/poll-voting/internal/infra/database"
	"github.com/nitoba/poll-voting/internal/infra/http"
	"github.com/nitoba/poll-voting/internal/infra/messaging"
	"github.com/nitoba/poll-voting/pkg/module"
)

type AppModule struct {
	module.Module
}

func NewAppModule() *AppModule {
	m := &AppModule{
		Module: module.Module{
			Imports: module.Imports{
				messaging.NewMessagingModule(),
				database.NewDatabaseModule(),
				http.NewHttpModule(),
			},
		},
	}
	m.Build()
	return m
}
