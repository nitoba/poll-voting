package infra

import (
	"github.com/nitoba/poll-voting/pkg/module"
)

type AppModule struct {
	module.Module
}

func NewAppModule(options module.Module) *AppModule {
	m := &AppModule{options}
	m.Build()
	return m
}
