package module

import "slices"

type Container interface {
	Get(name string) interface{}
}

type module interface {
	GetDependencies() Providers
	Build()
}

type provider struct {
	Name    string
	Provide func(ctn Container) (interface{}, error)
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

func (m *Module) GetDependencies() Providers {
	return m.Providers
}

func (m *Module) revolveProvidersFromImports() {
	for _, i := range m.Imports {
		importDeps := i.GetDependencies()
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
