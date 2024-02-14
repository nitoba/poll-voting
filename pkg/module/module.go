package module

import "slices"

type Container interface {
	Get(name string) interface{}
}

type Module interface {
	GetDependencies() []Provider
	Build()
}

type Provider struct {
	Name    string
	Provide func(ctn Container) (interface{}, error)
}

type NewModule struct {
	Imports   []Module
	Providers []Provider
}

func RevolveProvidersFromImports(imports []Module, incomingProviders []Provider) []Provider {
	var providers []Provider = incomingProviders
	for _, i := range imports {
		i.Build()
		importDeps := i.GetDependencies()

		for _, dep := range importDeps {
			alreadyInProviders := slices.ContainsFunc(providers, func(p Provider) bool {
				return p.Name == dep.Name
			})

			if !alreadyInProviders {
				providers = append(providers, dep)
			}
		}
	}
	return providers
}
