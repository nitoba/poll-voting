package module

import (
	"slices"

	"github.com/sarulabs/di/v2"
)

type module interface {
	GetProviders() Providers
	Build()
}

type provider struct {
	Name        string
	IsSingleton bool
	Provide     func(ctn di.Container) (interface{}, error)
}

type Imports []module
type Providers []provider

type Module struct {
	Imports   Imports
	Providers Providers
}

func (m *Module) Build() {
	m.revolveProvidersFromImports()
}

func (m *Module) GetProviders() Providers {
	return m.Providers
}

func (m *Module) revolveProvidersFromImports() {
	for _, i := range m.Imports {
		importDeps := i.GetProviders()
		if len(importDeps) == 0 {
			i.Build()
		}

		for _, dep := range importDeps {
			alreadyInProviders := slices.ContainsFunc(m.Providers, func(p provider) bool {
				return p.Name == dep.Name
			})

			if !alreadyInProviders {
				m.Providers = append(m.Providers, dep)
			}
		}
	}
}
