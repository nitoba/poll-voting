package infra

import (
	"slices"

	"github.com/nitoba/poll-voting/pkg/module"
)

type AppModule struct {
	Imports   []module.Module
	Providers []module.Provider
}

func (m *AppModule) Build() {
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

func NewAppModule(options module.NewModule) *AppModule {
	m := &AppModule{
		Imports:   options.Imports,
		Providers: options.Providers,
	}

	m.Build()

	return m
}
